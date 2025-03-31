package network_based_device_single

import (
	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// Extract is a function that accepts a result
// and extracts a Single Device resource.
func (r commonResult) Extract() (*SingleDeviceOrder, error) {
	var sdo SingleDeviceOrder
	err := r.ExtractInto(&sdo)
	return &sdo, err
}

// Extract interprets any commonResult as a Single Device if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Single Device.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Single Device.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Single Device.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	commonResult
}

// SingleDevice represents the result of a each element in
// response of single device api result.
type SingleDevice struct {
	ID   int      `json:"id"`
	Cell []string `json:"cell"`
}

// SingleDeviceOrder represents a Single Device's each order.
type SingleDeviceOrder struct {
	ID      string `json:"soId"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// SingleDevicePage is the page returned by a pager
// when traversing over a collection of Single Device.
type SingleDevicePage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of Single Device
// has reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r SingleDevicePage) NextPageURL() (string, error) {
	var s struct {
		Links []eclcloud.Link `json:"single_device_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return eclcloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a SingleDevicePage struct is empty.
func (r SingleDevicePage) IsEmpty() (bool, error) {
	is, err := ExtractSingleDevices(r)
	return len(is) == 0, err
}

// ExtractSingleDevices accepts a Page struct,
// specifically a SingleDevicePage struct, and extracts the elements
// into a slice of Single Device structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractSingleDevices(r pagination.Page) ([]SingleDevice, error) {
	var s []SingleDevice
	err := ExtractSingleDevicesInto(r, &s)
	return s, err
}

// ExtractSingleDevicesInto interprets the results of a single page from a List() call,
// producing a slice of Device entities.
func ExtractSingleDevicesInto(r pagination.Page, v interface{}) error {
	return r.(SingleDevicePage).Result.ExtractIntoSlicePtr(v, "rows")
}
