package devices

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// Device represents the result of a each element in
// response of device api result.
type Device struct {
	MSADeviceID        string `json:"msa_device_id"`
	OSServerID         string `json:"os_server_id"`
	OSServerName       string `json:"os_server_name"`
	OSAvailabilityZone string `json:"os_availability_zone"`
	OSAdminUserName    string `json:"os_admin_username"`
	MSADeviceType      string `json:"msa_device_type"`
	OSServerStatus     string `json:"os_server_status"`
}

// DevicePage is the page returned by a pager
// when traversing over a collection of Device.
type DevicePage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of Device
// has reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r DevicePage) NextPageURL() (string, error) {
	var s struct {
		Links []eclcloud.Link `json:"devices"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return eclcloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a Device struct is empty.
func (r DevicePage) IsEmpty() (bool, error) {
	is, err := ExtractDevices(r)
	return len(is) == 0, err
}

// ExtractDevices accepts a Page struct,
// specifically a DevicePage struct, and extracts the elements
// into a slice of Device structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractDevices(r pagination.Page) ([]Device, error) {
	var d []Device
	err := ExtractDevicesInto(r, &d)
	return d, err
}

// ExtractDevicesInto interprets the results of a single page from a List() call,
// producing a slice of Device entities.
func ExtractDevicesInto(r pagination.Page, v interface{}) error {
	return r.(DevicePage).Result.ExtractIntoSlicePtr(v, "devices")
}
