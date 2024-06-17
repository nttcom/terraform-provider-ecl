package fic_gateways

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToFICGatewaysListQuery() (string, error)
}

// ListOpts allows the filtering of paginated collections through the API.
// Filtering is achieved by passing in struct field values that map to
// the FIC Gateway attributes you want to see returned. Marker and Limit are used
// for pagination.
type ListOpts struct {
	// Description of the FIC Gateway resource.
	Description string `q:"description"`

	// 	FIC Service instantiated by this Gateway.
	FICServiceID string `q:"fic_service_id"`

	//Unique ID of the FIC Gateway resource.
	ID string `q:"id"`

	//Name of the FIC Gateway resource.
	Name string `q:"name"`

	// Quality of Service options selected for this Gateway.
	QoSOptionID string `q:"qos_option_id"`

	// The FIC Gateway status.
	Status string `q:"status"`

	// 	Tenant ID of the owner (UUID).
	TenantID string `q:"tenant_id"`
}

// ToFICGatewaysListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToFICGatewaysListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List makes a request against the API to list FIC Gateways accessible to you.
func List(client *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToFICGatewaysListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return FICGatewayPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific FIC Gateway based on its unique ID.
func Get(client *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}
