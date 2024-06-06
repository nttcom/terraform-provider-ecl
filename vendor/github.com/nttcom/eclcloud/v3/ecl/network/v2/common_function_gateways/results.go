package common_function_gateways

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// Extract is a function that accepts a result
// and extracts a common function gateway resource.
func (r commonResult) Extract() (*CommonFunctionGateway, error) {
	var cfgw CommonFunctionGateway
	err := r.ExtractInto(&cfgw)
	return &cfgw, err
}

// Extract interprets any commonResult as a Common Function Gateway, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "common_function_gateway")
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Common Function Gateway.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Common Function Gateway.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Common Function Gateway.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// CommonFunctionGateway represents, well, a common function gateway.
type CommonFunctionGateway struct {
	// UUID for the network
	ID string `json:"id"`

	// Human-readable name for the network. Might not be unique.
	Name string `json:"name"`

	Description string `json:"description"`

	CommonFunctionPoolID string `json:"common_function_pool_id"`

	NetworkID string `json:"network_id"`

	SubnetID string `json:"subnet_id"`
	Status   string `json:"status"`
	TenantID string `json:"tenant_id"`
}

// CommonFunctionGatewayPage is the page returned by a pager
// when traversing over a collection of common function gateway.
type CommonFunctionGatewayPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of common function gateway
//  has reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r CommonFunctionGatewayPage) NextPageURL() (string, error) {
	var s struct {
		Links []eclcloud.Link `json:"common_function_gateways_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return eclcloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a CommonFunctionGatewayPage struct is empty.
func (r CommonFunctionGatewayPage) IsEmpty() (bool, error) {
	is, err := ExtractCommonFunctionGateways(r)
	return len(is) == 0, err
}

// ExtractCommonFunctionGateways accepts a Page struct,
// specifically a NetworkPage struct, and extracts the elements
// into a slice of Common Function Gateway structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractCommonFunctionGateways(r pagination.Page) ([]CommonFunctionGateway, error) {
	var s []CommonFunctionGateway
	err := ExtractCommonFunctionGatewaysInto(r, &s)
	return s, err
}

// ExtractCommonFunctionGatewaysInto interprets the results of a single page from a List() call,
// producing a slice of Server entities.
func ExtractCommonFunctionGatewaysInto(r pagination.Page, v interface{}) error {
	return r.(CommonFunctionGatewayPage).Result.ExtractIntoSlicePtr(v, "common_function_gateways")
}
