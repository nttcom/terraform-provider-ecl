package routes

import (
	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// CreateResult represents the result of a Create operation.
// Call its Extract method to interpret it as a Route.
type CreateResult struct {
	commonResult
}

// ShowResult represents the result of a Show operation.
// Call its Extract method to interpret it as a Route.
type ShowResult struct {
	commonResult
}

// UpdateResult represents the result of a Update operation.
// Call its Extract method to interpret it as a Route.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a Delete operation.
// Call its ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// CreateStagedResult represents the result of a CreateStaged operation.
// Call its Extract method to interpret it as a Route.
type CreateStagedResult struct {
	commonResult
}

// ShowStagedResult represents the result of a ShowStaged operation.
// Call its Extract method to interpret it as a Route.
type ShowStagedResult struct {
	commonResult
}

// UpdateStagedResult represents the result of a UpdateStaged operation.
// Call its Extract method to interpret it as a Route.
type UpdateStagedResult struct {
	commonResult
}

// CancelStagedResult represents the result of a CancelStaged operation.
// Call its ExtractErr method to determine if the request succeeded or failed.
type CancelStagedResult struct {
	eclcloud.ErrResult
}

// ConfigurationInResponse represents a configuration in a route.
type ConfigurationInResponse struct {

	// - IP address of next hop for the (static) route
	NextHopIPAddress string `json:"next_hop_ip_address,omitempty"`
}

// Route represents a route.
type Route struct {

	// - ID of the (static) route
	ID string `json:"id"`

	// - Name of the (static) route
	Name string `json:"name"`

	// - Description of the (static) route
	Description string `json:"description"`

	// - Tags of the (static) route (JSON object format)
	Tags map[string]interface{} `json:"tags"`

	// - Configuration status of the (static) route
	//   - `"ACTIVE"`
	//     - There are no configurations of the (static) route that waiting to be applied
	//   - `"CREATE_STAGED"`
	//     - The (static) route has been added and waiting to be applied
	//   - `"UPDATE_STAGED"`
	//     - Changed configurations of the (static) route exists that waiting to be applied
	//   - `"DELETE_STAGED"`
	//     - The (static) route has been removed and waiting to be applied
	ConfigurationStatus string `json:"configuration_status"`

	// - Operation status of the load balancer which the (static) route belongs to
	//   - `"NONE"` :
	//     - There are no operations of the load balancer
	//     - The load balancer and related resources can be operated
	//   - `"PROCESSING"`
	//     - The latest operation of the load balancer is processing
	//     - The load balancer and related resources cannot be operated
	//   - `"COMPLETE"`
	//     - The latest operation of the load balancer has been succeeded
	//     - The load balancer and related resources can be operated
	//   - `"STUCK"`
	//     - The latest operation of the load balancer has been stopped
	//     - Operators of NTT Communications will investigate the operation
	//     - The load balancer and related resources cannot be operated
	//   - `"ERROR"`
	//     - The latest operation of the load balancer has been failed
	//     - The operation was roll backed normally
	//     - The load balancer and related resources can be operated
	OperationStatus string `json:"operation_status"`

	// - CIDR of destination for the (static) route
	DestinationCidr string `json:"destination_cidr,omitempty"`

	// - ID of the load balancer which the (static) route belongs to
	LoadBalancerID string `json:"load_balancer_id"`

	// - ID of the owner tenant of the (static) route
	TenantID string `json:"tenant_id"`

	// - IP address of next hop for the (static) route
	NextHopIPAddress string `json:"next_hop_ip_address,omitempty"`

	// - Running configurations of the (static) route
	// - If `changes` is `true`, return object
	// - If current configuration does not exist, return `null`
	Current ConfigurationInResponse `json:"current,omitempty"`

	// - Added or changed configurations of the (static) route that waiting to be applied
	// - If `changes` is `true`, return object
	// - If staged configuration does not exist, return `null`
	Staged ConfigurationInResponse `json:"staged,omitempty"`
}

// ExtractInto interprets any commonResult as a route, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "route")
}

// Extract is a function that accepts a result and extracts a Route resource.
func (r commonResult) Extract() (*Route, error) {
	var route Route

	err := r.ExtractInto(&route)

	return &route, err
}

// RoutePage is the page returned by a pager when traversing over a collection of route.
type RoutePage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a RoutePage struct is empty.
func (r RoutePage) IsEmpty() (bool, error) {
	is, err := ExtractRoutes(r)

	return len(is) == 0, err
}

// ExtractRoutesInto interprets the results of a single page from a List() call, producing a slice of route entities.
func ExtractRoutesInto(r pagination.Page, v interface{}) error {
	return r.(RoutePage).Result.ExtractIntoSlicePtr(v, "routes")
}

// ExtractRoutes accepts a Page struct, specifically a NetworkPage struct, and extracts the elements into a slice of Route structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractRoutes(r pagination.Page) ([]Route, error) {
	var s []Route

	err := ExtractRoutesInto(r, &s)

	return s, err
}
