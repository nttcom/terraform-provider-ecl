package internet_gateways

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

type ListOptsBuilder interface {
	ToInternetGatewayListQuery() (string, error)
}

type ListOpts struct {
	Description       string `q:"description"`
	ID                string `q:"id"`
	InternetServiceID string `q:"internet_service_id"`
	Name              string `q:"name"`
	QoSOptionID       string `q:"qos_option_id"`
	Status            string `q:"status"`
	TenantID          string `q:"tenant_id"`
}

func (opts ListOpts) ToInternetGatewayListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToInternetGatewayListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return InternetGatewayPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

func Get(c *eclcloud.ServiceClient, internetGatewayID string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, internetGatewayID), &r.Body, nil)
	return
}

type CreateOptsBuilder interface {
	ToInternetGatewayCreateMap() (map[string]interface{}, error)
}

type CreateOpts struct {
	Description       string `json:"description,omitempty"`
	InternetServiceID string `json:"internet_service_id" required:"true"`
	Name              string `json:"name,omitempty"`
	QoSOptionID       string `json:"qos_option_id" required:"true"`
	TenantID          string `json:"tenant_id,omitempty"`
}

func (opts CreateOpts) ToInternetGatewayCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "internet_gateway")
}

func Create(c *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToInternetGatewayCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(createURL(c), b, &r.Body, nil)
	return
}

type UpdateOptsBuilder interface {
	ToInternetGatewayUpdateMap() (map[string]interface{}, error)
}

type UpdateOpts struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	QoSOptionID *string `json:"qos_option_id,omitempty"`
}

func (opts UpdateOpts) ToInternetGatewayUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "internet_gateway")
}

func Update(c *eclcloud.ServiceClient, internetGatewayID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToInternetGatewayUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(updateURL(c, internetGatewayID), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}

func Delete(c *eclcloud.ServiceClient, internetGatewayID string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, internetGatewayID), nil)
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

	all, err := ExtractInternetGateways(pages)
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
		return "", eclcloud.ErrResourceNotFound{Name: name, ResourceType: "internet_gateway"}
	case 1:
		return id, nil
	default:
		return "", eclcloud.ErrMultipleResourcesFound{Name: name, Count: count, ResourceType: "internet_gateway"}
	}
}
