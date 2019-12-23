package ports

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// Extract is a function that accepts a result and extracts a port resource.
func (r commonResult) Extract() (*Port, error) {
	var s Port
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "port")
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Port.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Port.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Port.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// IP is a sub-struct that represents an individual IP.
type IP struct {
	SubnetID  string `json:"subnet_id"`
	IPAddress string `json:"ip_address,omitempty"`
}

// AddressPair contains the IP Address and the MAC address.
type AddressPair struct {
	IPAddress  string `json:"ip_address,omitempty"`
	MACAddress string `json:"mac_address,omitempty"`
}

// Port represents a Neutron port. See package documentation for a top-level
// description of what this is.
type Port struct {
	// Administrative state of port. If false (down), port does not forward
	// packets.
	AdminStateUp bool `json:"admin_state_up"`

	// Identifies the list of IP addresses the port will recognize/accept
	AllowedAddressPairs []AddressPair `json:"allowed_address_pairs"`

	// Description is description
	Description string `json:"description"`

	// Identifies the device (e.g., virtual server) using this port.
	DeviceID string `json:"device_id"`

	// Identifies the entity (e.g.: dhcp agent) using this port.
	DeviceOwner string `json:"device_owner"`

	// Specifies IP addresses for the port thus associating the port itself with
	// the subnets where the IP addresses are picked from
	FixedIPs []IP `json:"fixed_ips"`

	// UUID for the port.
	ID string `json:"id"`

	// Mac address to use on this port.
	MACAddress string `json:"mac_address"`

	// ManagedByService is set to true if only admin can modify it. Normal user has only read access.
	ManagedByService bool `json:"managed_by_service"`

	// Human-readable name for the port. Might not be unique.
	Name string `json:"name"`

	// Network that this port is associated with.
	NetworkID string `json:"network_id"`

	// SegmentationID is the segmenation ID used for this port (i.e. for vlan type it is vlan tag)
	SegmentationID int `json:"segmentation_id"`

	// SegmenationType is the segmentation type used for this port (i.e. vlan)
	SegmentationType string `json:"segmentation_type"`

	// Indicates whether network is currently operational. Possible values include
	// `ACTIVE', `DOWN', `BUILD', or `ERROR'. Plug-ins might define additional
	// values.
	Status string `json:"status"`

	// Tags optionally set via extensions/attributestags
	Tags map[string]string `json:"tags"`

	// TenantID is the project owner of the port.
	TenantID string `json:"tenant_id"`
}

// PortPage is the page returned by a pager when traversing over a collection
// of network ports.
type PortPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of ports has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r PortPage) NextPageURL() (string, error) {
	var s struct {
		Links []eclcloud.Link `json:"ports_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return eclcloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a PortPage struct is empty.
func (r PortPage) IsEmpty() (bool, error) {
	is, err := ExtractPorts(r)
	return len(is) == 0, err
}

// ExtractPorts accepts a Page struct, specifically a PortPage struct,
// and extracts the elements into a slice of Port structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractPorts(r pagination.Page) ([]Port, error) {
	var s []Port
	err := ExtractPortsInto(r, &s)
	return s, err
}

func ExtractPortsInto(r pagination.Page, v interface{}) error {
	return r.(PortPage).Result.ExtractIntoSlicePtr(v, "ports")
}
