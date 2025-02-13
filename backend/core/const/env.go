package constants

import (
	"os"

	"github.com/DigiConvent/testd9t/core/log"
)

const CERTIFICATES_PATH = "CERTIFICATES_PATH"
const DATABASE_PATH = "DATABASE_PATH"
const DOMAIN = "DOMAIN"
const HTTP_PORT = "HTTP_PORT"
const HTTPS_PORT = "HTTPS_PORT"
const MASTER_PASSWORD = "PASSWORD"
const SMTP_PORT = "SMTP_PORT"

const ENV_PATH = HOME_PATH + "env"
const HOME_PATH = "/home/testd9t/"

func CheckEnv() {
	shouldContinue := true
	shouldContinue = CheckIfSet(CERTIFICATES_PATH, false, "") && shouldContinue
	shouldContinue = CheckIfSet(DATABASE_PATH, false, "") && shouldContinue
	shouldContinue = CheckIfSet(DOMAIN, false, "") && shouldContinue
	shouldContinue = CheckIfSet(HTTP_PORT, true, "80") && shouldContinue
	shouldContinue = CheckIfSet(HTTPS_PORT, true, "443") && shouldContinue
	shouldContinue = CheckIfSet(SMTP_PORT, true, "587") && shouldContinue

	if !shouldContinue {
		os.Exit(1)
	}
}

func CheckIfSet(key string, optional bool, def string) bool {
	if os.Getenv(key) == "" {
		if !optional {
			log.Error(key + " environment variable is not set")
			return false
		} else {
			log.Info(key + " environment variable is not set, defaulting to " + def)
		}
	}
	return true
}
