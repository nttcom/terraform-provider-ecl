package system_updates

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

/*
List System Updates
*/

// ListOpts allows the filtering and sorting of paginated collections through the API.
// Filtering is achieved by passing in struct field values that map to the system update attributes you want to see returned.
type ListOpts struct {

	// - ID of the resource
	ID string `q:"id"`

	// - Name of the resource
	// - This field accepts single-byte characters only
	Name string `q:"name"`

	// - Description of the resource
	// - This field accepts single-byte characters only
	Description string `q:"description"`

	// - URL of announcement for the system update (for example, Knowledge Center news)
	Href string `q:"href"`

	// - Current revision for the system update
	CurrentRevision int `q:"current_revision"`

	// - Next revision for the system update
	NextRevision int `q:"next_revision"`

	// - Whether the system update can be applied to the load balancer
	Applicable bool `q:"applicable"`

	// - If `true` is set, only the latest resource is displayed
	Latest bool `q:"latest"`
}

// ToSystemUpdateListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToSystemUpdateListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)

	return q.String(), err
}

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToSystemUpdateListQuery() (string, error)
}

// List returns a Pager which allows you to iterate over a collection of system updates.
// It accepts a ListOpts struct, which allows you to filter and sort the returned collection for greater efficiency.
func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)

	if opts != nil {
		query, err := opts.ToSystemUpdateListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}

		url += query
	}

	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return SystemUpdatePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

/*
Show System Update
*/

// Show retrieves a specific system update based on its unique ID.
func Show(c *eclcloud.ServiceClient, id string) (r ShowResult) {
	_, r.Err = c.Get(showURL(c, id), &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}
