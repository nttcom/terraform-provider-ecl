package public_ips

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

type commonResult struct {
	eclcloud.Result
}

func (r commonResult) Extract() (*PublicIP, error) {
	var s PublicIP
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "public_ip")
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

type PublicIP struct {
	Cidr          string `json:"cidr"`
	Description   string `json:"description"`
	ID            string `json:"id"`
	InternetGwID  string `json:"internet_gw_id"`
	Name          string `json:"name"`
	Status        string `json:"status"`
	SubmaskLength int    `json:"submask_length"`
	TenantID      string `json:"tenant_id"`
}

type PublicIPPage struct {
	pagination.LinkedPageBase
}

func (r PublicIPPage) NextPageURL() (string, error) {
	var s struct {
		Links []eclcloud.Link `json:"public_ips_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return eclcloud.ExtractNextURL(s.Links)
}

func (r PublicIPPage) IsEmpty() (bool, error) {
	is, err := ExtractPublicIPs(r)
	return len(is) == 0, err
}

func ExtractPublicIPs(r pagination.Page) ([]PublicIP, error) {
	var s []PublicIP
	err := ExtractPublicIPsInto(r, &s)
	return s, err
}

func ExtractPublicIPsInto(r pagination.Page, v interface{}) error {
	return r.(PublicIPPage).Result.ExtractIntoSlicePtr(v, "public_ips")
}
