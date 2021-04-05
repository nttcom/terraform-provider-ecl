package approval_requests

import (
	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/pagination"
)

type Action struct {
	Service string `json:"service"`
	Region  string `json:"region"`
	APIPath string `json:"api_path"`
	Method  string `json:"method"`
	// Basically JSON is passed to Action.Body,
	// but depending on the value of the service, it may be a String, so it is set to interface{}.
	// If service is "provider-connectivity", body's type is JSON.
	// If service is "network", body's type is String.
	Body interface{} `json:"body"`
}

type Description struct {
	Lang string `json:"lang"`
	Text string `json:"text"`
}

// ApprovalRequest represents an ECL SSS Approval Request.
type ApprovalRequest struct {
	RequestID         string        `json:"request_id"`
	ExternalRequestID string        `json:"external_request_id"`
	ApproverType      string        `json:"approver_type"`
	ApproverID        string        `json:"approver_id"`
	RequestUserID     string        `json:"request_user_id"`
	Service           string        `json:"service"`
	Actions           []Action      `json:"actions"`
	Descriptions      []Description `json:"descriptions"`
	RequestUser       interface{}   `json:"request_user"`
	Approver          bool          `json:"approver"`
	ApprovalDeadLine  interface{}   `json:"approval_deadline"`
	ApprovalExpire    interface{}   `json:"approval_expire"`
	RegisteredTime    interface{}   `json:"registered_time"`
	UpdatedTime       interface{}   `json:"updated_time"`
	Status            string        `json:"status"`
}

type commonResult struct {
	eclcloud.Result
}

func (r commonResult) Extract() (*ApprovalRequest, error) {
	var ar ApprovalRequest
	err := r.ExtractInto(&ar)
	return &ar, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}

// GetResult is the response from a Get operation. Call its Extract method
// to interpret it as an approval request.
type GetResult struct {
	commonResult
}

// UpdateResult is the result of an Update request. Call its Extract method to
// interpret it as an approval request.
type UpdateResult struct {
	commonResult
}

// ApprovalRequestPage is a single page of approval request results.
type ApprovalRequestPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of approval requests contains any results.
func (r ApprovalRequestPage) IsEmpty() (bool, error) {
	resources, err := ExtractApprovalRequests(r)
	return len(resources) == 0, err
}

// ExtractApprovalRequests returns a slice of approval requests
// contained in a single page of results.
func ExtractApprovalRequests(r pagination.Page) ([]ApprovalRequest, error) {
	var s struct {
		ApprovalRequests []ApprovalRequest `json:"approval_requests"`
	}
	err := (r.(ApprovalRequestPage)).ExtractInto(&s)
	return s.ApprovalRequests, err
}

// ExtractApprovalRequestsInto interprets the results of a single page from a List() call,
// producing a slice of Approval Request entities.
func ExtractApprovalRequestsInto(r pagination.Page, v interface{}) error {
	return r.(ApprovalRequestPage).Result.ExtractIntoSlicePtr(v, "")
}
