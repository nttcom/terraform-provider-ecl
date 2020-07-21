package gateway_interfaces

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

type ListOptsBuilder interface {
	ToGatewayInterfaceListQuery() (string, error)
}

type ListOpts struct {
	AwsGwID       string `q:"aws_gw_id"`
	AzureGwID     string `q:"azure_gw_id"`
	Description   string `q:"description"`
	FICGatewayID  string `q:"fic_gw_id"`
	GcpGwID       string `q:"gcp_gw_id"`
	GwVipv4       string `q:"gw_vipv4"`
	GwVipv6       string `q:"gw_vipv6"`
	ID            string `q:"id"`
	InterdcGwID   string `q:"interdc_gw_id"`
	InternetGwID  string `q:"internet_gw_id"`
	Name          string `q:"name"`
	Netmask       int    `q:"netmask"`
	NetworkID     string `q:"network_id"`
	PrimaryIpv4   string `q:"primary_ipv4"`
	PrimaryIpv6   string `q:"primary_ipv6"`
	SecondaryIpv4 string `q:"secondary_ipv4"`
	SecondaryIpv6 string `q:"secondary_ipv6"`
	ServiceType   string `q:"service_type"`
	Status        string `q:"status"`
	TenantID      string `q:"tenant_id"`
	VpnGwID       string `q:"vpn_gw_id"`
	VRID          int    `q:"vrid"`
}

func (opts ListOpts) ToGatewayInterfaceListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToGatewayInterfaceListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return GatewayInterfacePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

func Get(c *eclcloud.ServiceClient, gatewayInterfaceID string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, gatewayInterfaceID), &r.Body, nil)
	return
}

type CreateOptsBuilder interface {
	ToGatewayInterfaceCreateMap() (map[string]interface{}, error)
}

type CreateOpts struct {
	AwsGwID       string `json:"aws_gw_id,omitempty"`
	AzureGwID     string `json:"azure_gw_id,omitempty"`
	Description   string `json:"description"`
	FICGatewayID  string `json:"fic_gw_id,omitempty"`
	GcpGwID       string `json:"gcp_gw_id,omitempty"`
	GwVipv4       string `json:"gw_vipv4" required:"true"`
	InterdcGwID   string `json:"interdc_gw_id,omitempty"`
	InternetGwID  string `json:"internet_gw_id,omitempty"`
	Name          string `json:"name"`
	Netmask       int    `json:"netmask" required:"true"`
	NetworkID     string `json:"network_id" required:"true"`
	PrimaryIpv4   string `json:"primary_ipv4" required:"true"`
	SecondaryIpv4 string `json:"secondary_ipv4" required:"true"`
	ServiceType   string `json:"service_type" required:"true"`
	TenantID      string `json:"tenant_id,omitempty"`
	VpnGwID       string `json:"vpn_gw_id,omitempty"`
	VRID          int    `json:"vrid" required:"true"`
}

func (opts CreateOpts) ToGatewayInterfaceCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "gw_interface")
}

func Create(c *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToGatewayInterfaceCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(createURL(c), b, &r.Body, nil)
	return
}

type UpdateOptsBuilder interface {
	ToGatewayInterfaceUpdateMap() (map[string]interface{}, error)
}

type UpdateOpts struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

func (opts UpdateOpts) ToGatewayInterfaceUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "gw_interface")
}

func Update(c *eclcloud.ServiceClient, gatewayInterfaceID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToGatewayInterfaceUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(updateURL(c, gatewayInterfaceID), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}

func Delete(c *eclcloud.ServiceClient, gatewayInterfaceID string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, gatewayInterfaceID), nil)
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

	all, err := ExtractGatewayInterfaces(pages)
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
		return "", eclcloud.ErrResourceNotFound{Name: name, ResourceType: "gw_interface"}
	case 1:
		return id, nil
	default:
		return "", eclcloud.ErrMultipleResourcesFound{Name: name, Count: count, ResourceType: "gw_interface"}
	}
}
