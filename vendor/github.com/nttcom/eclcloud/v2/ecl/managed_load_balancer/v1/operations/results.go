package operations

import (
	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// ShowResult represents the result of a Show operation.
// Call its Extract method to interpret it as a Operation.
type ShowResult struct {
	commonResult
}

// Operation represents a operation.
type Operation struct {

	// - ID of the operation
	ID string `json:"id"`

	// - ID of the resource
	ResourceID string `json:"resource_id"`

	// - Type of the resource
	ResourceType string `json:"resource_type"`

	// - The unique hyphenated UUID to identify the request
	//   - The UUID which has been set by X-MVNA-Request-Id in request headers
	RequestID string `json:"request_id"`

	// - Types of the request
	RequestTypes []string `json:"request_types"`

	// - Body of the request
	RequestBody map[string]interface{} `json:"request_body,omitempty"`

	// - Operation status of the resource
	Status string `json:"status"`

	// - The time when operation has been started by API execution
	// - Format: `"%Y-%m-%d %H:%M:%S"` (UTC)
	ReceptionDatetime string `json:"reception_datetime"`

	// - The time when operation has been finished
	// - Format: `"%Y-%m-%d %H:%M:%S"` (UTC)
	CommitDatetime string `json:"commit_datetime"`

	// - The warning message of operation that has been stopped or failed
	Warning string `json:"warning"`

	// - The error message of operation that has been stopped or failed
	Error string `json:"error"`

	// - ID of the owner tenant of the resource
	TenantID string `json:"tenant_id"`
}

// ExtractInto interprets any commonResult as a operation, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "operation")
}

// Extract is a function that accepts a result and extracts a Operation resource.
func (r commonResult) Extract() (*Operation, error) {
	var operation Operation

	err := r.ExtractInto(&operation)

	return &operation, err
}

// OperationPage is the page returned by a pager when traversing over a collection of operation.
type OperationPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a OperationPage struct is empty.
func (r OperationPage) IsEmpty() (bool, error) {
	is, err := ExtractOperations(r)

	return len(is) == 0, err
}

// ExtractOperationsInto interprets the results of a single page from a List() call, producing a slice of operation entities.
func ExtractOperationsInto(r pagination.Page, v interface{}) error {
	return r.(OperationPage).Result.ExtractIntoSlicePtr(v, "operations")
}

// ExtractOperations accepts a Page struct, specifically a NetworkPage struct, and extracts the elements into a slice of Operation structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractOperations(r pagination.Page) ([]Operation, error) {
	var s []Operation

	err := ExtractOperationsInto(r, &s)

	return s, err
}
