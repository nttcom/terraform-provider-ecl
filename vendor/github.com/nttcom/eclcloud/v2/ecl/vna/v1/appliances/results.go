package appliances

import (
	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// Extract is a function that accepts a result
// and extracts a virtual network appliance resource.
func (r commonResult) Extract() (*Appliance, error) {
	var vna Appliance
	err := r.ExtractInto(&vna)
	return &vna, err
}

// Extract interprets any commonResult as a Virtual Network Appliance, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "virtual_network_appliance")
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Virtual Network Appliance.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Virtual Network Appliance.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Virtual Network Appliance.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// FixedIPInResponse represents each element of fixed ips
// of virtual network appliance.
type FixedIPInResponse struct {
	IPAddress string `json:"ip_address"`
	SubnetID  string `json:"subnet_id"`
}

// AllowedAddressPairInResponse represents each element of
// allowed address pair of virtual network appliance.
type AllowedAddressPairInResponse struct {
	IPAddress  string      `json:"ip_address"`
	MACAddress string      `json:"mac_address"`
	Type       string      `json:"type"`
	VRID       interface{} `json:"vrid"`
}

// InterfaceInResponse works as parent element of
// each interface of virtual network appliance.
type InterfaceInResponse struct {
	Name                string                         `json:"name"`
	Description         string                         `json:"description"`
	NetworkID           string                         `json:"network_id"`
	Updatable           bool                           `json:"updatable"`
	Tags                map[string]string              `json:"tags"`
	FixedIPs            []FixedIPInResponse            `json:"fixed_ips"`
	AllowedAddressPairs []AllowedAddressPairInResponse `json:"allowed_address_pairs"`
}

// InterfacesInResponse works as list of interfaces
// of virtual network appliance.
type InterfacesInResponse struct {
	Interface1 InterfaceInResponse `json:"interface_1"`
	Interface2 InterfaceInResponse `json:"interface_2"`
	Interface3 InterfaceInResponse `json:"interface_3"`
	Interface4 InterfaceInResponse `json:"interface_4"`
	Interface5 InterfaceInResponse `json:"interface_5"`
	Interface6 InterfaceInResponse `json:"interface_6"`
	Interface7 InterfaceInResponse `json:"interface_7"`
	Interface8 InterfaceInResponse `json:"interface_8"`
}

// Appliance represents, well, a virtual network appliance.
type Appliance struct {
	Name               string               `json:"name"`
	ID                 string               `json:"id"`
	ApplianceType      string               `json:"appliance_type"`
	Description        string               `json:"description"`
	DefaultGateway     string               `json:"default_gateway"`
	AvailabilityZone   string               `json:"availability_zone"`
	OSMonitoringStatus string               `json:"os_monitoring_status"`
	OSLoginStatus      string               `json:"os_login_status"`
	VMStatus           string               `json:"vm_status"`
	OperationStatus    string               `json:"operation_status"`
	AppliancePlanID    string               `json:"virtual_network_appliance_plan_id"`
	TenantID           string               `json:"tenant_id"`
	Username           string               `json:"username"`
	Password           string               `json:"password"`
	Tags               map[string]string    `json:"tags"`
	Interfaces         InterfacesInResponse `json:"interfaces"`
}

// AppliancePage is the page returned by a pager
// when traversing over a collection of virtual network appliance.
type AppliancePage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of virtual network appliance
//  has reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r AppliancePage) NextPageURL() (string, error) {
	var s struct {
		Links []eclcloud.Link `json:"appliances_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return eclcloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a AppliancePage struct is empty.
func (r AppliancePage) IsEmpty() (bool, error) {
	is, err := ExtractAppliances(r)
	return len(is) == 0, err
}

// ExtractAppliances accepts a Page struct,
// specifically a NetworkPage struct, and extracts the elements
// into a slice of Virtual Network Appliance structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractAppliances(r pagination.Page) ([]Appliance, error) {
	var s []Appliance
	err := ExtractAppliancesInto(r, &s)
	return s, err
}

// ExtractAppliancesInto interprets the results of a single page from a List() call,
// producing a slice of Server entities.
func ExtractAppliancesInto(r pagination.Page, v interface{}) error {
	return r.(AppliancePage).Result.ExtractIntoSlicePtr(v, "virtual_network_appliances")
}
