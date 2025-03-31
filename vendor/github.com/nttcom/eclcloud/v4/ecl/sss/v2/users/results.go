package users

import (
	"encoding/json"
	"time"

	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/pagination"
)

type userResult struct {
	eclcloud.Result
}

// GetResult is the result of a Get request. Call its Extract method to
// interpret it as a User.
type GetResult struct {
	userResult
}

// CreateResult is the result of a Create request. Call its Extract method to
// interpret it as a User.
type CreateResult struct {
	userResult
}

// DeleteResult is the result of a Delete request. Call its ExtractErr method to
// determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// UpdateResult is the result of an Update request. Call its Extract method to
// interpret it as a User.
type UpdateResult struct {
	userResult
}

// User represents an ECL SSS User.
type User struct {
	LoginID             string    `json:"login_id"`
	MailAddress         string    `json:"mail_address"`
	UserID              string    `json:"user_id"`
	ContractOwner       bool      `json:"contract_owner"`
	Superuser           bool      `json:"super_user"`
	ApiAvailability     bool      `json:"api_availability"`
	KeystoneName        string    `json:"keystone_name"`
	KeystoneEndpoint    string    `json:"keystone_endpoint"`
	SSSEndpoint         string    `json:"sss_endpoint"`
	ContractID          string    `json:"contract_id"`
	LoginIntegration    string    `json:"login_integration"`
	ExternalReferenceID string    `json:"external_reference_id"`
	BrandID             string    `json:"brand_id"`
	OtpActivation       bool      `json:"otp_activation"`
	StartTime           time.Time `json:"-"`
}

// UnmarshalJSON creates JSON format of user
func (r *User) UnmarshalJSON(b []byte) error {
	type tmp User
	var s struct {
		tmp
		StartTime eclcloud.JSONRFC3339ZNoTNoZ `json:"start_time"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = User(s.tmp)

	r.StartTime = time.Time(s.StartTime)

	return err
}

// UserPage is a single page of User results.
type UserPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of User contains any results.
func (r UserPage) IsEmpty() (bool, error) {
	users, err := ExtractUsers(r)
	return len(users) == 0, err
}

// NextPageURL extracts the "next" link from the links section of the result.
func (r UserPage) NextPageURL() (string, error) {
	var s struct {
		Links struct {
			Next     string `json:"next"`
			Previous string `json:"previous"`
		} `json:"links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return s.Links.Next, err
}

// ExtractUsers returns a slice of Users contained in a single page of
// results.
func ExtractUsers(r pagination.Page) ([]User, error) {
	var s struct {
		ContractID string `json:"contract_id"`
		Users      []User `json:"users"`
	}

	// In list response case, each json element does not have contract_id.
	// It is set at out layer of each element.
	// So following logic set contract_id into inside of users slice forcibly.
	// In "show(get with ID of tenant)" case, this does not occur.
	err := (r.(UserPage)).ExtractInto(&s)
	contractID := s.ContractID

	for i := 0; i < len(s.Users); i++ {
		s.Users[i].ContractID = contractID
	}
	return s.Users, err
}

// Extract interprets any projectResults as a User.
func (r userResult) Extract() (*User, error) {
	var u *User
	err := r.ExtractInto(&u)
	return u, err
}
