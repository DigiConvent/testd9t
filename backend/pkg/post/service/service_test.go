package post_service_test

import (
	"testing"

	"github.com/DigiConvent/testd9t/core/db"
	post_repository "github.com/DigiConvent/testd9t/pkg/post/repository"
	post_service "github.com/DigiConvent/testd9t/pkg/post/service"
)

var testDB db.DatabaseInterface

func GetTestPostService(dbName string) post_service.PostServiceInterface {
	if testDB == nil {
		testDB = db.NewTestSqliteDB(dbName)
	}
	repo := post_repository.NewPostRepository(testDB)
	return post_service.NewPostService(repo, true)
}

func TestMain(m *testing.M) {
	GetTestPostService("post")
	defer testDB.DeleteDatabase()
	m.Run()

}
