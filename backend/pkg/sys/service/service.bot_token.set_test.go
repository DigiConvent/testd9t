package sys_service_test

import (
	"os"
	"testing"

	constants "github.com/DigiConvent/testd9t/core/const"
)

func TestSetBotToken(t *testing.T) {
	sysService := GetTestSysService("sys")

	if os.Getenv(constants.TELEGRAM_BOT_TOKEN) != "" {
		t.Fatal("Expected no bot token")
	}

	status := sysService.SetBotToken("thisissomebottoken")

	if status.Err() {
		t.Fatalf("Error: %v", status)
	}

	configuration, status := sysService.GetConfiguration()

	if status.Err() {
		t.Fatalf("Error: %v", status)
	}

	if configuration.TelegramBotToken != "thisissomebottoken" {
		t.Fatalf("Expected: thisissomebottoken, Got: %v", configuration.TelegramBotToken)
	}

	if os.Getenv(constants.TELEGRAM_BOT_TOKEN) != "thisissomebottoken" {
		t.Fatalf("Expected: thisissomebottoken, Got: %v", os.Getenv(constants.TELEGRAM_BOT_TOKEN))
	}
}
