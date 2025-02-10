package post_service_test

import "testing"

func TestListEmailAddresses(t *testing.T) {
	postService := GetTestPostService("post")

	addresses, status := postService.ListEmailAddresses()
	if status.Err() {
		t.Errorf("Failed to list email addresses: %s", status.Message)
	}
	if addresses == nil {
		t.Errorf("Failed to list email addresses: %s", status.Message)
	}
}
