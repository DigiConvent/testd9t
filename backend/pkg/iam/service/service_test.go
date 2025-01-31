package iam_service_test

import (
	"fmt"
	"testing"

	db_test "github.com/DigiConvent/testd9t/core/db"
	iam_repository "github.com/DigiConvent/testd9t/pkg/iam/repository"
	iam_service "github.com/DigiConvent/testd9t/pkg/iam/service"
)

var testDB db_test.DatabaseInterface

func GetTestIAMService(dbName string) iam_service.IAMServiceInterface {
	testDB = db_test.NewTestSqliteDB(dbName)
	testDB.MigratePackage()
	repo := iam_repository.NewIAMRepository(testDB)
	return iam_service.NewIAMService(repo)
}

func TestMain(m *testing.M) {
	GetTestIAMService("iam")

	m.Run()
	fmt.Println("Don't do anything")

	testDB.DeleteDatabase()
}
