package load_balancer_plans

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the Load Balancer Plan attributes you want to see returned. SortKey allows you to sort
// by a particular Load Balancer Plan attribute. SortDir sets the direction, and is either
// `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	Description          string `q:"description"`
	ID                   string `q:"id"`
	MaximumSyslogServers int    `q:"maximum_syslog_servers"`
	Name                 string `q:"name"`
	Vendor               string `q:"vendor"`
	Version              string `q:"version"`
}

// ToLoadBalancerPlansListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToLoadBalancerPlansListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// Load Balancer Plans. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those Load Balancer Plans that are owned by the tenant
// who submits the request, unless the request is submitted by a user with
// administrative rights.
func List(c *eclcloud.ServiceClient, opts ListOpts) pagination.Pager {
	url := listURL(c)
	query, err := opts.ToLoadBalancerPlansListQuery()
	if err != nil {
		return pagination.Pager{Err: err}
	}
	url += query
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return LoadBalancerPlanPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific Load Balancer Plan based on its unique ID.
func Get(c *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, id), &r.Body, nil)
	return
}

// IDFromName is a convenience function that returns a Load Balancer Plan's ID,
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

	all, err := ExtractLoadBalancerPlans(pages)
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
		return "", eclcloud.ErrResourceNotFound{Name: name, ResourceType: "load_balancer_plan"}
	case 1:
		return id, nil
	default:
		return "", eclcloud.ErrMultipleResourcesFound{Name: name, Count: count, ResourceType: "load_balancer_plan"}
	}
}
