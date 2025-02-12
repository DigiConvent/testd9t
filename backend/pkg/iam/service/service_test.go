package iam_service_test

import (
	"os"
	"testing"

	constants "github.com/DigiConvent/testd9t/core/const"
	"github.com/DigiConvent/testd9t/core/db"
	iam_repository "github.com/DigiConvent/testd9t/pkg/iam/repository"
	iam_service "github.com/DigiConvent/testd9t/pkg/iam/service"
	iam_setup "github.com/DigiConvent/testd9t/pkg/iam/setup"
)

var testDB db.DatabaseInterface

func GetTestIAMService(dbName string) iam_service.IAMServiceInterface {
	os.Setenv(constants.CERTIFICATES_PATH, "/tmp/testd9t/certificates")
	if testDB == nil {
		iam_setup.Setup()
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
