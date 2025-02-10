package post_service_test

import (
	"testing"

	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
)

func TestDeleteEmailAddress(t *testing.T) {
	postService := GetTestPostService("post")

	id, _ := postService.CreateEmailAddress(&post_domain.EmailAddressWrite{
		Name:   "TestDeleteEmailAddress",
		Domain: "test.test",
	})

	status := postService.DeleteEmailAddress(id)
	if status.Err() {
		t.Errorf("Failed to delete email address: %s", status.Message)
	}

	if status.Code != 204 {
		t.Errorf("Failed to delete email address: %s", status.Message)
	}
}
