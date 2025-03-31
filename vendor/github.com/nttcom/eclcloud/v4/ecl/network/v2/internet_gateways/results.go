package internet_gateways

import (
	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/pagination"
)

type commonResult struct {
	eclcloud.Result
}

func (r commonResult) Extract() (*InternetGateway, error) {
	var s InternetGateway
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "internet_gateway")
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

type InternetGateway struct {
	ID                string `json:"id"`
	Description       string `json:"description"`
	InternetServiceID string `json:"internet_service_id"`
	Name              string `json:"name"`
	QoSOptionID       string `json:"qos_option_id"`
	Status            string `json:"status"`
	TenantID          string `json:"tenant_id"`
}

type InternetGatewayPage struct {
	pagination.LinkedPageBase
}

func (r InternetGatewayPage) NextPageURL() (string, error) {
	var s struct {
		Links []eclcloud.Link `json:"internet_gateways_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return eclcloud.ExtractNextURL(s.Links)
}

func (r InternetGatewayPage) IsEmpty() (bool, error) {
	is, err := ExtractInternetGateways(r)
	return len(is) == 0, err
}

func ExtractInternetGateways(r pagination.Page) ([]InternetGateway, error) {
	var s []InternetGateway
	err := ExtractInternetGatewaysInto(r, &s)
	return s, err
}

func ExtractInternetGatewaysInto(r pagination.Page, v interface{}) error {
	return r.(InternetGatewayPage).Result.ExtractIntoSlicePtr(v, "internet_gateways")
}
