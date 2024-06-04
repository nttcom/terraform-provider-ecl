package tenants

import (
	"encoding/json"
	"time"

	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
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

// DeleteResult is the result of a Delete request. Call its ExtractErr method to
// determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// UpdateResult is the result of an Update request. Call its Extract method to
// interpret it as a Tenant.
type UpdateResult struct {
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
	// TenantRegion the tenant blongs.
	TenantRegion string `json:"region"`
	// Time that the tenant is created.
	StartTime time.Time `json:"-"`
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
	// In "show(get with ID of tennat)" case, this does not occur.
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
