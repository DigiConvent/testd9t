package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DigiConvent/testd9t/cli"
	constants "github.com/DigiConvent/testd9t/core/const"
	"github.com/DigiConvent/testd9t/core/db"
	"github.com/DigiConvent/testd9t/core/log"
	core_utils "github.com/DigiConvent/testd9t/core/utils"
	router "github.com/DigiConvent/testd9t/delivery"
	packages "github.com/DigiConvent/testd9t/pkg"
	sys_domain "github.com/DigiConvent/testd9t/pkg/sys/domain"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	sys_domain.StartTime = time.Now()
	log.Info("Build      " + sys_domain.CompiledAt)
	log.Info("Running at " + sys_domain.StartTime.Format(core_utils.FormattedTime))
	run := core_utils.Contains(os.Args, "--run")
	local := sys_domain.ProgramVersion == "dev"

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\nReceived shutdown signal, closing DB...")
		db.CloseAllDatabases()
		os.Exit(0)
	}()

	if run {
		services := packages.InitiateServices(!local)
		router.SetupRouter(services)
	} else {
		cli.HandleFlags()
	}

	<-sigChan
	log.Info("Closing DigiConvent")
}

func loadEnv() {
	if sys_domain.ProgramVersion == "dev" {
		err := godotenv.Load("env")
		if err != nil {
			log.Error(err.Error())
		}
	}
	constants.CheckEnv()
}
