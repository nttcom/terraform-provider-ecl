package security_groups

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToSecurityGroupListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the security group attributes you want to see returned.
type ListOpts struct {
	Description string `q:"description"`
	ID          string `q:"id"`
	Name        string `q:"name"`
	Status      string `q:"status"`
	TenantID    string `q:"tenant_id"`
}

// ToSecurityGroupListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToSecurityGroupListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// security groups. It accepts a ListOpts struct, which allows you to filter
// and sort the returned collection for greater efficiency.
func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToSecurityGroupListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return SecurityGroupPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific security group based on its unique ID.
func Get(c *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, id), &r.Body, nil)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToSecurityGroupCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents options used to create a security group.
type CreateOpts struct {
	Description string            `json:"description,omitempty"`
	Name        string            `json:"name,omitempty"`
	Tags        map[string]string `json:"tags,omitempty"`
	TenantID    string            `json:"tenant_id,omitempty"`
}

// ToSecurityGroupCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToSecurityGroupCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "security_group")
}

// Create accepts a CreateOpts struct and creates a new security group using
// the values provided. This operation does not actually require a request
// body, i.e. the CreateOpts struct argument can be empty.
func Create(c *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToSecurityGroupCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(createURL(c), b, &r.Body, nil)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToSecurityGroupUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents options used to update a security group.
type UpdateOpts struct {
	Description *string            `json:"description,omitempty"`
	Name        *string            `json:"name,omitempty"`
	Tags        *map[string]string `json:"tags,omitempty"`
}

// ToSecurityGroupUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToSecurityGroupUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "security_group")
}

// Update accepts a UpdateOpts struct and updates an existing security group
// using the values provided.
func Update(c *eclcloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToSecurityGroupUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(updateURL(c, id), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete accepts a unique ID and deletes the security group associated with it.
func Delete(c *eclcloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, id), nil)
	return
}
