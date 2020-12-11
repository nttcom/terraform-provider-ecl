package tenant_connection_requests

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

// TenantConnectionRequest represents Tenant Connection Request.
type TenantConnectionRequest struct {
	ID                string            `json:"id"`
	Status            string            `json:"status"`
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	Tags              map[string]string `json:"tags"`
	TenantID          string            `json:"tenant_id"`
	NameOther         string            `json:"name_other"`
	DescriptionOther  string            `json:"description_other"`
	TagsOther         map[string]string `json:"tags_other"`
	TenantIDOther     string            `json:"tenant_id_other"`
	NetworkID         string            `json:"network_id"`
	ApprovalRequestID string            `json:"approval_request_id"`
}

type commonResult struct {
	eclcloud.Result
}

// GetResult is the response from a Get operation. Call its Extract method
// to interpret it as a Tenant Connection Request.
type GetResult struct {
	commonResult
}

// CreateResult is the response from a Create operation. Call its Extract method
// to interpret it as a Tenant Connection Request.
type CreateResult struct {
	commonResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr to
// determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// UpdateResult is the result of an Update request. Call its Extract method to
// interpret it as a Tenant Connection Request.
type UpdateResult struct {
	commonResult
}

// TenantConnectionRequestPage is a single page of Tenant Connection Request results.
type TenantConnectionRequestPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of Tenant Connection Request contains any results.
func (r TenantConnectionRequestPage) IsEmpty() (bool, error) {
	resources, err := ExtractTenantConnectionRequests(r)
	return len(resources) == 0, err
}

// ExtractTenantConnectionRequests returns a slice of Tenant Connection Requests contained in a
// single page of results.
func ExtractTenantConnectionRequests(r pagination.Page) ([]TenantConnectionRequest, error) {
	var s struct {
		TenantConnectionRequest []TenantConnectionRequest `json:"tenant_connection_requests"`
	}
	err := (r.(TenantConnectionRequestPage)).ExtractInto(&s)
	return s.TenantConnectionRequest, err
}

// Extract interprets any commonResult as a Tenant Connection Request.
func (r commonResult) Extract() (*TenantConnectionRequest, error) {
	var s struct {
		TenantConnectionRequest *TenantConnectionRequest `json:"tenant_connection_request"`
	}
	err := r.ExtractInto(&s)
	return s.TenantConnectionRequest, err
}
