package appliance_plans

import (
	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// Extract is a function that accepts a result and extracts a Virtual Network Appliance Plan resource.
func (r commonResult) Extract() (*VirtualNetworkAppliancePlan, error) {
	var s struct {
		VirtualNetworkAppliancePlan *VirtualNetworkAppliancePlan `json:"virtual_network_appliance_plan"`
	}
	err := r.ExtractInto(&s)
	return s.VirtualNetworkAppliancePlan, err
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Virtual Network Appliance Plan.
type GetResult struct {
	commonResult
}

// License of Virtual Network Appliance
type License struct {
	LicenseType string `json:"license_type"`
}

// Availability Zone of Virtual Network Appliance
type AvailabilityZone struct {
	AvailabilityZone string `json:"availability_zone"`
	Available        bool   `json:"available"`
	Rank             int    `json:"rank"`
}

// VirtualNetworkAppliancePlan represents a Virtual Network Appliance Plan.
// See package documentation for a top-level description of what this is.
type VirtualNetworkAppliancePlan struct {

	// UUID representing the Virtual Network Appliance Plan.
	ID string `json:"id"`

	// Name of the Virtual Network Appliance Plan.
	Name string `json:"name"`

	// Description is description
	Description string `json:"description"`

	// Type of appliance
	ApplianceType string `json:"appliance_type"`

	// Version name
	Version string `json:"version"`

	// Nova flavor
	Flavor string `json:"flavor"`

	// Number of Interfaces
	NumberOfInterfaces int `json:"number_of_interfaces"`

	// Is user allowed to create new firewalls with this plan.
	Enabled bool `json:"enabled"`

	// Max Number of allowed_address_pairs
	MaxNumberOfAap int `json:"max_number_of_aap"`

	// Licenses
	Licenses []License `json:"licenses"`

	// AvailabilityZones
	AvailabilityZones []AvailabilityZone `json:"availability_zones"`
}

// VirtualNetworkAppliancePlanPage is the page returned by a pager when traversing over a collection
// of virtual network appliance plans.
type VirtualNetworkAppliancePlanPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a VirtualNetworkAppliancePlanPage struct is empty.
func (r VirtualNetworkAppliancePlanPage) IsEmpty() (bool, error) {
	is, err := ExtractVirtualNetworkAppliancePlans(r)
	return len(is) == 0, err
}

// ExtractVirtualNetworkAppliancePlans accepts a Page struct, specifically a VirtualNetworkAppliancePlanPage struct,
// and extracts the elements into a slice of Virtual Network Appliance Plan structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractVirtualNetworkAppliancePlans(r pagination.Page) ([]VirtualNetworkAppliancePlan, error) {
	var s struct {
		VirtualNetworkAppliancePlans []VirtualNetworkAppliancePlan `json:"virtual_network_appliance_plans"`
	}
	err := (r.(VirtualNetworkAppliancePlanPage)).ExtractInto(&s)
	return s.VirtualNetworkAppliancePlans, err
}
