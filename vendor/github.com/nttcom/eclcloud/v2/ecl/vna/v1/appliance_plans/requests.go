package appliance_plans

import (
	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/pagination"
)

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the Virtual Network Appliance Plan attributes you want to see returned.
type ListOpts struct {
	ID                        string `q:"id"`
	Name                      string `q:"name"`
	Description               string `q:"description"`
	ApplianceType             string `q:"appliance_type"`
	Version                   string `q:"version"`
	Flavor                    string `q:"flavor"`
	NumberOfInterfaces        int    `q:"number_of_interfaces"`
	Enabled                   bool   `q:"enabled"`
	MaxNumberOfAap            int    `q:"max_number_of_aap"`
	Details                   bool   `q:"details"`
	AvailabilityZone          string `q:"availability_zone"`
	AvailabilityZoneAvailable bool   `q:"availability_zone.available"`
}

// ToVirtualNetworkAppliancePlanListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToVirtualNetworkAppliancePlanListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// Virtual Network Appliance Plans. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(c *eclcloud.ServiceClient, opts ListOpts) pagination.Pager {
	url := listURL(c)
	query, err := opts.ToVirtualNetworkAppliancePlanListQuery()
	if err != nil {
		return pagination.Pager{Err: err}
	}
	url += query
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return VirtualNetworkAppliancePlanPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// GetOptsBuilder allows extensions to add additional parameters to
// the Virtual Network Appliance Plan API request
type GetOptsBuilder interface {
	ToProcessQuery() (string, error)
}

// GetOpts represents result of Virtual Network Appliance Plan API response.
type GetOpts struct {
	VirtualNetworkAppliancePlanId string `q:"virtual_network_appliance_plan_id"`
	Details                       bool   `q:"details"`
}

// ToProcessQuery formats a GetOpts into a query string.
func (opts GetOpts) ToProcessQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// Get retrieves a specific Virtual Network Appliance Plan based on its unique ID.
func Get(c *eclcloud.ServiceClient, id string, opts GetOptsBuilder) (r GetResult) {
	url := getURL(c, id)
	if opts != nil {
		query, _ := opts.ToProcessQuery()
		url += query
	}
	_, r.Err = c.Get(url, &r.Body, nil)
	return
}

// IDFromName is a convenience function that returns a Virtual Network Appliance Plan's ID,
// given its name.
func IDFromName(client *eclcloud.ServiceClient, name string) (string, error) {
	count := 0
	id := ""

	listOpts := ListOpts{
		Name: name,
	}

	pages, err := List(client, listOpts).AllPages()
	if err != nil {
		return "", err
	}

	all, err := ExtractVirtualNetworkAppliancePlans(pages)
	if err != nil {
		return "", err
	}

	for _, s := range all {
		if s.Name == name {
			count++
			id = s.ID
		}
	}

	switch count {
	case 0:
		return "", eclcloud.ErrResourceNotFound{Name: name, ResourceType: "virtual_network_appliance_plan"}
	case 1:
		return id, nil
	default:
		return "", eclcloud.ErrMultipleResourcesFound{Name: name, Count: count, ResourceType: "virtual_network_appliance_plan"}
	}
}
