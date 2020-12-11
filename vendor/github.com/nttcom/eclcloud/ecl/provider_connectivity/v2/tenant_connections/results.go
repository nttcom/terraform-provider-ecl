package tenant_connections

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

// TenantConnection represents Tenant Connection.
// TagsOther is interface{} because the data type returned by Create API depends on the value of device_type.
// When the device_type of Create Request is ECL::Compute::Server, the data type of tags_other is map[].
// When the device_type of Create Request is ECL::Baremetal::Server or ECL::VirtualNetworkAppliance::VSRX, the data type of tags_other is string.
type TenantConnection struct {
	ID                        string            `json:"id"`
	TenantConnectionRequestID string            `json:"tenant_connection_request_id"`
	Name                      string            `json:"name"`
	Description               string            `json:"description"`
	Tags                      map[string]string `json:"tags"`
	TenantID                  string            `json:"tenant_id"`
	NameOther                 string            `json:"name_other"`
	DescriptionOther          string            `json:"description_other"`
	TagsOther                 interface{}       `json:"tags_other"`
	TenantIDOther             string            `json:"tenant_id_other"`
	NetworkID                 string            `json:"network_id"`
	DeviceType                string            `json:"device_type"`
	DeviceID                  string            `json:"device_id"`
	DeviceInterfaceID         string            `json:"device_interface_id"`
	PortID                    string            `json:"port_id"`
	Status                    string            `json:"status"`
}

type commonResult struct {
	eclcloud.Result
}

// GetResult is the response from a Get operation. Call its Extract method
// to interpret it as a Tenant Connection.
type GetResult struct {
	commonResult
}

// CreateResult is the response from a Create operation. Call its Extract method
// to interpret it as a Tenant Connection.
type CreateResult struct {
	commonResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr to
// determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// UpdateResult is the result of an Update request. Call its Extract method to
// interpret it as a Tenant Connection.
type UpdateResult struct {
	commonResult
}

// TenantConnectionPage is a single page of Tenant Connection results.
type TenantConnectionPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of Tenant Connection contains any results.
func (r TenantConnectionPage) IsEmpty() (bool, error) {
	resources, err := ExtractTenantConnections(r)
	return len(resources) == 0, err
}

// ExtractTenantConnections returns a slice of Tenant Connections contained in a
// single page of results.
func ExtractTenantConnections(r pagination.Page) ([]TenantConnection, error) {
	var s struct {
		TenantConnection []TenantConnection `json:"tenant_connections"`
	}
	err := (r.(TenantConnectionPage)).ExtractInto(&s)
	return s.TenantConnection, err
}

// Extract interprets any commonResult as a Tenant Connection.
func (r commonResult) Extract() (*TenantConnection, error) {
	var s struct {
		TenantConnection *TenantConnection `json:"tenant_connection"`
	}
	err := r.ExtractInto(&s)
	return s.TenantConnection, err
}
