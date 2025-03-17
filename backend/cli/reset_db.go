package cli

import (
	"fmt"
	"os"
	"path"

	constants "github.com/DigiConvent/testd9t/core/const"
)

func ResetDB() {

	fmt.Println("--reset-db")

	dbPath := os.Getenv(constants.DATABASE_PATH)
	entries, err := os.ReadDir(dbPath)
	if err != nil {
		fmt.Println("Could not find db directory", dbPath)
	}

	for _, entry := range entries {
		dbName := entry.Name()
		if !entry.IsDir() {
			continue
		}

		dbPath := path.Join(dbPath, dbName+".db")
		err := os.RemoveAll(dbPath)
		if err != nil {
			fmt.Println("Error removing db ", err)
		}

		fmt.Println("Removed", dbPath)
	}
	os.Exit(0)
}
