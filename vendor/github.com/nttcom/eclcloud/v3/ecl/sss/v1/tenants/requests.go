package tenants

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
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
	// Name of this tenant.
	TenantName string `json:"tenant_name" required:"true"`

	// Description of the tenant.
	Description string `json:"description" required:"true"`

	// TenantRegion of the tenant.
	TenantRegion string `json:"region" required:"true"`

	// ID of contract which new tenant belongs.
	ContractID string `json:"contract_id,omitempty"`
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

// Delete deletes a tenant.
func Delete(client *eclcloud.ServiceClient, tenantID string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, tenantID), nil)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateOptsBuilder interface {
	ToTenantUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents parameters to update a tenant.
type UpdateOpts struct {
	// Description of the tenant.
	Description *string `json:"description,omitempty"`
}

// ToTenantUpdateMap formats a UpdateOpts into an update request.
func (opts UpdateOpts) ToTenantUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "")
}

// Update modifies the attributes of a tenant.
// SSS Tenant PUT API does not have response body,
// so set JSONResponse option as nil.
func Update(client *eclcloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToTenantUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(
		updateURL(client, id),
		b,
		nil,
		&eclcloud.RequestOpts{
			OkCodes: []int{204},
		},
	)
	return
}
