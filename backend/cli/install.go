package cli

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/DigiConvent/testd9t/core/file_repo"
	"github.com/DigiConvent/testd9t/core/log"
	sys_service "github.com/DigiConvent/testd9t/pkg/sys/service"
)

type Script struct {
	Step          int      `json:"step"`
	Optional      bool     `json:"optional"`
	RequiresSteps []int    `json:"requires_steps"`
	RequiresFiles []string `json:"requires_files"`
	// this means that it is a shell script that is downloadable from
	// https://raw.githubusercontent.com/DigiConvent/testd9t/refs/heads/main/install/<flavour>/<name>.sh
	Name string `json:"name"`
	// input are basically named arguments and they are stored from other steps into files inside
	// /tmp/name.txt or, if it doesn't exist, use the default or prompt the user for input
	Input []struct {
		Name    string `json:"name"`
		Default string `json:"default"`
		Value   string `json:"value"`
	}
	// these are the output files that are stored under /tmp/name.txt
	Output []string `json:"output"`
}

type Input struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (input *Input) promptUser() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", input.Name)
	userInput, _ := reader.ReadString('\n')
	userInput = userInput[:len(userInput)-1]
	input.Value = userInput
}

func (s Script) Prepare(flavour string) {
	repo := file_repo.NewRepoRemote()
	if len(s.RequiresFiles) > 0 {
		for _, file := range s.RequiresFiles {
			contents, err := repo.ReadRawFile("install/" + flavour + "/" + file)
			if err != nil {
				fmt.Println("Error downloading file:", err)
				return
			}
			storeFile(file, string(contents))
		}
	}

	doScriptContents, err := repo.ReadRawFile("install/" + flavour + "/do_" + s.Name + ".sh")
	if err != nil {
		fmt.Println("Error downloading script:", err)
		return
	}
	undoScriptContents, err := repo.ReadRawFile("install/" + flavour + "/undo_" + s.Name + ".sh")
	if err != nil {
		fmt.Println("Error downloading script:", err)
		return
	}

	storeFile("do_"+s.Name+".sh", string(doScriptContents))
	storeFile("undo_"+s.Name+".sh", string(undoScriptContents))
}

func promptUser(prompt string, defaultValue string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s [%s]: ", prompt, defaultValue)
	userInput, _ := reader.ReadString('\n')
	userInput = userInput[:len(userInput)-1]
	if userInput == "" {
		return defaultValue
	}
	return userInput
}
func (s Script) Do(fix, verbose bool, inputs map[string]Input) error {
	args := []string{getFilePath("do_" + s.Name + ".sh")}
	for _, input := range s.Input {
		args = append(args, inputs[input.Name].Value)
	}
	cmd := exec.Command("bash", args...)
	if verbose {
		fmt.Println(cmd.String())
	}

	result, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("❌ do_" + s.Name + "...")
		if verbose {
			fmt.Println("failed: " + string(result))
		}
		if fix {
			err := s.Undo(verbose)
			if err == nil {
				s.Do(false, verbose, inputs)
			}
		} else {
			fmt.Println(string(result))
		}
		return err
	}
	fmt.Println("✅ do_" + s.Name)
	return nil
}

func (s Script) Undo(verbose bool) error {
	args := []string{getFilePath("undo_" + s.Name + ".sh")}
	for _, input := range s.Input {
		args = append(args, input.Value)
	}
	cmd := exec.Command("bash", args...)
	if verbose {
		fmt.Println(cmd.String())
	}
	result, err := cmd.CombinedOutput()
	if err != nil {
		if verbose {
			fmt.Println("failed: " + string(result))
		}
		return err
	}
	return nil
}

type InstallationProtocol struct {
	Scripts []Script `json:"scripts"`
	Files   []string `json:"path"`
}

func Install(sysService sys_service.SysServiceInterface, flavour *string, force bool, verbose bool) map[string]Input {
	uid := os.Geteuid()

	if uid != 0 {
		fmt.Println("You need to be root to install")
		os.Exit(1)
	}

	*flavour = strings.ToLower(*flavour)
	fmt.Println("--install", *flavour)

	flavours, status := sysService.ListFlavours()
	if status.Err() {
		fmt.Println("Error fetching flavours:", status.Message)
		os.Exit(1)
	}

	found := false
	for _, availableFlavour := range flavours {
		if availableFlavour == *flavour {
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Flavour", *flavour, "not found")
		choices := strings.Join(flavours, ", ")
		fmt.Println("Available flavours:", choices)
		os.Exit(1)
	}

	inputs := map[string]Input{
		"domain": {Name: "Domain", Value: ""},
		"email":  {Name: "Email", Value: ""},
	}

	for _, input := range inputs {
		input.promptUser()
	}

	repo := file_repo.NewRepoRemote()
	var protocol InstallationProtocol

	raw, err := repo.ReadRawFile("install/" + *flavour + "/install.json")
	if err != nil {
		fmt.Println("Error fetching installation protocol:", err)
		return nil
	}

	err = json.Unmarshal(raw, &protocol)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}

	for i := 0; i < len(protocol.Scripts); i++ {
		for j := i + 1; j < len(protocol.Scripts); j++ {
			if protocol.Scripts[i].Step > protocol.Scripts[j].Step {
				protocol.Scripts[i], protocol.Scripts[j] = protocol.Scripts[j], protocol.Scripts[i]
			}
		}
	}

	for _, script := range protocol.Scripts {
		script.Prepare(*flavour)
		script.Do(force, verbose, inputs)
	}

	bytes := make([]byte, 64)
	_, err = rand.Read(bytes)
	if err != nil {
		fmt.Println("Failed to generate a master password, no idea what to do from here.", err.Error())
	}
	password := hex.EncodeToString(bytes)

	contents, err := os.ReadFile("/home/testd9t/env")
	if err != nil {
		log.Error("Could not store the master password in the env file. No idea what to do from here. " + err.Error())
	}

	contents = []byte("MASTER_PASSWORD=" + password + "\n" + string(contents))
	err = os.WriteFile("/home/testd9t/env", contents, 0644)
	if err != nil {
		log.Error("Could not store the master password in the env file. No idea what to do from here. " + err.Error())
	}
	log.Success("Master password:\n" + password)

	return inputs
}

const dirToStore = "/testd9t/"

func getFilePath(fileName string) string {
	return path.Join(os.TempDir(), dirToStore, fileName)
}
func storeFile(fileName string, contents string) error {
	prefix := path.Join(os.TempDir(), dirToStore)
	os.MkdirAll(prefix, 0755)
	uri := path.Join(prefix, fileName)

	os.Remove(uri)
	file, err := os.Create(uri)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(contents)
	if err != nil {
		return err
	}

	if strings.HasSuffix(fileName, ".sh") {
		err = exec.Command("chmod", "+x", uri).Run()
		if err != nil {
			return err
		}
	}
	return nil
}
