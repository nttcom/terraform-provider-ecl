package licenses

import (
	"time"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

// License represents guest image license key information.
type License struct {
	ID           string     `json:"id"`
	Key          string     `json:"key"`
	AssignedFrom time.Time  `json:"assigned_from"`
	ExpiresAt    *time.Time `json:"expires_at"`
	LicenseType  string     `json:"license_type"`
}

type commonResult struct {
	eclcloud.Result
}

// CreateResult is the response from a Create operation. Call its Extract method
// to interpret it as a License.
type CreateResult struct {
	commonResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr to
// determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// LicensePage is a single page of License results.
type LicensePage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of Licenses contains any results.
func (r LicensePage) IsEmpty() (bool, error) {
	licenses, err := ExtractLicenses(r)
	return len(licenses) == 0, err
}

// ExtractLicenses returns a slice of Licenses contained in a single page of
// results.
func ExtractLicenses(r pagination.Page) ([]License, error) {
	var s struct {
		Licenses []License `json:"licenses"`
	}
	err := (r.(LicensePage)).ExtractInto(&s)
	return s.Licenses, err
}

// ExtractLicenseInfo interprets any commonResult as a License.
func (r commonResult) ExtractLicenseInfo() (*License, error) {
	var s struct {
		License *License `json:"license"`
	}
	err := r.ExtractInto(&s)
	return s.License, err
}
