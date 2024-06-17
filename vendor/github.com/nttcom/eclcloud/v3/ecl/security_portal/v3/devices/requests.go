package devices

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToDevicesQuery() (string, error)
}

// ListOpts enables filtering of a list request.
type ListOpts struct {
	TenantID  string `q:"tenantid"`
	UserToken string `q:"usertoken"`
}

// ToDevicesQuery formats a ListOpts into a query string.
func (opts ListOpts) ToDevicesQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over
// a collection of devices.
func List(client *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToDevicesQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return DevicePage{pagination.LinkedPageBase{PageResult: r}}
	})
}
