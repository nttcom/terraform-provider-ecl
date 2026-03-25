package security_groups

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

type commonResult struct {
	eclcloud.Result
}

func (r commonResult) Extract() (*SecurityGroup, error) {
	var s SecurityGroup
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "security_group")
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

// SecurityGroupRule represents a rule within a security group
type SecurityGroupRule struct {
	Description     string  `json:"description"`
	Direction       string  `json:"direction"`
	Ethertype       string  `json:"ethertype"`
	ID              string  `json:"id"`
	PortRangeMax    *int    `json:"port_range_max"`
	PortRangeMin    *int    `json:"port_range_min"`
	Protocol        string  `json:"protocol"`
	RemoteGroupID   *string `json:"remote_group_id"`
	RemoteIPPrefix  *string `json:"remote_ip_prefix"`
	SecurityGroupID string  `json:"security_group_id"`
	TenantID        string  `json:"tenant_id"`
}

// SecurityGroup represents a security group
type SecurityGroup struct {
	Description        string              `json:"description"`
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	SecurityGroupRules []SecurityGroupRule `json:"security_group_rules"`
	Status             string              `json:"status"`
	Tags               map[string]string   `json:"tags"`
	TenantID           string              `json:"tenant_id"`
}

type SecurityGroupPage struct {
	pagination.LinkedPageBase
}

func (r SecurityGroupPage) NextPageURL() (string, error) {
	var s struct {
		Links []eclcloud.Link `json:"security_groups_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return eclcloud.ExtractNextURL(s.Links)
}

func (r SecurityGroupPage) IsEmpty() (bool, error) {
	is, err := ExtractSecurityGroups(r)
	return len(is) == 0, err
}

func ExtractSecurityGroups(r pagination.Page) ([]SecurityGroup, error) {
	var s []SecurityGroup
	err := ExtractSecurityGroupsInto(r, &s)
	return s, err
}

func ExtractSecurityGroupsInto(r pagination.Page, v interface{}) error {
	return r.(SecurityGroupPage).Result.ExtractIntoSlicePtr(v, "security_groups")
}
