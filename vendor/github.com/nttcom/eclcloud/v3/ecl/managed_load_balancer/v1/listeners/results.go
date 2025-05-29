package listeners

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// CreateResult represents the result of a Create operation.
// Call its Extract method to interpret it as a Listener.
type CreateResult struct {
	commonResult
}

// ShowResult represents the result of a Show operation.
// Call its Extract method to interpret it as a Listener.
type ShowResult struct {
	commonResult
}

// UpdateResult represents the result of a Update operation.
// Call its Extract method to interpret it as a Listener.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a Delete operation.
// Call its ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// CreateStagedResult represents the result of a CreateStaged operation.
// Call its Extract method to interpret it as a Listener.
type CreateStagedResult struct {
	commonResult
}

// ShowStagedResult represents the result of a ShowStaged operation.
// Call its Extract method to interpret it as a Listener.
type ShowStagedResult struct {
	commonResult
}

// UpdateStagedResult represents the result of a UpdateStaged operation.
// Call its Extract method to interpret it as a Listener.
type UpdateStagedResult struct {
	commonResult
}

// CancelStagedResult represents the result of a CancelStaged operation.
// Call its ExtractErr method to determine if the request succeeded or failed.
type CancelStagedResult struct {
	eclcloud.ErrResult
}

// ConfigurationInResponse represents a configuration in a listener.
type ConfigurationInResponse struct {

	// - IP address of the listener for listening
	IPAddress string `json:"ip_address,omitempty"`

	// - Port number of the listener for listening
	Port int `json:"port,omitempty"`

	// - Protocol of the listener for listening
	Protocol string `json:"protocol,omitempty"`
}

// Listener represents a listener.
type Listener struct {

	// - ID of the listener
	ID string `json:"id"`

	// - Name of the listener
	Name string `json:"name"`

	// - Description of the listener
	Description string `json:"description"`

	// - Tags of the listener (JSON object format)
	Tags map[string]interface{} `json:"tags"`

	// - Configuration status of the listener
	//   - `"ACTIVE"`
	//     - There are no configurations of the listener that waiting to be applied
	//   - `"CREATE_STAGED"`
	//     - The listener has been added and waiting to be applied
	//   - `"UPDATE_STAGED"`
	//     - Changed configurations of the listener exists that waiting to be applied
	//   - `"DELETE_STAGED"`
	//     - The listener has been removed and waiting to be applied
	// - For detail, refer to the API reference appendix
	//     - https://sdpf.ntt.com/services/docs/managed-lb/service-descriptions/api_reference_appendix.html
	ConfigurationStatus string `json:"configuration_status"`

	// - Operation status of the load balancer which the listener belongs to
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
	// - For detail, refer to the API reference appendix
	//     - https://sdpf.ntt.com/services/docs/managed-lb/service-descriptions/api_reference_appendix.html
	OperationStatus string `json:"operation_status"`

	// - ID of the load balancer which the listener belongs to
	LoadBalancerID string `json:"load_balancer_id"`

	// - ID of the owner tenant of the listener
	TenantID string `json:"tenant_id"`

	// - IP address of the listener for listening
	IPAddress string `json:"ip_address,omitempty"`

	// - Port number of the listener for listening
	Port int `json:"port,omitempty"`

	// - Protocol of the listener for listening
	Protocol string `json:"protocol,omitempty"`

	// - Running configurations of the listener
	// - If `changes` is `true`, return object
	// - If current configuration does not exist, return `null`
	Current ConfigurationInResponse `json:"current,omitempty"`

	// - Added or changed configurations of the listener that waiting to be applied
	// - If `changes` is `true`, return object
	// - If staged configuration does not exist, return `null`
	Staged ConfigurationInResponse `json:"staged,omitempty"`
}

// ExtractInto interprets any commonResult as a listener, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "listener")
}

// Extract is a function that accepts a result and extracts a Listener resource.
func (r commonResult) Extract() (*Listener, error) {
	var listener Listener

	err := r.ExtractInto(&listener)

	return &listener, err
}

// ListenerPage is the page returned by a pager when traversing over a collection of listener.
type ListenerPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a ListenerPage struct is empty.
func (r ListenerPage) IsEmpty() (bool, error) {
	is, err := ExtractListeners(r)

	return len(is) == 0, err
}

// ExtractListenersInto interprets the results of a single page from a List() call, producing a slice of listener entities.
func ExtractListenersInto(r pagination.Page, v interface{}) error {
	return r.(ListenerPage).Result.ExtractIntoSlicePtr(v, "listeners")
}

// ExtractListeners accepts a Page struct, specifically a NetworkPage struct, and extracts the elements into a slice of Listener structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractListeners(r pagination.Page) ([]Listener, error) {
	var s []Listener

	err := ExtractListenersInto(r, &s)

	return s, err
}
