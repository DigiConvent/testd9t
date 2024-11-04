package installation_utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(url, filepath string) error {
	// Create the file where the download will be saved
	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	// Make an HTTP GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}
	defer resp.Body.Close()

	// Check for a successful HTTP response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Copy the response body to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}

	fmt.Printf("File downloaded successfully to %s\n", filepath)
	return nil
}
