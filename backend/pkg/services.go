package services

import (
	"fmt"

	"github.com/DigiConvent/testd9t/core"
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

	update(sysService)

	iamDB := db.NewSqliteDB("iam")
	iamRepo := iam_repository.NewIAMRepository(iamDB)
	iamService := iam_service.NewIAMService(iamRepo)

	return &Services{
		SysService: sysService,
		IAMService: iamService,
	}
}

func update(s sys_service.SysServiceInterface) *core.Status {
	sysStatus, status := s.GetSystemStatus()

	if status.Err() {
		return core.InternalError(status.Message)
	}

	packages := db.ListPackages()

	for _, pkg := range packages {
		status := s.MigratePackage(pkg, sysStatus.ProgramVersion)
		if status.Err() && status.Code != 404 {
			fmt.Println("Error migrating package", pkg, ":", status.Message)
		}
	}

	return nil
}
