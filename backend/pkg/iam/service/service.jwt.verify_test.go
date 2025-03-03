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
	token, status := iamService.GenerateJwt(&someId)
	if status.Err() {
		t.Fatal(status.Message)
	}

	theId, status := iamService.VerifyJwt(token)
	if status.Err() {
		t.Fatal(status.Message)
	}

	if theId == nil {
		t.Fatal("Expected a result")
	}

	if theId.String() != someId.String() {
		t.Fatal("Expected ", someId.String(), " instead I got ", theId.String())
	}
}
