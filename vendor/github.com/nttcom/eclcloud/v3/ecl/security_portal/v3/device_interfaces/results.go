package device_interfaces

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// DeviceInterface represents the result of a each element in
// response of device interface api result.
type DeviceInterface struct {
	OSIPAddress  string `json:"os_ip_address"`
	MSAPortID    string `json:"msa_port_id"`
	OSPortName   string `json:"os_port_name"`
	OSPortID     string `json:"os_port_id"`
	OSNetworkID  string `json:"os_network_id"`
	OSPortStatus string `json:"os_port_status"`
	OSMACAddress string `json:"os_mac_address"`
	OSSubnetID   string `json:"os_subnet_id"`
}

// DeviceInterfacePage is the page returned by a pager
// when traversing over a collection of Device Interface.
type DeviceInterfacePage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of Single Device Interface
// has reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r DeviceInterfacePage) NextPageURL() (string, error) {
	var s struct {
		Links []eclcloud.Link `json:"device_interfaces"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return eclcloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a DeviceInterfacePage struct is empty.
func (r DeviceInterfacePage) IsEmpty() (bool, error) {
	is, err := ExtractDeviceInterfaces(r)
	return len(is) == 0, err
}

// ExtractDeviceInterfaces accepts a Page struct,
// specifically a DeviceInterfacePage struct, and extracts the elements
// into a slice of Device Interface structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractDeviceInterfaces(r pagination.Page) ([]DeviceInterface, error) {
	var d []DeviceInterface
	err := ExtractDeviceInterfacesInto(r, &d)
	return d, err
}

// ExtractDeviceInterfacesInto interprets the results of a single page from a List() call,
// producing a slice of Device Interface entities.
func ExtractDeviceInterfacesInto(r pagination.Page, v interface{}) error {
	return r.(DeviceInterfacePage).Result.ExtractIntoSlicePtr(v, "device_interfaces")
}
