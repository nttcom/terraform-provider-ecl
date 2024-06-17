package device_interfaces

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToDeviceInterfaceQuery() (string, error)
}

// ListOpts converts tenant id and token as query string
type ListOpts struct {
	TenantID  string `q:"tenantid"`
	UserToken string `q:"usertoken"`
}

// ToDeviceInterfaceQuery formats a ListOpts into a query string.
func (opts ListOpts) ToDeviceInterfaceQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// device interfaces.
func List(client *eclcloud.ServiceClient, serverUUID string, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client, serverUUID)
	if opts != nil {
		query, err := opts.ToDeviceInterfaceQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return DeviceInterfacePage{pagination.LinkedPageBase{PageResult: r}}
	})
}
