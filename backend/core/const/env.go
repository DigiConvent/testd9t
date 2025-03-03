package constants

import (
	"os"
	"sort"
	"strings"

	"github.com/DigiConvent/testd9t/core/log"
)

const CERTIFICATES_PATH = "CERTIFICATES_PATH"
const DATABASE_PATH = "DATABASE_PATH"
const DOMAIN = "DOMAIN"
const HTTP_PORT = "HTTP_PORT"
const HTTPS_PORT = "HTTPS_PORT"
const MASTER_PASSWORD = "PASSWORD"
const MASTER_EMAILADDRESS = "EMAILADDRESS"
const SMTP_PORT = "SMTP_PORT"
const TELEGRAM_BOT_TOKEN = "TELEGRAM_BOT_TOKEN"

const HOME_PATH = "/home/testd9t/"
const ENV_PATH = HOME_PATH + "env"

func CheckEnv() {
	shouldContinue := true
	shouldContinue = CheckIfSet(CERTIFICATES_PATH, false, "") && shouldContinue
	shouldContinue = CheckIfSet(DATABASE_PATH, false, "") && shouldContinue
	shouldContinue = CheckIfSet(DOMAIN, false, "") && shouldContinue
	shouldContinue = CheckIfSet(HTTP_PORT, true, "80") && shouldContinue
	shouldContinue = CheckIfSet(HTTPS_PORT, true, "443") && shouldContinue
	shouldContinue = CheckIfSet(SMTP_PORT, true, "587") && shouldContinue

	shouldContinue = CheckIfSet(MASTER_PASSWORD, false, "") && shouldContinue
	shouldContinue = CheckIfSet(MASTER_EMAILADDRESS, false, "") && shouldContinue

	shouldContinue = CheckIfSet(TELEGRAM_BOT_TOKEN, true, "") && shouldContinue

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

func SetEnvVariable(key string, value string) {
	os.Setenv(key, value)
	contents, err := os.ReadFile(ENV_PATH)
	if err != nil {
		panic(err)
	}

	pairs := strings.Split(string(contents), "\n")
	for i := range pairs {
		pair := strings.Split(pairs[i], "=")
		if pair[0] == key {
			pairs[i] = key + "=" + value
		}
	}

	sort.Strings(pairs)

	contents = []byte(strings.Join(pairs, "\n"))
	if contents[len(contents)-1] != '\n' {
		contents = append(contents, '\n')
	}

	err = os.WriteFile(ENV_PATH, contents, 0644)
	if err != nil {
		panic(err)
	}
}
