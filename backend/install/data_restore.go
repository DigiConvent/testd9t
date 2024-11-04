package install

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func RestoreBackup(zipPath string) error {
	timestamp := time.Now().Unix()

	destFolder := "backup_" + strconv.Itoa(int(timestamp))
	dest := filepath.Join(os.TempDir(), destFolder)

	os.Mkdir(dest, 0755)

	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return fmt.Errorf("failed to open zip file: %w", err)
	}
	defer r.Close()

	for _, f := range r.File {
		dstFilePath := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(dstFilePath, os.ModePerm); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", dstFilePath, err)
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(dstFilePath), os.ModePerm); err != nil {
			return fmt.Errorf("failed to create directory for file %s: %w", dstFilePath, err)
		}

		srcFile, err := f.Open()
		if err != nil {
			return fmt.Errorf("failed to open file %s: %w", f.Name, err)
		}
		defer srcFile.Close()

		dstFile, err := os.Create(dstFilePath)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", dstFilePath, err)
		}
		defer dstFile.Close()

		if _, err := io.Copy(dstFile, srcFile); err != nil {
			return fmt.Errorf("failed to copy file contents to %s: %w", dstFilePath, err)
		}
	}

	err = RestoreDatabase(dest + "/dump.sql")
	if err != nil {
		return fmt.Errorf("failed to restore database: %w", err)
	}

	return nil
}
