package operations

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

/*
List Operations
*/

// ListOpts allows the filtering and sorting of paginated collections through the API.
// Filtering is achieved by passing in struct field values that map to the operation attributes you want to see returned.
type ListOpts struct {

	// - ID of the resource
	ID string `q:"id"`

	// - ID of the resource
	ResourceID string `q:"resource_id"`

	// - Type of the resource
	ResourceType string `q:"resource_type"`

	// - The unique hyphenated UUID to identify the request
	//   - The UUID which has been set by `X-MVNA-Request-Id` in request headers
	RequestID string `q:"request_id"`

	// - Type of the request
	RequestType string `q:"request_type"`

	// - Operation status of the resource
	Status string `q:"status"`

	// - ID of the owner tenant of the resource
	TenantID string `q:"tenant_id"`

	// - If `true` is set, operations of deleted resource is not displayed
	NoDeleted bool `q:"no_deleted"`

	// - If `true` is set, only the latest operation of each resource is displayed
	Latest bool `q:"latest"`
}

// ToOperationListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToOperationListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)

	return q.String(), err
}

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToOperationListQuery() (string, error)
}

// List returns a Pager which allows you to iterate over a collection of operations.
// It accepts a ListOpts struct, which allows you to filter and sort the returned collection for greater efficiency.
func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)

	if opts != nil {
		query, err := opts.ToOperationListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}

		url += query
	}

	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return OperationPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

/*
Show Operation
*/

// Show retrieves a specific operation based on its unique ID.
func Show(c *eclcloud.ServiceClient, id string) (r ShowResult) {
	_, r.Err = c.Get(showURL(c, id), &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}
