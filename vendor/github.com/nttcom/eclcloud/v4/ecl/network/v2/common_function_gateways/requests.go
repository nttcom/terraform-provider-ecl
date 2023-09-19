package common_function_gateways

import (
	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToCommonFunctionGatewayListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the common function gateway attributes you want to see returned.
type ListOpts struct {
	CommonFunctionPoolID string `q:"common_function_pool_id"`
	Description          string `q:"description"`
	ID                   string `q:"id"`
	Name                 string `q:"name"`
	NetworkID            string `q:"network_id"`
	Status               string `q:"status"`
	SubnetID             string `q:"subnet_id"`
	TenantID             string `q:"tenant_id"`
}

// ToCommonFunctionGatewayListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToCommonFunctionGatewayListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// common function gateways.
// It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToCommonFunctionGatewayListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return CommonFunctionGatewayPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific common function gateway based on its unique ID.
func Get(c *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, id), &r.Body, nil)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToCommonFunctionGatewayCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents options used to create a common function gateway.
type CreateOpts struct {
	Name                 string `json:"name,omitempty"`
	Description          string `json:"description,omitempty"`
	CommonFunctionPoolID string `json:"common_function_pool_id,omitempty"`
	TenantID             string `json:"tenant_id,omitempty"`
}

// ToCommonFunctionGatewayCreateMap builds a request body from CreateOpts.
// func (opts CreateOpts) ToCommonFunctionGatewayCreateMap() (map[string]interface{}, error) {
func (opts CreateOpts) ToCommonFunctionGatewayCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "common_function_gateway")
}

// Create accepts a CreateOpts struct and creates a new common function gateway
// using the values provided.
// This operation does not actually require a request body, i.e. the
// CreateOpts struct argument can be empty.
//
// The tenant ID that is contained in the URI is the tenant that creates the
// common function gateway.
// An admin user, however, has the option of specifying another tenant
// ID in the CreateOpts struct.
func Create(c *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToCommonFunctionGatewayCreateMap()
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
	ToCommonFunctionGatewayUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents options used to update a common function gateway.
type UpdateOpts struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// ToCommonFunctionGatewayUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToCommonFunctionGatewayUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "common_function_gateway")
}

// Update accepts a UpdateOpts struct and updates an existing common function gateway
// using the values provided. For more information, see the Create function.
func Update(c *eclcloud.ServiceClient, commonFunctionGatewayID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToCommonFunctionGatewayUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(updateURL(c, commonFunctionGatewayID), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}

// Delete accepts a unique ID and deletes the common function gateway associated with it.
func Delete(c *eclcloud.ServiceClient, commonFunctionGatewayID string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, commonFunctionGatewayID), nil)
	return
}

// IDFromName is a convenience function that returns a common function gateway's
// ID, given its name.
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

	all, err := ExtractCommonFunctionGateways(pages)
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
		return "", eclcloud.ErrResourceNotFound{Name: name, ResourceType: "common_function_gateway"}
	case 1:
		return id, nil
	default:
		return "", eclcloud.ErrMultipleResourcesFound{Name: name, Count: count, ResourceType: "common_function_gateway"}
	}
}
