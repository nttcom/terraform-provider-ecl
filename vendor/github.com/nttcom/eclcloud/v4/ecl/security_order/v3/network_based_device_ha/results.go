package network_based_device_ha

import (
	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// Extract is a function that accepts a result
// and extracts a HA Device resource.
func (r commonResult) Extract() (*HADeviceOrder, error) {
	var sdo HADeviceOrder
	err := r.ExtractInto(&sdo)
	return &sdo, err
}

// Extract interprets any commonResult as a HA Device if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a HA Device.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a HA Device.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a HA Device.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	commonResult
}

// HADevice represents the result of a each element in
// response of HA Device api result.
type HADevice struct {
	ID   int      `json:"id"`
	Cell []string `json:"cell"`
}

// HADeviceOrder represents a HA Device's each order.
type HADeviceOrder struct {
	ID      string `json:"soId"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// HADevicePage is the page returned by a pager
// when traversing over a collection of HA Device.
type HADevicePage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of HA Device
// has reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r HADevicePage) NextPageURL() (string, error) {
	var s struct {
		Links []eclcloud.Link `json:"ha_device_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return eclcloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a HADevicePage struct is empty.
func (r HADevicePage) IsEmpty() (bool, error) {
	is, err := ExtractHADevices(r)
	return len(is) == 0, err
}

// ExtractHADevices accepts a Page struct,
// specifically a HADevicePage struct, and extracts the elements
// into a slice of HA Device structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractHADevices(r pagination.Page) ([]HADevice, error) {
	var s []HADevice
	err := ExtractHADevicesInto(r, &s)
	return s, err
}

// ExtractHADevicesInto interprets the results of a single page from a List() call,
// producing a slice of Device entities.
func ExtractHADevicesInto(r pagination.Page, v interface{}) error {
	return r.(HADevicePage).Result.ExtractIntoSlicePtr(v, "rows")
}
