package static_routes

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

type commonResult struct {
	eclcloud.Result
}

func (r commonResult) Extract() (*StaticRoute, error) {
	var s StaticRoute
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "static_route")
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

type StaticRoute struct {
	AwsGwID      string `json:"aws_gw_id"`
	AzureGwID    string `json:"azure_gw_id"`
	Description  string `json:"description"`
	Destination  string `json:"destination"`
	GcpGwID      string `json:"gcp_gw_id"`
	ID           string `json:"id"`
	InterdcGwID  string `json:"interdc_gw_id"`
	InternetGwID string `json:"internet_gw_id"`
	Name         string `json:"name"`
	Nexthop      string `json:"nexthop"`
	ServiceType  string `json:"service_type"`
	Status       string `json:"status"`
	TenantID     string `json:"tenant_id"`
	VpnGwID      string `json:"vpn_gw_id"`
}

type StaticRoutePage struct {
	pagination.LinkedPageBase
}

func (r StaticRoutePage) NextPageURL() (string, error) {
	var s struct {
		Links []eclcloud.Link `json:"static_routes_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return eclcloud.ExtractNextURL(s.Links)
}

func (r StaticRoutePage) IsEmpty() (bool, error) {
	is, err := ExtractStaticRoutes(r)
	return len(is) == 0, err
}

func ExtractStaticRoutes(r pagination.Page) ([]StaticRoute, error) {
	var s []StaticRoute
	err := ExtractStaticRoutesInto(r, &s)
	return s, err
}

func ExtractStaticRoutesInto(r pagination.Page, v interface{}) error {
	return r.(StaticRoutePage).Result.ExtractIntoSlicePtr(v, "static_routes")
}
