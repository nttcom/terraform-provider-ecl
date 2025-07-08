package health_monitors

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// CreateResult represents the result of a Create operation.
// Call its Extract method to interpret it as a HealthMonitor.
type CreateResult struct {
	commonResult
}

// ShowResult represents the result of a Show operation.
// Call its Extract method to interpret it as a HealthMonitor.
type ShowResult struct {
	commonResult
}

// UpdateResult represents the result of a Update operation.
// Call its Extract method to interpret it as a HealthMonitor.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a Delete operation.
// Call its ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// CreateStagedResult represents the result of a CreateStaged operation.
// Call its Extract method to interpret it as a HealthMonitor.
type CreateStagedResult struct {
	commonResult
}

// ShowStagedResult represents the result of a ShowStaged operation.
// Call its Extract method to interpret it as a HealthMonitor.
type ShowStagedResult struct {
	commonResult
}

// UpdateStagedResult represents the result of a UpdateStaged operation.
// Call its Extract method to interpret it as a HealthMonitor.
type UpdateStagedResult struct {
	commonResult
}

// CancelStagedResult represents the result of a CancelStaged operation.
// Call its ExtractErr method to determine if the request succeeded or failed.
type CancelStagedResult struct {
	eclcloud.ErrResult
}

// ConfigurationInResponse represents a configuration in a health monitor.
type ConfigurationInResponse struct {

	// - Port number of the health monitor for healthchecking
	// - If `protocol` is `"icmp"`, returns `0`
	Port int `json:"port,omitempty"`

	// - Protocol of the health monitor for healthchecking
	Protocol string `json:"protocol,omitempty"`

	// - Interval of healthchecking (in seconds)
	Interval int `json:"interval,omitempty"`

	// - Retry count of healthchecking
	// - Initial monitoring is not included
	// - Retry is executed at the interval set in `interval`
	Retry int `json:"retry,omitempty"`

	// - Timeout of healthchecking (in seconds)
	Timeout int `json:"timeout,omitempty"`

	// - URL path of healthchecking
	// - If `protocol` is `"http"` or `"https"`, uses this parameter
	Path string `json:"path,omitempty"`

	// - HTTP status codes expected in healthchecking
	// - If `protocol` is `"http"` or `"https"`, uses this parameter
	// - Format: `"xxx"` or `"xxx-xxx"` ( `xxx` between [100, 599])
	HttpStatusCode string `json:"http_status_code,omitempty"`
}

// HealthMonitor represents a health monitor.
type HealthMonitor struct {

	// - ID of the health monitor
	ID string `json:"id"`

	// - Name of the health monitor
	Name string `json:"name"`

	// - Description of the health monitor
	Description string `json:"description"`

	// - Tags of the health monitor (JSON object format)
	Tags map[string]interface{} `json:"tags"`

	// - Configuration status of the health monitor
	//   - `"ACTIVE"`
	//     - There are no configurations of the health monitor that waiting to be applied
	//   - `"CREATE_STAGED"`
	//     - The health monitor has been added and waiting to be applied
	//   - `"UPDATE_STAGED"`
	//     - Changed configurations of the health monitor exists that waiting to be applied
	//   - `"DELETE_STAGED"`
	//     - The health monitor has been removed and waiting to be applied
	// - For detail, refer to the API reference appendix
	//     - https://sdpf.ntt.com/services/docs/managed-lb/service-descriptions/api_reference_appendix.html
	ConfigurationStatus string `json:"configuration_status"`

	// - Operation status of the load balancer which the health monitor belongs to
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
	//     - The operators will investigate the operation
	//     - The load balancer and related resources cannot be operated
	//   - `"ERROR"`
	//     - The latest operation of the load balancer has been failed
	//     - The operation was roll backed normally
	//     - The load balancer and related resources can be operated
	// - For detail, refer to the API reference appendix
	//     - https://sdpf.ntt.com/services/docs/managed-lb/service-descriptions/api_reference_appendix.html
	OperationStatus string `json:"operation_status"`

	// - ID of the load balancer which the health monitor belongs to
	LoadBalancerID string `json:"load_balancer_id"`

	// - ID of the owner tenant of the health monitor
	TenantID string `json:"tenant_id"`

	// - Port number of the health monitor for healthchecking
	// - If `protocol` is `"icmp"`, returns `0`
	Port int `json:"port,omitempty"`

	// - Protocol of the health monitor for healthchecking
	Protocol string `json:"protocol,omitempty"`

	// - Interval of healthchecking (in seconds)
	Interval int `json:"interval,omitempty"`

	// - Retry count of healthchecking
	// - Initial monitoring is not included
	// - Retry is executed at the interval set in `interval`
	Retry int `json:"retry,omitempty"`

	// - Timeout of healthchecking (in seconds)
	Timeout int `json:"timeout,omitempty"`

	// - URL path of healthchecking
	// - If `protocol` is `"http"` or `"https"`, uses this parameter
	Path string `json:"path,omitempty"`

	// - HTTP status codes expected in healthchecking
	// - If `protocol` is `"http"` or `"https"`, uses this parameter
	// - Format: `"xxx"` or `"xxx-xxx"` ( `xxx` between [100, 599])
	HttpStatusCode string `json:"http_status_code,omitempty"`

	// - Running configurations of the health monitor
	// - If `changes` is `true`, return object
	// - If current configuration does not exist, return `null`
	Current ConfigurationInResponse `json:"current,omitempty"`

	// - Added or changed configurations of the health monitor that waiting to be applied
	// - If `changes` is `true`, return object
	// - If staged configuration does not exist, return `null`
	Staged ConfigurationInResponse `json:"staged,omitempty"`
}

// ExtractInto interprets any commonResult as a health monitor, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "health_monitor")
}

// Extract is a function that accepts a result and extracts a HealthMonitor resource.
func (r commonResult) Extract() (*HealthMonitor, error) {
	var healthMonitor HealthMonitor

	err := r.ExtractInto(&healthMonitor)

	return &healthMonitor, err
}

// HealthMonitorPage is the page returned by a pager when traversing over a collection of health monitor.
type HealthMonitorPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a HealthMonitorPage struct is empty.
func (r HealthMonitorPage) IsEmpty() (bool, error) {
	is, err := ExtractHealthMonitors(r)

	return len(is) == 0, err
}

// ExtractHealthMonitorsInto interprets the results of a single page from a List() call, producing a slice of health monitor entities.
func ExtractHealthMonitorsInto(r pagination.Page, v interface{}) error {
	return r.(HealthMonitorPage).Result.ExtractIntoSlicePtr(v, "health_monitors")
}

// ExtractHealthMonitors accepts a Page struct, specifically a NetworkPage struct, and extracts the elements into a slice of HealthMonitor structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractHealthMonitors(r pagination.Page) ([]HealthMonitor, error) {
	var s []HealthMonitor

	err := ExtractHealthMonitorsInto(r, &s)

	return s, err
}
