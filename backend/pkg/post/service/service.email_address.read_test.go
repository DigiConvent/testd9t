package post_service_test

import (
	"testing"

	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
)

func TestReadEmailAddress(t *testing.T) {
	postService := GetTestPostService("post")

	id, _ := postService.CreateEmailAddress(&post_domain.EmailAddressWrite{
		Name:   "TestReadEmailAddress",
		Domain: "test.test",
	})

	address, status := postService.ReadEmailAddress(id)
	if status.Err() {
		t.Errorf("Failed to read email address: %s", status.Message)
	}
	if address == nil {
		t.Fatalf("Failed to read email address: %s", status.Message)
	}

	if address.Name != "testreademailaddress" {
		t.Errorf("Failed to read email address: %s using id %s", status.Message, id)
	}
	if address.Domain != "test.test" {
		t.Errorf("Failed to read email address: %s", status.Message)
	}
}
