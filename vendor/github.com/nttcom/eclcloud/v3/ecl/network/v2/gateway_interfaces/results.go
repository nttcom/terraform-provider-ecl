package gateway_interfaces

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

type commonResult struct {
	eclcloud.Result
}

func (r commonResult) Extract() (*GatewayInterface, error) {
	var s GatewayInterface
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "gw_interface")
}

type CreateResult struct {
	commonResult
}

type GetResult struct {
	commonResult
}

type UpdateResult struct {
	commonResult
}

type DeleteResult struct {
	eclcloud.ErrResult
}

type GatewayInterface struct {
	AwsGwID       string `json:"aws_gw_id"`
	AzureGwID     string `json:"azure_gw_id"`
	Description   string `json:"description"`
	FICGatewayID  string `json:"fic_gw_id"`
	GcpGwID       string `json:"gcp_gw_id"`
	GwVipv4       string `json:"gw_vipv4"`
	GwVipv6       string `json:"gw_vipv6"`
	ID            string `json:"id"`
	InterdcGwID   string `json:"interdc_gw_id"`
	InternetGwID  string `json:"internet_gw_id"`
	Name          string `json:"name"`
	Netmask       int    `json:"netmask"`
	NetworkID     string `json:"network_id"`
	PrimaryIpv4   string `json:"primary_ipv4"`
	PrimaryIpv6   string `json:"primary_ipv6"`
	SecondaryIpv4 string `json:"secondary_ipv4"`
	SecondaryIpv6 string `json:"secondary_ipv6"`
	ServiceType   string `json:"service_type"`
	Status        string `json:"status"`
	TenantID      string `json:"tenant_id"`
	VpnGwID       string `json:"vpn_gw_id"`
	VRID          int    `json:"vrid"`
}

type GatewayInterfacePage struct {
	pagination.LinkedPageBase
}

func (r GatewayInterfacePage) NextPageURL() (string, error) {
	var s struct {
		Links []eclcloud.Link `json:"gw_interfaces_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return eclcloud.ExtractNextURL(s.Links)
}

func (r GatewayInterfacePage) IsEmpty() (bool, error) {
	is, err := ExtractGatewayInterfaces(r)
	return len(is) == 0, err
}

func ExtractGatewayInterfaces(r pagination.Page) ([]GatewayInterface, error) {
	var s []GatewayInterface
	err := ExtractGatewayInterfacesInto(r, &s)
	return s, err
}

func ExtractGatewayInterfacesInto(r pagination.Page, v interface{}) error {
	return r.(GatewayInterfacePage).Result.ExtractIntoSlicePtr(v, "gw_interfaces")
}
