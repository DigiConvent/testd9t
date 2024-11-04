package install

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/DigiConvent/testd9t/install/version"
)

var dumpPath = "/tmp/dump.sql"
var versionPath = "/tmp/version.txt"

func CreateBackup(filename *string) {
	migrationVersion := version.MigrationVersion()
	programVersion := version.ProgramVersion()
	var runtype string
	if version.Dev() {
		runtype = "dev"
	} else {
		runtype = "prod"
	}

	timeOfBackup := time.Now().UTC().Format(time.RFC3339)
	os.WriteFile(versionPath, []byte(fmt.Sprintf("Run type: %v\nMigration version: %v\nProgram version: %v\nTime of backup: %v", runtype, migrationVersion, programVersion, timeOfBackup)), 0644)
	err := BackupDatabase(dumpPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	files := []string{dumpPath, EnvironmentVars.DataPath, versionPath}
	if !strings.HasSuffix(*filename, ".zip") {
		*filename = *filename + ".zip"
	}
	err = createZip(*filename, files)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Run the following to restore the backup:")
	if version.Dev() {
		fmt.Println("go run main.go --restore " + *filename)
	} else {
		fmt.Println(os.Args[0], "--restore", *filename)
	}

	os.Remove(dumpPath)
	os.Remove(versionPath)
}

func createZip(zipFileName string, paths []string) error {
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %v", err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range paths {
		fileInfo, err := os.Stat(file)
		if err != nil {
			return fmt.Errorf("failed to get file info for %s: %v", file, err)
		}
		if fileInfo.IsDir() {
			err = filepath.Walk(file, func(filePath string, info os.FileInfo, err error) error {
				if err != nil {
					return fmt.Errorf("failed to access path %q: %v", filePath, err)
				}
				if info.IsDir() {
					return nil
				}

				return addFileToZip(zipWriter, filePath, file)
			})
		} else {
			err = addFileToZip(zipWriter, file, "")
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func addFileToZip(zipWriter *zip.Writer, filePath, folderPath string) error {
	fileToZip, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %v", filePath, err)
	}
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info for %s: %v", filePath, err)
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return fmt.Errorf("failed to create header for %s: %v", filePath, err)
	}

	if folderPath != "" {
		relativePath, err := filepath.Rel(folderPath, filePath)
		if err != nil {
			return fmt.Errorf("failed to create relative path for %s: %v", filePath, err)
		}
		header.Name = relativePath
	} else {
		segments := strings.Split(filePath, "/")
		header.Name = segments[len(segments)-1]
	}

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return fmt.Errorf("failed to create zip entry for %s: %v", filePath, err)
	}

	_, err = io.Copy(writer, fileToZip)
	if err != nil {
		return fmt.Errorf("failed to copy file content for %s: %v", filePath, err)
	}

	return nil
}
