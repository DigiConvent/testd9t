package packages

import (
	"os"

	constants "github.com/DigiConvent/testd9t/core/const"
	"github.com/DigiConvent/testd9t/core/db"
	"github.com/DigiConvent/testd9t/core/log"
	iam_repository "github.com/DigiConvent/testd9t/pkg/iam/repository"
	iam_service "github.com/DigiConvent/testd9t/pkg/iam/service"
	iam_setup "github.com/DigiConvent/testd9t/pkg/iam/setup"
	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
	post_repository "github.com/DigiConvent/testd9t/pkg/post/repository"
	post_service "github.com/DigiConvent/testd9t/pkg/post/service"
	post_setup "github.com/DigiConvent/testd9t/pkg/post/setup"
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
	initStatus := sysService.Init()

	iam_setup.Setup()
	iamDB := db.NewSqliteDB("iam")
	iamRepo := iam_repository.NewIAMRepository(iamDB)
	iamService := iam_service.NewIAMService(iamRepo)

	post_setup.Setup()
	postDB := db.NewSqliteDB("post")
	postRepo := post_repository.NewPostRepository(postDB)
	postService := post_service.NewPostService(postRepo, live)

	services := &Services{
		SysService:  sysService,
		IAMService:  iamService,
		PostService: postService,
	}

	if initStatus.Code == 200 {
		log.Info("Migrating databases to newest version")
		status := sysService.MigratePackageDatabases(nil)
		if status.Err() {
			log.Error("Could not migrate databases to newest version: " + status.Message)
		}
		DoFirstTimeStuff(services)
	} else {
		log.Warning("Could not migrate system database to newest version")
	}

	return services
}

func DoFirstTimeStuff(services *Services) {
	log.Info("Doing first time stuff")
	emailAddress := os.Getenv("EMAIL")
	sendFrom, status := services.PostService.CreateEmailAddress(&post_domain.EmailAddressWrite{
		Name:   "admin",
		Domain: os.Getenv(constants.DOMAIN),
	})

	if status.Err() {
		log.Error("Could not create email admin address: " + status.Message)
	} else {
		log.Info("Created email admin address")
	}

	status = services.PostService.SendEmail(sendFrom, emailAddress, "Login credentials", "Here are the login credentials for "+os.Getenv(constants.DOMAIN)+":\n\nEmail: "+emailAddress+"\nPassword: "+os.Getenv(constants.MASTER_PASSWORD))

	if status.Err() {
		log.Error("Could not send email: from " + sendFrom.String() + " to " + emailAddress + ": " + status.Message)
	} else {
		log.Success("Sent password to " + emailAddress)
	}
}
