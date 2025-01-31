package sys_service_test

import (
	"testing"

	"github.com/DigiConvent/testd9t/core/db"
	sys_repository "github.com/DigiConvent/testd9t/pkg/sys/repository"
	sys_service "github.com/DigiConvent/testd9t/pkg/sys/service"
)

func TestSysService(t *testing.T) {
	t.Skipped()
}

func GetTestSysService() sys_service.SysServiceInterface {
	testDB := db.NewTestSqliteDB("sys")
	mockRepo := sys_repository.NewSysRepository(testDB)
	return sys_service.NewSysService(mockRepo)
}
