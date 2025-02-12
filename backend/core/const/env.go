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

	if os.Getenv("SMTP_PORT") == "" {
		log.Info("SMTP_PORT environment variable is not set, defaulting to 465")
		os.Setenv("SMTP_PORT", "465")
	}

	if os.Getenv("HTTPS_PORT") == "" {
		log.Info("HTTPS_PORT environment variable is not set, defaulting to 443")
		os.Setenv("HTTPS_PORT", "443")
	}

	if os.Getenv("HTTP_PORT") == "" {
		log.Info("HTTP_PORT environment variable is not set, defaulting to 80")
		os.Setenv("HTTP_PORT", "80")
	}

	if !shouldContinue {
		os.Exit(1)
	}
}
