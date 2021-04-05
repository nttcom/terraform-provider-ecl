package public_ips

import (
	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/pagination"
)

type ListOptsBuilder interface {
	ToPublicIPListQuery() (string, error)
}

type ListOpts struct {
	Cidr          string `q:"cidr"`
	Description   string `q:"description"`
	ID            string `q:"id"`
	InternetGwID  string `q:"internet_gw_id"`
	Name          string `q:"name"`
	Status        string `q:"status"`
	SubmaskLength int    `q:"submask_length"`
	TenantID      string `q:"tenant_id"`
}

func (opts ListOpts) ToPublicIPListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToPublicIPListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return PublicIPPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

func Get(c *eclcloud.ServiceClient, publicIPID string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, publicIPID), &r.Body, nil)
	return
}

type CreateOptsBuilder interface {
	ToPublicIPCreateMap() (map[string]interface{}, error)
}

type CreateOpts struct {
	Description   string `json:"description,omitempty"`
	InternetGwID  string `json:"internet_gw_id" required:"true"`
	Name          string `json:"name,omitempty"`
	SubmaskLength int    `json:"submask_length" required:"true"`
	TenantID      string `json:"tenant_id,omitempty"`
}

func (opts CreateOpts) ToPublicIPCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "public_ip")
}

func Create(c *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToPublicIPCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(createURL(c), b, &r.Body, nil)
	return
}

type UpdateOptsBuilder interface {
	ToPublicIPUpdateMap() (map[string]interface{}, error)
}

type UpdateOpts struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

func (opts UpdateOpts) ToPublicIPUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "public_ip")
}

func Update(c *eclcloud.ServiceClient, publicIPID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToPublicIPUpdateMap()
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

	all, err := ExtractPublicIPs(pages)
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
		return "", eclcloud.ErrResourceNotFound{Name: name, ResourceType: "public_ip"}
	case 1:
		return id, nil
	default:
		return "", eclcloud.ErrMultipleResourcesFound{Name: name, Count: count, ResourceType: "public_ip"}
	}
}
