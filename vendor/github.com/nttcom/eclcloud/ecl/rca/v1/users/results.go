package users

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

// User represents VPN user.
type User struct {
	Name         string        `json:"name"`
	Password     string        `json:"password"`
	VPNEndpoints []VPNEndpoint `json:"vpn_endpoints"`
}

// VPNEndpoint represents VPN Endpoint.
type VPNEndpoint struct {
	Endpoint string `json:"endpoint"`
	Type     string `json:"type"`
}

type commonResult struct {
	eclcloud.Result
}

// GetResult is the response from a Get operation. Call its Extract method
// to interpret it as a user.
type GetResult struct {
	commonResult
}

// CreateResult is the response from a Create operation. Call its Extract method
// to interpret it as a user.
type CreateResult struct {
	commonResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr to
// determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// UpdateResult is the result of an Update request. Call its Extract method to
// interpret it as a user.
type UpdateResult struct {
	commonResult
}

// UserPage is a single page of user results.
type UserPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of users contains any results.
func (r UserPage) IsEmpty() (bool, error) {
	resources, err := ExtractUsers(r)
	return len(resources) == 0, err
}

// ExtractUsers returns a slice of users contained in a single page of
// results.
func ExtractUsers(r pagination.Page) ([]User, error) {
	var s struct {
		Users []User `json:"users"`
	}
	err := (r.(UserPage)).ExtractInto(&s)
	return s.Users, err
}

// Extract interprets any commonResult as a user.
func (r commonResult) Extract() (*User, error) {
	var s struct {
		User *User `json:"user"`
	}
	err := r.ExtractInto(&s)
	return s.User, err
}
