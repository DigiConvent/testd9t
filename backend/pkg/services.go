package services

import (
	"os"

	"github.com/DigiConvent/testd9t/core/db"
	"github.com/DigiConvent/testd9t/core/log"
	iam_repository "github.com/DigiConvent/testd9t/pkg/iam/repository"
	iam_service "github.com/DigiConvent/testd9t/pkg/iam/service"
	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
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

func InitiateServices() *Services {
	sysDB := db.NewSqliteDB("sys")
	sysRepo := sys_repository.NewSysRepository(sysDB)
	sysService := sys_service.NewSysService(sysRepo)
	initStatus := sysService.Init()

	keyPath := "/home/testd9t/certs/privkey.pem"
	iamDB := db.NewSqliteDB("iam")
	iamRepo := iam_repository.NewIAMRepository(iamDB, keyPath, false)
	iamService := iam_service.NewIAMService(iamRepo)

	postDB := db.NewSqliteDB("post")
	postRepo := post_repository.NewPostRepository(postDB)
	postService := post_service.NewPostService(postRepo, false)

	services := &Services{
		SysService:  sysService,
		IAMService:  iamService,
		PostService: postService,
	}
	if initStatus.Code == 200 {
		log.Info("Migrating databases to newest version")
		sysService.MigratePackageDatabases(nil)
		DoFirstTimeStuff(services)
	}
	return services
}

func DoFirstTimeStuff(services *Services) {
	log.Success("First time setup")
	sendFrom, status := services.PostService.CreateEmailAddress(&post_domain.EmailAddressWrite{
		Name:   "admin",
		Domain: os.Getenv("DOMAIN"),
	})

	if !status.Err() {
		log.Error(status.Message)
	} else {
		log.Success(status.Message)
	}

	status = services.PostService.SendEmail(sendFrom, os.Getenv("EMAIL"), "Login credentials", "Here are the login credentials for "+os.Getenv("DOMAIN")+":\n\nEmail: "+os.Getenv("EMAIL")+"\nPassword: "+os.Getenv("PASSWORD"))

	if !status.Err() {
		log.Error(status.Message)
	} else {
		log.Success(status.Message)
	}
}
