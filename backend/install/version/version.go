package version

import (
	"context"
	"fmt"
	"os"
	"strings"

	db "github.com/DigiConvent/testd9t/core/pg"
)

var RepositoryUser string = "Digiconvent"
var RepositoryName string = "testd9t"

func Prod() bool {
	return !Dev()
}

func Dev() bool {
	return strings.HasPrefix(os.Args[0], "/tmp/go-build")
}

func ProgramVersion() string {
	if os.Getenv("VERSION") != "" {
		return os.Getenv("VERSION")
	}
	return "-1.-1.-1"
}

func MigrationVersion() *SemVer {
	conn := db.GetPGDBConnection()

	version := SemVer{Major: -1, Minor: -1, Patch: -1}
	row := conn.QueryRow(context.Background(), "SELECT major, minor, patch FROM versions ORDER BY major DESC, minor DESC, patch DESC LIMIT 1")
	err := row.Scan(&version.Major, &version.Minor, &version.Patch)
	if err != nil {
		fmt.Println("Error fetching migration version:", err)
		return nil
	}
	return &version
}
