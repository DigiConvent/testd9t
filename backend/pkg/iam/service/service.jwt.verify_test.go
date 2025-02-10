package iam_service_test

import (
	"testing"

	"github.com/google/uuid"
)

func TestVerifyJwt(t *testing.T) {
	iamService := GetTestIAMService("iam")

	_, status := iamService.VerifyJwt("badtoken")
	if !status.Err() {
		t.Fatal("Expected an error")
	}

	_, status = iamService.VerifyJwt("")
	if !status.Err() {
		t.Fatal("Expected an error")
	}

	someId, _ := uuid.NewV7()
	token, _ := iamService.GenerateJwt(&someId)

	_, status = iamService.VerifyJwt(token)
	if status.Err() {
		t.Fatal(status.Message)
	}

}
