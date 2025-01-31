package install

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

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
	}
	// these are the output files that are stored under /tmp/name.txt
	Output []string `json:"output"`
}

type InstallationProtocol struct {
	Scripts []Script `json:"scripts"`
	Files   []string `json:"path"`
}

func Install(flavour *string, force bool, clearCache bool) {
	var requiredFiles map[string]string = make(map[string]string)
	var scripts map[string]string = make(map[string]string)
	repo := file_repo.NewRepoLocal()
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

		prefix := "install/" + *flavour + "/"
		for _, file := range protocol.Scripts[i].RequiresFiles {
			contents, err := repo.GetRawFile(prefix + file)
			if err != nil {
				fmt.Println("Error downloading file:", err)
				return
			}
			requiredFiles[file] = string(contents)
		}

		for _, method := range []string{"do", "undo"} {
			contents, err := repo.GetRawFile(prefix + method + "_" + protocol.Scripts[i].Name + ".sh")
			scripts[method+"_"+protocol.Scripts[i].Name+".sh"] = string(contents)

			if err != nil {
				fmt.Println("Error downloading script:", err)
				return
			}

			panic("This still needs to be implemented")
			// err = exec.Command("chmod", "+x", method+"_"+protocol.Scripts[i].Name+".sh").Run()
			// if err != nil {
			// 	fmt.Println("Error making script executable:", err)
			// 	return
			// }
		}
	}

	scriptsRan := make([]string, len(protocol.Scripts))
	for i := 0; i < len(protocol.Scripts); i++ {
		for _, file := range protocol.Scripts[i].RequiresFiles {
			if _, err := os.Stat(file); os.IsNotExist(err) {
				fmt.Println("Error: file", file, "does not exist")
				return
			}
		}

		for _, input := range protocol.Scripts[i].Input {
			if _, err := os.Stat(input.Name + ".txt"); os.IsNotExist(err) {
				err = exec.Command("echo", input.Default).Run()
				if err != nil {
					fmt.Println("Error writing default value to file", input.Name+".txt:", err)
					return
				}
			}
		}

		output, err := ExecuteScript(protocol.Scripts[i].Name, force)
		scriptsRan[i] = protocol.Scripts[i].Name
		if err != nil {
			fmt.Println("Error running script do_"+protocol.Scripts[i].Name+":", err)
			fmt.Println("Output:", string(output))
			for j := i; j >= 0; j-- {
				fmt.Println("Undoing script", protocol.Scripts[j].Name)
				exec.Command("undo_" + protocol.Scripts[j].Name + ".sh").Run()
			}
			return
		} else {
			fmt.Println("✅ Script do_" + protocol.Scripts[i].Name)
		}
	}
}

func ExecuteScript(scriptName string, force bool) (string, error) {
	output, err := exec.Command("/bin/sh", "/do_"+scriptName+".sh").CombinedOutput()
	if err != nil {
		if !force {
			return string(output), err
		} else {
			fmt.Println("❌ Script do_" + scriptName + " failed, doing a fix 👷🏻‍♂️")
		}
	}
	output, err = exec.Command("/bin/sh", "/undo_"+scriptName+".sh").CombinedOutput()
	if err != nil {
		fmt.Println("Error fixing ", scriptName, ":", string(output))
		return string(output), err
	}

	output, err = exec.Command("/bin/sh", "/do_"+scriptName+".sh").CombinedOutput()
	return string(output), err
}
