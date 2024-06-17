package service_order_status

import (
	"github.com/nttcom/eclcloud/v3"
)

// GetOptsBuilder allows extensions to add additional parameters to
// the order progress API request
type GetOptsBuilder interface {
	ToServiceOrderQuery() (string, error)
}

// GetOpts represents result of order progress API response.
type GetOpts struct {
	TenantID string `q:"tenant_id"`
	Locale   string `q:"locale"`
	SoID     string `q:"soid"`
}

// ToServiceOrderQuery formats a GetOpts into a query string.
func (opts GetOpts) ToServiceOrderQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// Get retrieves details of an order progress, by SoId.
func Get(client *eclcloud.ServiceClient, deviceType string, opts GetOptsBuilder) (r GetResult) {
	url := getURL(client, deviceType)
	if opts != nil {
		query, _ := opts.ToServiceOrderQuery()
		url += query
	}

	_, r.Err = client.Get(url, &r.Body, nil)
	return
}
