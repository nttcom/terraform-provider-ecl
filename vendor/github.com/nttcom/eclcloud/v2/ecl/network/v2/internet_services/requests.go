package internet_services

import (
	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/pagination"
)

type ListOptsBuilder interface {
	ToInternetServiceListQuery() (string, error)
}

type ListOpts struct {
	Description          string `q:"description"`
	ID                   string `q:"id"`
	MinimalSubmaskLength int    `q:"minimal_submask_length"`
	Name                 string `q:"name"`
	Zone                 string `q:"zone"`
}

func (opts ListOpts) ToInternetServiceListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToInternetServiceListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return InternetServicePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

func Get(c *eclcloud.ServiceClient, internetServiceID string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, internetServiceID), &r.Body, nil)
	return
}
