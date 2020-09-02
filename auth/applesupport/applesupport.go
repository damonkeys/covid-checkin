package applesupport

import (
	"encoding/json"
	"net/http"
)

// UserNameData is the value of the form Parameter user that apple sends with the request - outside of the oauth ID data.
// This happens only on the first time an authorization callback is been made to goth from apple. Additionally it only happens
// if the right scope is set: apple.ScopeName.
type UserNameData struct {
	UserNames `json:"name"`
}

// UserNames is the container that holds the real user data within FormUser struct
type UserNames struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

//UserNameResolver is anything the fetches first and last name form request data
type UserNameResolver interface {
	ResolveUserNames(req *http.Request) error
}

// ResolveUserNames Resolves the usernames from form fields. A special case for "sign in with apple"
func (userNameData *UserNameData) ResolveUserNames(req *http.Request) error {
	userData := req.FormValue("user")
	if userData != "" {
		err := json.Unmarshal([]byte(req.FormValue("user")), userNameData)
		return err
	}
	return nil
}
