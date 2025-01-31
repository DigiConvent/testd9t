package sys_repository_test

import (
	"testing"

	"github.com/DigiConvent/testd9t/core/db"
	sys_repository "github.com/DigiConvent/testd9t/pkg/sys/repository"
)

func TestListFlavours(t *testing.T) {
	testDB := db.NewTestSqliteDB("sys.testlistflavours")
	flavours, status := sys_repository.NewSysRepository(testDB).ListFlavoursForVersion()

	if status.Err() {
		t.Errorf("Failed to list flavours: %s", status.Message)
	}

	if len(flavours) == 0 {
		t.Errorf("No flavours found")
	}
}
