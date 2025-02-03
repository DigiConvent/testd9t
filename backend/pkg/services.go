package services

import (
	"github.com/DigiConvent/testd9t/core/db"
	iam_repository "github.com/DigiConvent/testd9t/pkg/iam/repository"
	iam_service "github.com/DigiConvent/testd9t/pkg/iam/service"
	sys_repository "github.com/DigiConvent/testd9t/pkg/sys/repository"
	sys_service "github.com/DigiConvent/testd9t/pkg/sys/service"
)

type Services struct {
	IAMService iam_service.IAMServiceInterface
	SysService sys_service.SysServiceInterface
}

func InitiateServices() *Services {
	sysDB := db.NewSqliteDB("sys")
	sysRepo := sys_repository.NewSysRepository(sysDB)
	sysService := sys_service.NewSysService(sysRepo)
	sysService.Init()

	iamDB := db.NewSqliteDB("iam")
	iamRepo := iam_repository.NewIAMRepository(iamDB)
	iamService := iam_service.NewIAMService(iamRepo)

	// this needs to be called after all the databases are initialised as it migrates only packages of which a database file exists
	sysService.MigrateDatabase(nil)

	return &Services{
		SysService: sysService,
		IAMService: iamService,
	}
}
