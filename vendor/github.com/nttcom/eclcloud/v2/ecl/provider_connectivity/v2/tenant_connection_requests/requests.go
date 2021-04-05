package tenant_connection_requests

import (
	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToTenantConnectionRequestListQuery() (string, error)
}

// ListOpts provides options to filter the List results.
type ListOpts struct {
	TenantConnectionRequestID string `q:"tenant_connection_request_id"`
	Status                    string `q:"status"`
	Name                      string `q:"name"`
	TenantID                  string `q:"tenant_id"`
	NameOther                 string `q:"name_other"`
	TenantIDOther             string `q:"tenant_id_other"`
	NetworkID                 string `q:"network_id"`
	ApprovalRequestID         string `q:"approval_request_id"`
}

// ToTenantConnectionRequestListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToTenantConnectionRequestListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

// List retrieves a list of Tenant Connection Requests.
func List(client *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToTenantConnectionRequestListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return TenantConnectionRequestPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves details of an Tenant Connection Request.
func Get(client *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToTenantConnectionRequestCreateMap() (map[string]interface{}, error)
}

// CreateOpts provides options used to create a Tenant Connection Request.
type CreateOpts struct {
	TenantIDOther string            `json:"tenant_id_other" required:"true"`
	NetworkID     string            `json:"network_id" required:"true"`
	Name          string            `json:"name,omitempty"`
	Description   string            `json:"description,omitempty"`
	Tags          map[string]string `json:"tags,omitempty"`
}

// ToTenantConnectionRequestCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToTenantConnectionRequestCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "tenant_connection_request")
}

// Create creates a new Tenant Connection Request.
func Create(client *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToTenantConnectionRequestCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), &b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete deletes a Tenant Connection Request.
func Delete(client *eclcloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, id), &eclcloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateOptsBuilder interface {
	ToTenantConnectionRequestUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents parameters to update a Tenant Connection Request.
type UpdateOpts struct {
	Name             *string            `json:"name,omitempty"`
	Description      *string            `json:"description,omitempty"`
	Tags             *map[string]string `json:"tags,omitempty"`
	NameOther        *string            `json:"name_other,omitempty"`
	DescriptionOther *string            `json:"description_other,omitempty"`
	TagsOther        *map[string]string `json:"tags_other,omitempty"`
}

// ToResourceUpdateCreateMap formats a UpdateOpts into an update request.
func (opts UpdateOpts) ToTenantConnectionRequestUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "tenant_connection_request")
}

// Update modifies the attributes of a Tenant Connection Request.
func Update(client *eclcloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToTenantConnectionRequestUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(updateURL(client, id), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
