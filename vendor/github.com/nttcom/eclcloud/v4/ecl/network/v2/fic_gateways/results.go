package fic_gateways

import (
	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/pagination"
)

type FICGatewayPage struct {
	pagination.LinkedPageBase
}

type commonResult struct {
	eclcloud.Result
}

// GetResult is the result of Get operations. Call its Extract method to
// interpret it as a FICGateway.
type GetResult struct {
	commonResult
}

// FICGateway represents a FIC Gateway.
type FICGateway struct {
	Description  string `json:"description"`
	FICServiceID string `json:"fic_service_id"`
	ID           string `json:"id"`
	Name         string `json:"name"`
	QoSOptionID  string `json:"qos_option_id"`
	Status       string `json:"status"`
	TenantID     string `json:"tenant_id"`
}

// IsEmpty checks whether a FICGatewayPage struct is empty.
func (r FICGatewayPage) IsEmpty() (bool, error) {
	is, err := ExtractFICGateways(r)
	return len(is) == 0, err
}

// ExtractFICGateways accepts a Page struct, specifically a FICGatewayPage struct,
// and extracts the elements into a slice of ListOpts structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractFICGateways(r pagination.Page) ([]FICGateway, error) {
	var s []FICGateway
	err := r.(FICGatewayPage).Result.ExtractIntoSlicePtr(&s, "fic_gateways")
	return s, err
}

// Extract is a function that accepts a result and extracts a FICGateway.
func (r GetResult) Extract() (*FICGateway, error) {
	var l FICGateway
	err := r.Result.ExtractIntoStructPtr(&l, "fic_gateway")
	return &l, err
}
