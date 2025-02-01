package cli

import (
	"fmt"
	"os"
	"path"

	"github.com/DigiConvent/testd9t/core/db"
)

func ResetDB() {

	fmt.Println("--reset-db")

	dbPath := db.DatabasePath
	entries, err := os.ReadDir(dbPath)
	if err != nil {
		fmt.Println("Could not find db directory ", dbPath)
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
