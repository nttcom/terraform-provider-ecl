package approval_requests

import (
	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToApprovalRequestListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the approval request attributes you want to see returned.
type ListOpts struct {
	Status  string `q:"status"`
	Service string `q:"service"`
}

// ToApprovalRequestListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToApprovalRequestListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List retrieves a list of approval requests.
func List(client *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToApprovalRequestListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ApprovalRequestPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves details of an approval request.
func Get(client *eclcloud.ServiceClient, name string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, name), &r.Body, nil)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateOptsBuilder interface {
	ToResourceUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents parameters to update an approval request.
type UpdateOpts struct {
	Status string `json:"status" required:"true"`
}

// ToResourceUpdateMap formats a UpdateOpts to update approval request.
func (opts UpdateOpts) ToResourceUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "")
}

// Update modifies the attributes of an approval request.
func Update(client *eclcloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToResourceUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(updateURL(client, id), b, nil, &eclcloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}
