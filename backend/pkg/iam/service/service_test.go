package iam_service_test

import (
	"testing"

	"github.com/DigiConvent/testd9t/core/db"
	iam_repository "github.com/DigiConvent/testd9t/pkg/iam/repository"
	iam_service "github.com/DigiConvent/testd9t/pkg/iam/service"
)

var testDB db.DatabaseInterface

func GetTestIAMService(dbName string) iam_service.IAMServiceInterface {
	if testDB == nil {
		testDB = db.NewTestSqliteDB(dbName)
	}
	repo := iam_repository.NewIAMRepository(testDB)
	return iam_service.NewIAMService(repo)
}

func TestMain(m *testing.M) {
	GetTestIAMService("iam")
	defer testDB.DeleteDatabase()
	m.Run()

}
