package load_balancers

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the Load Balancer attributes you want to see returned. SortKey allows you to sort
// by a particular Load Balancer attribute. SortDir sets the direction, and is either
// `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	AdminUsername      string `q:"admin_username"`
	AvailabilityZone   string `q:"availability_zone"`
	DefaultGateway     string `q:"default_gateway"`
	Description        string `q:"description"`
	ID                 string `q:"id"`
	LoadBalancerPlanID string `q:"load_balancer_plan_id"`
	Name               string `q:"name"`
	Status             string `q:"status"`
	TenantID           string `q:"tenant_id"`
	UserUsername       string `q:"user_username"`
}

// ToLoadBalancersListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToLoadBalancersListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// Load Balancers. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those Load Balancers that are owned by the tenant
// who submits the request, unless the request is submitted by a user with
// administrative rights.
func List(c *eclcloud.ServiceClient, opts ListOpts) pagination.Pager {
	url := listURL(c)
	query, err := opts.ToLoadBalancersListQuery()
	if err != nil {
		return pagination.Pager{Err: err}
	}
	url += query
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return LoadBalancerPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific Load Balancer based on its unique ID.
func Get(c *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, id), &r.Body, nil)
	return
}

// CreateOpts represents the attributes used when creating a new Load Balancer.
type CreateOpts struct {

	// AvailabilityZone is one of the Virtual Server (Nova)’s availability zone.
	AvailabilityZone string `json:"availability_zone,omitempty"`

	// Description is description
	Description string `json:"description,omitempty"`

	// LoadBalancerPlanID is the UUID of Load Balancer Plan.
	LoadBalancerPlanID string `json:"load_balancer_plan_id" required:"true"`

	// Name is a human-readable name of the Load Balancer.
	Name string `json:"name,omitempty"`

	// The UUID of the project who owns the Load Balancer. Only administrative users
	// can specify a project UUID other than their own.
	TenantID string `json:"tenant_id,omitempty"`
}

// ToLoadBalancerCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToLoadBalancerCreateMap() (map[string]interface{}, error) {
	b, err := eclcloud.BuildRequestBody(opts, "load_balancer")
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Create accepts a CreateOpts struct and creates a new Load Balancer using the values
// provided. You must remember to provide a valid LoadBalancerPlanID.
func Create(c *eclcloud.ServiceClient, opts CreateOpts) (r CreateResult) {
	b, err := opts.ToLoadBalancerCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(createURL(c), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{201},
	})
	return
}

// UpdateOpts represents the attributes used when updating an existing Load Balancer.
type UpdateOpts struct {

	// Description is description
	DefaultGateway *interface{} `json:"default_gateway,omitempty"`

	// Description is description
	Description *string `json:"description,omitempty"`

	// LoadBalancerPlanID is the UUID of Load Balancer Plan.
	LoadBalancerPlanID string `json:"load_balancer_plan_id,omitempty"`

	// Name is a human-readable name of the Load Balancer.
	Name *string `json:"name,omitempty"`
}

// ToLoadBalancerUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToLoadBalancerUpdateMap() (map[string]interface{}, error) {
	b, err := eclcloud.BuildRequestBody(opts, "load_balancer")
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Update accepts a UpdateOpts struct and updates an existing Load Balancer using the
// values provided.
func Update(c *eclcloud.ServiceClient, id string, opts UpdateOpts) (r UpdateResult) {
	b, err := opts.ToLoadBalancerUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(updateURL(c, id), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete accepts a unique ID and deletes the Load Balancer associated with it.
func Delete(c *eclcloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, id), nil)
	return
}

// IDFromName is a convenience function that returns a Load Balancer's ID,
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

	all, err := ExtractLoadBalancers(pages)
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
		return "", eclcloud.ErrResourceNotFound{Name: name, ResourceType: "load_balancer"}
	case 1:
		return id, nil
	default:
		return "", eclcloud.ErrMultipleResourcesFound{Name: name, Count: count, ResourceType: "load_balancer"}
	}
}
