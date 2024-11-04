package env_test

import (
	"os"
	"testing"

	"github.com/DigiConvent/testd9t/core/env"
)

func TestReplaceEnvVar(t *testing.T) {
	envVar := "TEST_ENV_VAR"
	envVal := "test"

	err := os.Setenv(envVar, envVal)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	str := "This is a test string ${TEST_ENV_VAR} a test environment variable: ${TEST_ENV_VAR}"
	expected := "This is a test string test a test environment variable: test"

	result := env.ReplaceEnvVar(str)

	if result != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
