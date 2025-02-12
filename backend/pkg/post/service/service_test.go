package post_service_test

import (
	"os"
	"testing"

	constants "github.com/DigiConvent/testd9t/core/const"
	"github.com/DigiConvent/testd9t/core/db"
	post_repository "github.com/DigiConvent/testd9t/pkg/post/repository"
	post_service "github.com/DigiConvent/testd9t/pkg/post/service"
	post_setup "github.com/DigiConvent/testd9t/pkg/post/setup"
)

var testDB db.DatabaseInterface

func GetTestPostService(dbName string) post_service.PostServiceInterface {
	os.Setenv(constants.CERTIFICATES_PATH, "/tmp/testd9t/certificates")
	if testDB == nil {
		post_setup.Setup()
		testDB = db.NewTestSqliteDB(dbName)
	}
	repo := post_repository.NewPostRepository(testDB)
	return post_service.NewPostService(repo, false)
}

func TestMain(m *testing.M) {
	GetTestPostService("post")
	defer testDB.DeleteDatabase()
	m.Run()

}
