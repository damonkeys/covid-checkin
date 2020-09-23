package applesupport

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestResolveWrongUserNameFromRequest(t *testing.T) {
	// idea from here: https: //www.reddit.com/r/golang/comments/6bortg/how_to_test_post_requests_with_data_other_than/
	t.Parallel()
	userJSON := `{"this_is_invalid"}`
	form := url.Values{}
	form.Add("user", userJSON)

	/* strings.NewReader(userJSON) */
	req, err := http.NewRequest("POST", "http://example.com", strings.NewReader(form.Encode()))
	req.Form = form
	if err != nil {
		t.Fatalf("Couldn't create request. Error: %s", err)
	}
	userNameData := &UserNameData{}
	err = userNameData.ResolveUserNames(req)

	if err == nil {
		t.Fatalf("Invalid json errorfree parsed into userNameData. Error: %s", err)
	}
}

func TestResolveGoodUserNameFromRequest(t *testing.T) {
	// idea from here: https: //www.reddit.com/r/golang/comments/6bortg/how_to_test_post_requests_with_data_other_than/
	t.Parallel()
	userJSON := `{"name":{"firstName":"first", "lastName": "last"}, "email":"foo@example.com"}`
	form := url.Values{}
	form.Add("user", userJSON)
	req, err := http.NewRequest("POST", "http://example.com", strings.NewReader(form.Encode()))
	req.Form = form

	userNameData := &UserNameData{}
	err = userNameData.ResolveUserNames(req)

	if err != nil {
		t.Fatalf("Couldn't create request. Error: %s", err)
	}
	if userNameData.FirstName != "first" {
		t.Fatalf("Wrong first name: %s", userNameData.FirstName)
	}

	if userNameData.LastName != "last" {
		t.Fatalf("Wrong first name: %s", userNameData.LastName)
	}
}
