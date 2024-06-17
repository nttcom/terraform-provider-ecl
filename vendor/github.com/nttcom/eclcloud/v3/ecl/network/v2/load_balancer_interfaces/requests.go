package load_balancer_interfaces

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the Load Balancer Interface attributes you want to see returned. SortKey allows you to sort
// by a particular Load Balancer Interface attribute. SortDir sets the direction, and is either
// `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	Description      string `q:"description"`
	ID               string `q:"id"`
	IPAddress        string `q:"ip_address"`
	LoadBalancerID   string `q:"load_balancer_id"`
	Name             string `q:"name"`
	NetworkID        string `q:"network_id"`
	SlotNumber       int    `q:"slot_number"`
	Status           string `q:"status"`
	TenantID         string `q:"tenant_id"`
	VirtualIPAddress string `q:"virtual_ip_address"`
}

// ToLoadBalancerInterfacesListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToLoadBalancerInterfacesListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// Load Balancer Interfaces. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those Load Balancer Interfaces that are owned by the tenant
// who submits the request, unless the request is submitted by a user with
// administrative rights.
func List(c *eclcloud.ServiceClient, opts ListOpts) pagination.Pager {
	url := listURL(c)
	query, err := opts.ToLoadBalancerInterfacesListQuery()
	if err != nil {
		return pagination.Pager{Err: err}
	}
	url += query
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return LoadBalancerInterfacePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific Load Balancer Interface based on its unique ID.
func Get(c *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, id), &r.Body, nil)
	return
}

// UpdateOpts represents the attributes used when updating an existing Load Balancer Interface.
type UpdateOpts struct {

	// Description is description
	Description *string `json:"description,omitempty"`

	// IP Address
	IPAddress string `json:"ip_address,omitempty"`

	// Name of the Load Balancer Interface
	Name *string `json:"name,omitempty"`

	// UUID of the parent network.
	NetworkID *interface{} `json:"network_id,omitempty"`

	// Virtual IP Address
	VirtualIPAddress *interface{} `json:"virtual_ip_address,omitempty"`

	// Properties used for virtual IP address
	VirtualIPProperties *VirtualIPProperties `json:"virtual_ip_properties,omitempty"`
}

// ToLoadBalancerUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToLoadBalancerInterfaceUpdateMap() (map[string]interface{}, error) {
	b, err := eclcloud.BuildRequestBody(opts, "load_balancer_interface")
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Update accepts a UpdateOpts struct and updates an existing Load Balancer Interface using the
// values provided.
func Update(c *eclcloud.ServiceClient, id string, opts UpdateOpts) (r UpdateResult) {
	b, err := opts.ToLoadBalancerInterfaceUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(updateURL(c, id), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// IDFromName is a convenience function that returns a Load Balancer Interface's ID,
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

	all, err := ExtractLoadBalancerInterfaces(pages)
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
		return "", eclcloud.ErrResourceNotFound{Name: name, ResourceType: "load_balancer_interface"}
	case 1:
		return id, nil
	default:
		return "", eclcloud.ErrMultipleResourcesFound{Name: name, Count: count, ResourceType: "load_balancer_interface"}
	}
}
