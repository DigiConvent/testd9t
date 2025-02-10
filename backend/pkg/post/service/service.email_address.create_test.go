package post_service_test

import (
	"testing"

	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
)

func TestCreateEmailAddress(t *testing.T) {
	postService := GetTestPostService("post")

	credentials := &post_domain.EmailAddressWrite{
		Name:   "TestCreateEmailAddress",
		Domain: "test.test",
	}

	id, status := postService.CreateEmailAddress(credentials)

	if status.Err() {
		t.Errorf("Failed to create email address: %s", status.Message)
	}
	if id == nil {
		t.Errorf("Failed to create email address: %s", status.Message)
	}
}
