package tenants

import (
	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToTenantListQuery() (string, error)
}

// ListOpts enables filtering of a list request.
// Currently SSS Tenant API does not support any of query parameters.
type ListOpts struct {
}

// ToTenantListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToTenantListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List enumerates the Projects to which the current token has access.
func List(client *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToTenantListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return TenantPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves details on a single tenant, by ID.
func Get(client *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToTenantCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents parameters used to create a tenant.
type CreateOpts struct {
	// Workspace ID.
	WorkspaceID string `json:"workspace_id" required:"true"`
	// TenantRegion of the tenant.
	TenantRegion string `json:"region" required:"true"`
}

// ToTenantCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToTenantCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "")
}

// Create creates a new Project.
func Create(client *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToTenantCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), &b, &r.Body, nil)
	return
}
