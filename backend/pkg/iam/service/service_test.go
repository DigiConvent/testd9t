package iam_service_test

import (
	"os"
	"testing"

	constants "github.com/DigiConvent/testd9t/core/const"
	"github.com/DigiConvent/testd9t/core/db"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	iam_repository "github.com/DigiConvent/testd9t/pkg/iam/repository"
	iam_service "github.com/DigiConvent/testd9t/pkg/iam/service"
	iam_setup "github.com/DigiConvent/testd9t/pkg/iam/setup"
)

var testDB db.DatabaseInterface

func getRootPermissionGroup() string {
	testService := GetTestIAMService("iam")
	facades, _ := testService.ListPermissionGroups()
	for _, facade := range facades {
		if facade.Name == "root" {
			return facade.Id.String()
		}
	}

	testRepo := getTestRepo("iam")
	id, _ := testRepo.CreatePermissionGroup(&iam_domain.PermissionGroupWrite{
		Name: "root",
	})

	return id.String()
}

func GetTestIAMService(dbName string) iam_service.IAMServiceInterface {
	os.Setenv(constants.CERTIFICATES_PATH, "/tmp/testd9t/certificates")
	if testDB == nil {
		iam_setup.Setup()
		testDB = db.NewTestSqliteDB(dbName)
	}
	repo := iam_repository.NewIamRepository(testDB)
	return iam_service.NewIamService(repo)
}

func getTestRepo(dbName string) iam_repository.IAMRepositoryInterface {
	os.Setenv(constants.CERTIFICATES_PATH, "/tmp/testd9t/certificates")
	if testDB == nil {
		iam_setup.Setup()
		testDB = db.NewTestSqliteDB(dbName)
	}
	repo := iam_repository.NewIamRepository(testDB)
	return repo
}

func TestMain(m *testing.M) {
	GetTestIAMService("iam")
	defer testDB.DeleteDatabase()
	m.Run()
}
