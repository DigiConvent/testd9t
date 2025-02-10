package post_service_test

import (
	"testing"

	post_domain "github.com/DigiConvent/testd9t/pkg/post/domain"
)

func TestUpdateEmailAddresses(t *testing.T) {
	postService := GetTestPostService("post")

	id, _ := postService.CreateEmailAddress(&post_domain.EmailAddressWrite{
		Name:   "TestUpdateEmailAddress",
		Domain: "test.test",
	})

	status := postService.UpdateEmailAddresses(id, &post_domain.EmailAddressWrite{
		Name:   "TestUpdateEmailAddress2",
		Domain: "test.test3",
	})
	if status.Err() {
		t.Errorf("Failed to update email address: %s", status.Message)
	}
	if status.Code != 204 {
		t.Errorf("Failed to update email address: %s", status.Message)
	}

	updatedAddress, _ := postService.ReadEmailAddress(id)

	if updatedAddress.Name != "testupdateemailaddress2" {
		t.Errorf("Failed to update email address: %s", status.Message)
	}
	if updatedAddress.Domain != "test.test3" {
		t.Errorf("Failed to update email address: %s", status.Message)
	}
}
