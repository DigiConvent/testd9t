package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DigiConvent/testd9t/cli"
	"github.com/DigiConvent/testd9t/core/db"
	"github.com/DigiConvent/testd9t/core/log"
	core_utils "github.com/DigiConvent/testd9t/core/utils"
	router "github.com/DigiConvent/testd9t/delivery"
	services "github.com/DigiConvent/testd9t/pkg"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	sys_domain.StartTime = time.Now()
	log.Info("Build      " + sys_domain.CompiledAt)
	log.Info("Running at " + sys_domain.StartTime.Format(core_utils.FormattedTime))
	if sys_domain.ProgramVersion == "dev" {
		godotenv.Load("env")
		log.SetLogLevel(4)
	} else {
		gin.SetMode(gin.ReleaseMode)
		log.SetLogLevel(0)
		db.DatabasePath = os.Getenv("DATABASE_PATH")
		err := godotenv.Load("/home/testd9t/env")
		if err != nil {
			log.Error(err.Error())
		}
	}

	services := services.InitiateServices()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\nReceived shutdown signal, closing DB...")
		db.CloseAllDatabases()
		os.Exit(0)
	}()

	cli.HandleFlags(services.SysService)

	router.SetupRouter(services)

	<-sigChan
	log.Info("Closing DigiConvent")
}
