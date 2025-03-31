package tenants

import (
	"encoding/json"
	"time"

	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/pagination"
)

type tenantResult struct {
	eclcloud.Result
}

// GetResult is the result of a Get request. Call its Extract method to
// interpret it as a Tenant.
type GetResult struct {
	tenantResult
}

// CreateResult is the result of a Create request. Call its Extract method to
// interpret it as a Tenant.
type CreateResult struct {
	tenantResult
}

// Tenant represents an ECL SSS Tenant.
type Tenant struct {
	// 	ID of contract which owns these tenants.
	ContractID string `json:"contract_id"`
	// ID is the unique ID of the tenant.
	TenantID string `json:"tenant_id"`
	// Name of the tenant.
	TenantName string `json:"tenant_name"`
	// Description of the tenant.
	Description string `json:"description"`
	// TenantRegion the tenant belongs.
	TenantRegion string `json:"region"`
	// Time that the tenant is created.
	StartTime time.Time `json:"-"`
	// SSS API endpoint for the region.
	RegionApiEndpoint string `json:"region_api_endpoint"`
	// Users information who have access to this tenant.
	User []User `json:"users"`
	// Brand ID which this tenant belongs. (ex. ecl2)
	BrandID string `json:"brand_id"`
	// Workspace ID of the tenant.
	WorkspaceID string `json:"workspace_id"`
}

type User struct {
	// ID of the users who have access to this tenant.
	UserID string `json:"user_id"`
	// Contract which owns the tenant.
	ContractID string `json:"contract_id"`
	// This user is contract owner / or not.
	ContractOwner bool `json:"contract_owner"`
}

// UnmarshalJSON creates JSON format of tenant
func (r *Tenant) UnmarshalJSON(b []byte) error {
	type tmp Tenant
	var s struct {
		tmp
		StartTime eclcloud.JSONRFC3339ZNoTNoZ `json:"start_time"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Tenant(s.tmp)

	r.StartTime = time.Time(s.StartTime)

	return err
}

// TenantPage is a single page of Tenant results.
type TenantPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of Tenants contains any results.
func (r TenantPage) IsEmpty() (bool, error) {
	tenants, err := ExtractTenants(r)
	return len(tenants) == 0, err
}

// NextPageURL extracts the "next" link from the links section of the result.
func (r TenantPage) NextPageURL() (string, error) {
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

// ExtractTenants returns a slice of Tenants contained in a single page of
// results.
func ExtractTenants(r pagination.Page) ([]Tenant, error) {
	var s struct {
		ContractID string   `json:"contract_id"`
		Tenants    []Tenant `json:"tenants"`
	}

	// In list response case, each json element does not have contract_id.
	// It is set at out layer of each element.
	// So following logic set contract_id into inside of tenants slice forcibly.
	// In "show(get with ID of tenant)" case, this does not occur.
	err := (r.(TenantPage)).ExtractInto(&s)
	contractID := s.ContractID

	for i := 0; i < len(s.Tenants); i++ {
		s.Tenants[i].ContractID = contractID
	}
	return s.Tenants, err
}

// Extract interprets any projectResults as a Tenant.
func (r tenantResult) Extract() (*Tenant, error) {
	var s *Tenant
	err := r.ExtractInto(&s)
	return s, err
}
