package sys_service_test

import (
	"testing"

	"github.com/DigiConvent/testd9t/core/db"
	sys_repository "github.com/DigiConvent/testd9t/pkg/sys/repository"
	sys_service "github.com/DigiConvent/testd9t/pkg/sys/service"
)

var testDB db.DatabaseInterface

func GetTestSysService(dbName string) sys_service.SysServiceInterface {
	if testDB == nil {
		testDB = db.NewTestSqliteDB(dbName)
	}
	repo := sys_repository.NewSysRepository(testDB)
	return sys_service.NewSysService(repo)
}

func TestMain(m *testing.M) {
	GetTestSysService("sys")
	defer testDB.DeleteDatabase()
	m.Run()
}
