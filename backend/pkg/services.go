package services

import (
	"github.com/DigiConvent/testd9t/core/db"
	iam_repository "github.com/DigiConvent/testd9t/pkg/iam/repository"
	iam_service "github.com/DigiConvent/testd9t/pkg/iam/service"
	post_repository "github.com/DigiConvent/testd9t/pkg/post/repository"
	post_service "github.com/DigiConvent/testd9t/pkg/post/service"
	sys_repository "github.com/DigiConvent/testd9t/pkg/sys/repository"
	sys_service "github.com/DigiConvent/testd9t/pkg/sys/service"
)

type Services struct {
	IAMService  iam_service.IAMServiceInterface
	SysService  sys_service.SysServiceInterface
	PostService post_service.PostServiceInterface
}

func InitiateServices(live bool) *Services {
	sysDB := db.NewSqliteDB("sys")
	sysRepo := sys_repository.NewSysRepository(sysDB)
	sysService := sys_service.NewSysService(sysRepo)
	sysService.Init()

	keyPath := "/home/testd9t/certs/privkey.pem"
	iamDB := db.NewSqliteDB("iam")
	iamRepo := iam_repository.NewIAMRepository(iamDB, keyPath, live)
	iamService := iam_service.NewIAMService(iamRepo)

	postDB := db.NewSqliteDB("post")
	postRepo := post_repository.NewPostRepository(postDB)
	postService := post_service.NewPostService(postRepo, live)

	// this needs to be called after all the databases are initialised as it migrates only packages of which a database file exists
	sysService.MigrateDatabase(nil)

	return &Services{
		SysService:  sysService,
		IAMService:  iamService,
		PostService: postService,
	}
}
