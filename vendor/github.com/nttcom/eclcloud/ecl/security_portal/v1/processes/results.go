package processes

import (
	"github.com/nttcom/eclcloud"
)

type commonResult struct {
	eclcloud.Result
}

// Extract is a function that accepts a result
// and extracts a Process.
func (r commonResult) Extract() (*ProcessInstance, error) {
	var pr ProcessInstance
	err := r.ExtractInto(&pr)
	return &pr, err
}

// Extract interprets any commonResult as a Process, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "processInstance")
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Process.
type GetResult struct {
	commonResult
}

type ProcessStatus struct {
	Status string `json:"status"`
}

type ProcessInstance struct {
	Status ProcessStatus `json:"status"`
}
