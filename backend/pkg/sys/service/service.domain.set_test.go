package sys_service_test

import (
	"os"
	"testing"

	constants "github.com/DigiConvent/testd9t/core/const"
)

func TestSetDomain(t *testing.T) {
	sysService := GetTestSysService("sys")

	if os.Getenv(constants.DOMAIN) != "" {
		t.Fatal("Expected no domain, instead got ", os.Getenv(constants.DOMAIN))
	}

	status := sysService.SetDomain("thisisadomain")

	if status.Err() {
		t.Fatalf("Error: %v", status)
	}

	configuration, status := sysService.GetConfiguration()

	if status.Err() {
		t.Fatalf("Error: %v", status)
	}

	if configuration.Domain != "thisisadomain" {
		t.Fatalf("Expected: thisisadomain, Got: %v", configuration.Domain)
	}

	if os.Getenv(constants.DOMAIN) != "thisisadomain" {
		t.Fatalf("Expected: thisisadomain, Got: %v", os.Getenv(constants.DOMAIN))
	}
}
