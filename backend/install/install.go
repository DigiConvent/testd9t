package install

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/DigiConvent/testd9t/core/file_repo"
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

func (s Script) Prepare(flavour string) {
	repo := file_repo.NewRepoRemote()
	if len(s.RequiresFiles) > 0 {
		for _, file := range s.RequiresFiles {
			contents, err := repo.GetRawFile("install/" + flavour + "/" + file)
			if err != nil {
				fmt.Println("Error downloading file:", err)
				return
			}
			storeFile(file, string(contents))
		}
	}

	doScriptContents, err := repo.GetRawFile("install/" + flavour + "/do_" + s.Name + ".sh")
	if err != nil {
		fmt.Println("Error downloading script:", err)
		return
	}
	undoScriptContents, err := repo.GetRawFile("install/" + flavour + "/undo_" + s.Name + ".sh")
	if err != nil {
		fmt.Println("Error downloading script:", err)
		return
	}

	storeFile("do_"+s.Name+".sh", string(doScriptContents))
	storeFile("undo_"+s.Name+".sh", string(undoScriptContents))

	if len(s.Input) > 0 {
		for i, input := range s.Input {
			s.Input[i].Value = promptUser(input.Name, input.Default)
		}
	}
}

// Function to prompt the user for input
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
func (s Script) Do(fix bool) error {
	args := []string{"-c", getFilePath("do_" + s.Name + ".sh")}
	for _, input := range s.Input {
		args = append(args, input.Value)
	}
	cmd := exec.Command("bash", args...)
	result, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("❌ do_" + s.Name + "...")
		if fix {
			fmt.Println(" fixing...")
			err := s.Undo()
			if err == nil {
				s.Do(false)
			}
		} else {
			fmt.Println(string(result))
		}
		return err
	}
	fmt.Println("✅ do_" + s.Name)
	return nil
}

func (s Script) Undo() error {
	args := []string{"-c", getFilePath("undo_" + s.Name + ".sh")}
	for _, input := range s.Input {
		args = append(args, input.Value)
	}
	cmd := exec.Command("bash", args...)
	result, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("failed: " + string(result))
		return err
	}
	return nil
}

type InstallationProtocol struct {
	Scripts []Script `json:"scripts"`
	Files   []string `json:"path"`
}

func Install(flavour *string, force bool, clearCache bool) {
	repo := file_repo.NewRepoRemote()
	var protocol InstallationProtocol

	raw, err := repo.GetRawFile("install/" + *flavour + "/install.json")
	if err != nil {
		fmt.Println("Error fetching installation protocol:", err)
		return
	}

	err = json.Unmarshal(raw, &protocol)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
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
		script.Do(force)
	}
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
