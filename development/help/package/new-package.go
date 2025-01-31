package help

import (
	"development/util"
	"fmt"
	"os"
)

func CreateNewPackage(name string) {
	// check if already exists
	info, err := os.Stat(util.GetBackendPath() + "pkg" + name)

	if err == nil && info.IsDir() {
		println("Package already exists")
		return
	} else {
		fmt.Println("Creating package", name, " at ", util.GetBackendPath()+"pkg/"+name)
		err := os.Mkdir(util.GetBackendPath()+"pkg/"+name, 0755)
		if err != nil {
			fmt.Println("Error creating package:", err)
			return
		} else {
			CreateFile(util.GetBackendPath() + "pkg/" + name + "/service/service.go")
		}
	}
}

func CreateFile(path string) {
	info, err := os.Stat(path)

	if err == nil && !info.IsDir() {
		println("File already exists")
		return
	} else {
		// if the folder does not exist, create folder with -p
		folder := path[:len(path)-len(info.Name())]

		if _, err := os.Stat(folder); os.IsNotExist(err) {
			err := os.MkdirAll(folder, 0755)
			if err != nil {
				fmt.Println("Error creating folder:", err)
				return
			}
		}

		file, err := os.Create(path)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}

		defer file.Close()
	}
}
