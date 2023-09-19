package service_order_status

import (
	"github.com/nttcom/eclcloud/v4"
)

type commonResult struct {
	eclcloud.Result
}

// Extract is a function that accepts a result
// and extracts an Order Progress resource.
func (r commonResult) Extract() (*OrderProgress, error) {
	var sd OrderProgress
	err := r.ExtractInto(&sd)
	return &sd, err
}

// Extract interprets any commonResult as an Order Progress, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as an Order.
type GetResult struct {
	commonResult
}

// OrderProgress represents an Order Progress response.
type OrderProgress struct {
	Status       int    `json:"status"`
	Code         string `json:"code"`
	Message      string `json:"message"`
	ProgressRate int    `json:"progressRate"`
}
