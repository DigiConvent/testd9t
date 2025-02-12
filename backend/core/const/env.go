package constants

import (
	"os"

	"github.com/DigiConvent/testd9t/core/log"
)

const DOMAIN = "DOMAIN"
const DATABASE_PATH = "DATABASE_PATH"
const CERTIFICATES_PATH = "CERTIFICATES_PATH"
const SMTP_PORT = "SMTP_PORT"
const HTTP_PORT = "HTTP_PORT"
const HTTPS_PORT = "HTTPS_PORT"

const HOME_PATH = "/home/testd9t/"
const ENV_PATH = HOME_PATH + "env"

func CheckEnv() {
	shouldContinue := true
	shouldContinue = CheckIfSet(DOMAIN, false, "") && shouldContinue
	shouldContinue = CheckIfSet(DATABASE_PATH, false, "") && shouldContinue
	shouldContinue = CheckIfSet(CERTIFICATES_PATH, false, "") && shouldContinue
	shouldContinue = CheckIfSet(SMTP_PORT, true, "465") && shouldContinue
	shouldContinue = CheckIfSet(HTTP_PORT, true, "80") && shouldContinue
	shouldContinue = CheckIfSet(HTTPS_PORT, true, "443") && shouldContinue

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
