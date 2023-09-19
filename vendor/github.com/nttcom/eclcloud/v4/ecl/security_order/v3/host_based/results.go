package host_based

import (
	"github.com/nttcom/eclcloud/v4"
)

type commonResult struct {
	eclcloud.Result
}

// Extract is a function that accepts a result
// and extracts a Host Based Security resource.
func (r commonResult) Extract() (*HostBasedOrder, error) {
	var hbo HostBasedOrder
	err := r.ExtractInto(&hbo)
	return &hbo, err
}

// Extract interprets any commonResult as a Host Based Security if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Host Based Security.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Host Based Security.
type GetResult struct {
	commonResult
}

// HostBasedSecurity represents a Host Based Security's each order.
type HostBasedSecurity struct {
	Code                string      `json:"code"`
	Message             string      `json:"message"`
	Region              string      `json:"region"`
	TenantName          string      `json:"tenant_name"`
	TenantDescription   string      `json:"tenant_description"`
	ContractID          string      `json:"contract_id"`
	ServiceOrderService string      `json:"service_order_service"`
	MaxAgentValue       interface{} `json:"max_agent_value"`
	TimeZone            string      `json:"time_zone"`
	CustomerName        string      `json:"customer_name"`
	MailAddress         string      `json:"mailaddress"`
	DSMLang             string      `json:"dsm_lang"`
	TenantFlg           bool        `json:"tenant_flg"`
	Status              int         `json:"status"`
}

// Extract is a function that accepts a result
// and extracts a Host Based Security resource.
func (r GetResult) Extract() (*HostBasedSecurity, error) {
	var h HostBasedSecurity
	err := r.ExtractInto(&h)
	return &h, err
}

// ExtractInto interprets any commonResult as a Host Based Security if possible.
func (r GetResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Host Based Security.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	commonResult
}

// HostBasedOrder represents a Host Based Security's each order.
type HostBasedOrder struct {
	ID      string `json:"soId"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}
