package processes

import (
	"github.com/nttcom/eclcloud"
)

// GetOptsBuilder allows extensions to add additional parameters to
// the order API request
type GetOptsBuilder interface {
	ToProcessQuery() (string, error)
}

// GetOpts represents result of order API response.
type GetOpts struct {
	TenantID  string `q:"tenantid"`
	UserToken string `q:"usertoken"`
}

// ToProcessQuery formats a GetOpts into a query string.
func (opts GetOpts) ToProcessQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// Get retrieves details on a single order, by ID.
func Get(client *eclcloud.ServiceClient, processID string, opts GetOptsBuilder) (r GetResult) {
	url := getURL(client, processID)
	if opts != nil {
		query, _ := opts.ToProcessQuery()
		url += query
	}

	_, r.Err = client.Get(url, &r.Body, nil)
	return
}
