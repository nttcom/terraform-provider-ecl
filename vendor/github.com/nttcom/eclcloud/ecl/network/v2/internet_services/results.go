package internet_services

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

type commonResult struct {
	eclcloud.Result
}

func (r commonResult) Extract() (*InternetService, error) {
	var s InternetService
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "internet_service")
}

type GetResult struct {
	commonResult
}

type InternetService struct {
	Description          string `json:"description"`
	ID                   string `json:"id"`
	MinimalSubmaskLength int    `json:"minimal_submask_length"`
	Name                 string `json:"name"`
	Zone                 string `json:"zone"`
}

type InternetServicePage struct {
	pagination.LinkedPageBase
}

func (r InternetServicePage) NextPageURL() (string, error) {
	var s struct {
		Links []eclcloud.Link `json:"internet_services_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return eclcloud.ExtractNextURL(s.Links)
}

func (r InternetServicePage) IsEmpty() (bool, error) {
	is, err := ExtractInternetServices(r)
	return len(is) == 0, err
}

func ExtractInternetServices(r pagination.Page) ([]InternetService, error) {
	var s []InternetService
	err := ExtractInternetServicesInto(r, &s)
	return s, err
}

func ExtractInternetServicesInto(r pagination.Page, v interface{}) error {
	return r.(InternetServicePage).Result.ExtractIntoSlicePtr(v, "internet_services")
}
