package load_balancer_interfaces

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Load Balancer Interface.
type UpdateResult struct {
	commonResult
}

// Extract is a function that accepts a result and extracts a Load Balancer Interface resource.
func (r commonResult) Extract() (*LoadBalancerInterface, error) {
	var s struct {
		LoadBalancerInterface *LoadBalancerInterface `json:"load_balancer_interface"`
	}
	err := r.ExtractInto(&s)
	return s.LoadBalancerInterface, err
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Load Balancer Interface.
type GetResult struct {
	commonResult
}

// Properties used for virtual IP address
type VirtualIPProperties struct {
	Protocol string `json:"protocol"`
	Vrid     int    `json:"vrid"`
}

// LoadBalancerInterface represents a Load Balancer Interface. See package documentation for a top-level
// description of what this is.
type LoadBalancerInterface struct {

	// Description is description
	Description string `json:"description"`

	// UUID representing the Load Balancer Interface.
	ID string `json:"id"`

	// IP Address
	IPAddress *string `json:"ip_address"`

	// The ID of load_balancer this load_balancer_interface belongs to.
	LoadBalancerID string `json:"load_balancer_id"`

	// Name of the Load Balancer Interface
	Name string `json:"name"`

	// UUID of the parent network.
	NetworkID *string `json:"network_id"`

	// Slot Number
	SlotNumber int `json:"slot_number"`

	// Load Balancer Interface status
	Status string `json:"status"`

	// Tenant ID of the owner (UUID)
	TenantID string `json:"tenant_id"`

	// Load Balancer Interface type
	Type string `json:"type"`

	// Virtual IP Address
	VirtualIPAddress *string `json:"virtual_ip_address"`

	// Properties used for virtual IP address
	VirtualIPProperties *VirtualIPProperties `json:"virtual_ip_properties"`
}

// LoadBalancerPage is the page returned by a pager when traversing over a collection
// of load balancers.
type LoadBalancerInterfacePage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a LoadBalancerInterfacePage struct is empty.
func (r LoadBalancerInterfacePage) IsEmpty() (bool, error) {
	is, err := ExtractLoadBalancerInterfaces(r)
	return len(is) == 0, err
}

// ExtractLoadBalancerInterfaces accepts a Page struct, specifically a LoadBalancerPage struct,
// and extracts the elements into a slice of Load Balancer Interface structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractLoadBalancerInterfaces(r pagination.Page) ([]LoadBalancerInterface, error) {
	var s struct {
		LoadBalancerInterfaces []LoadBalancerInterface `json:"load_balancer_interfaces"`
	}
	err := (r.(LoadBalancerInterfacePage)).ExtractInto(&s)
	return s.LoadBalancerInterfaces, err
}
