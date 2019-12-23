package static_routes

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

type ListOptsBuilder interface {
	ToStaticRouteListQuery() (string, error)
}

type ListOpts struct {
	AwsGwID      string `q:"aws_gw_id"`
	AzureGwID    string `q:"azure_gw_id"`
	Description  string `q:"description"`
	Destination  string `q:"destination"`
	GcpGwID      string `q:"gcp_gw_id"`
	ID           string `q:"id"`
	InterdcGwID  string `q:"inter_dc_id"`
	InternetGwID string `q:"internet_gw_id"`
	Name         string `q:"name"`
	Nexthop      string `q:"nexthop"`
	ServiceType  string `q:"service_type"`
	Status       string `q:"status"`
	TenantID     string `q:"tenant_id"`
	VpnGwID      string `q:"vpn_gw_id"`
}

func (opts ListOpts) ToStaticRouteListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToStaticRouteListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return StaticRoutePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

func Get(c *eclcloud.ServiceClient, publicIPID string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, publicIPID), &r.Body, nil)
	return
}

type CreateOptsBuilder interface {
	ToStaticRouteCreateMap() (map[string]interface{}, error)
}

type CreateOpts struct {
	AwsGwID      string `json:"aws_gw_id,omitempty"`
	AzureGwID    string `json:"azure_gw_id,omitempty"`
	Description  string `json:"description"`
	Destination  string `json:"destination" required:"true"`
	GcpGwID      string `json:"gcp_gw_id,omitempty"`
	InterdcGwID  string `json:"inter_dc_id,omitempty"`
	InternetGwID string `json:"internet_gw_id,omitempty"`
	Name         string `json:"name"`
	Nexthop      string `json:"nexthop" required:"true"`
	ServiceType  string `json:"service_type" required:"true"`
	TenantID     string `json:"tenant_id,omitempty"`
	VpnGwID      string `json:"vpn_gw_id,omitempty"`
}

func (opts CreateOpts) ToStaticRouteCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "static_route")
}

func Create(c *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToStaticRouteCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(createURL(c), b, &r.Body, nil)
	return
}

type UpdateOptsBuilder interface {
	ToStaticRouteUpdateMap() (map[string]interface{}, error)
}

type UpdateOpts struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

func (opts UpdateOpts) ToStaticRouteUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "static_route")
}

func Update(c *eclcloud.ServiceClient, publicIPID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToStaticRouteUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(updateURL(c, publicIPID), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}

func Delete(c *eclcloud.ServiceClient, publicIPID string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, publicIPID), nil)
	return
}

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

	all, err := ExtractStaticRoutes(pages)
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
		return "", eclcloud.ErrResourceNotFound{Name: name, ResourceType: "static_route"}
	case 1:
		return id, nil
	default:
		return "", eclcloud.ErrMultipleResourcesFound{Name: name, Count: count, ResourceType: "static_route"}
	}
}
