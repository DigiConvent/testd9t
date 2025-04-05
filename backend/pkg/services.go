package packages

import (
	"os"

	constants "github.com/DigiConvent/testd9t/core/const"
	"github.com/DigiConvent/testd9t/core/db"
	"github.com/DigiConvent/testd9t/core/log"
	iam_domain "github.com/DigiConvent/testd9t/pkg/iam/domain"
	iam_repository "github.com/DigiConvent/testd9t/pkg/iam/repository"
	iam_service "github.com/DigiConvent/testd9t/pkg/iam/service"
	iam_setup "github.com/DigiConvent/testd9t/pkg/iam/setup"
	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
	post_repository "github.com/DigiConvent/testd9t/pkg/post/repository"
	post_service "github.com/DigiConvent/testd9t/pkg/post/service"
	post_setup "github.com/DigiConvent/testd9t/pkg/post/setup"
	sys_repository "github.com/DigiConvent/testd9t/pkg/sys/repository"
	sys_service "github.com/DigiConvent/testd9t/pkg/sys/service"
	"github.com/google/uuid"
)

type Services struct {
	IamService  iam_service.IAMServiceInterface
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
	iamRepo := iam_repository.NewIamRepository(iamDB)
	iamService := iam_service.NewIamService(iamRepo)

	post_setup.Setup()
	postDB := db.NewSqliteDB("post")
	postRepo := post_repository.NewPostRepository(postDB)
	postService := post_service.NewPostService(postRepo, live)

	services := &Services{
		SysService:  sysService,
		IamService:  iamService,
		PostService: postService,
	}

	if initStatus.Code == 200 {
		log.Info("Migrating databases to newest version")
		status := sysService.MigratePackageDatabases(nil)
		if status.Err() {
			log.Error("Could not migrate system database to newest version: " + status.Message)
		}
		DoFirstTimeStuff(services)
	} else {
		log.Warning("Could not migrate system database to newest version: " + initStatus.Message)
	}

	return services
}

// initialise the things that sqlite can't do on its own
func DoFirstTimeStuff(services *Services) {
	services.SysService.SetBotToken(os.Getenv(constants.TELEGRAM_BOT_TOKEN))
	emailAddress := os.Getenv(constants.MASTER_EMAILADDRESS)
	id := uuid.MustParse("00000000-0000-0000-0000-000000000000")
	status := services.PostService.UpdateEmailAddresses(&id, &post_domain.EmailAddressWrite{
		Name:   "admin",
		Domain: os.Getenv(constants.DOMAIN),
	})

	if status.Err() {
		log.Error("Could not update email admin address: " + status.Message)
	}

	status = services.IamService.SetEnabled(&id, true)
	if status.Err() {
		log.Error("Could not enable super user: " + status.Message)
	}

	status = services.IamService.UpdateUser(&id, &iam_domain.UserWrite{
		Emailaddress: emailAddress,
		FirstName:    "admin",
		LastName:     "admin",
	})

	if status.Err() {
		log.Error("Could not update super user: " + status.Message)
	}

	status = services.IamService.SetUserPassword(&id, os.Getenv(constants.MASTER_PASSWORD))
	if status.Err() {
		log.Error("Could not set super user password: " + status.Message)
	}

	// create a root permission_group
	rootId, status := services.IamService.CreatePermissionGroup(&iam_domain.PermissionGroupWrite{
		Name:        "root",
		Abbr:        "root",
		IsGroup:     false,
		IsNode:      false,
		Description: "",
	})

	if status.Err() {
		log.Error("Could not create root permission group: " + status.Message)
	}

	status = services.IamService.SetParentPermissionGroup(&iam_domain.PermissionGroupSetParent{
		ID:     &id,
		Parent: rootId,
	})

	if status.Err() {
		log.Error("Could not set super user as root permission group: " + status.Message)
	}
}
