package security_group_rules

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

type commonResult struct {
	eclcloud.Result
}

func (r commonResult) Extract() (*SecurityGroupRule, error) {
	var s SecurityGroupRule
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "security_group_rule")
}

type CreateResult struct {
	commonResult
}

type GetResult struct {
	commonResult
}

type DeleteResult struct {
	eclcloud.ErrResult
}

// SecurityGroupRule represents a security group rule
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

type SecurityGroupRulePage struct {
	pagination.LinkedPageBase
}

func (r SecurityGroupRulePage) NextPageURL() (string, error) {
	var s struct {
		Links []eclcloud.Link `json:"security_group_rules_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return eclcloud.ExtractNextURL(s.Links)
}

func (r SecurityGroupRulePage) IsEmpty() (bool, error) {
	is, err := ExtractSecurityGroupRules(r)
	return len(is) == 0, err
}

func ExtractSecurityGroupRules(r pagination.Page) ([]SecurityGroupRule, error) {
	var s []SecurityGroupRule
	err := ExtractSecurityGroupRulesInto(r, &s)
	return s, err
}

func ExtractSecurityGroupRulesInto(r pagination.Page, v interface{}) error {
	return r.(SecurityGroupRulePage).Result.ExtractIntoSlicePtr(v, "security_group_rules")
}
