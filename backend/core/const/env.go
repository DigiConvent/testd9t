package constants

import (
	"os"

	"github.com/DigiConvent/testd9t/core/log"
)

const DOMAIN = "DOMAIN"
const DATABASE_PATH = "DATABASE_PATH"
const CERTIFICATES_PATH = "CERTIFICATES_PATH"

func CheckEnv() {
	shouldContinue := true
	if os.Getenv(DOMAIN) == "" {
		log.Error("DOMAIN environment variable is not set")
		shouldContinue = false
	}

	if os.Getenv(DATABASE_PATH) == "" {
		log.Error("DATABASE_PATH environment variable is not set")
		shouldContinue = false
	}

	if os.Getenv(CERTIFICATES_PATH) == "" {
		log.Error("CERTIFICATES_PATH environment variable is not set")
		shouldContinue = false
	}

	if !shouldContinue {
		os.Exit(1)
	}
}
