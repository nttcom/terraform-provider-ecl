package security_group_rules

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToSecurityGroupRuleListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the security group rule attributes you want to see returned.
type ListOpts struct {
	Description     string `q:"description"`
	Direction       string `q:"direction"`
	Ethertype       string `q:"ethertype"`
	ID              string `q:"id"`
	PortRangeMax    int    `q:"port_range_max"`
	PortRangeMin    int    `q:"port_range_min"`
	Protocol        string `q:"protocol"`
	RemoteGroupID   string `q:"remote_group_id"`
	RemoteIPPrefix  string `q:"remote_ip_prefix"`
	SecurityGroupID string `q:"security_group_id"`
	TenantID        string `q:"tenant_id"`
}

// ToSecurityGroupRuleListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToSecurityGroupRuleListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// security group rules. It accepts a ListOpts struct, which allows you to filter
// and sort the returned collection for greater efficiency.
func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToSecurityGroupRuleListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return SecurityGroupRulePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific security group rule based on its unique ID.
func Get(c *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, id), &r.Body, nil)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToSecurityGroupRuleCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents options used to create a security group rule.
type CreateOpts struct {
	Description     string  `json:"description,omitempty"`
	Direction       string  `json:"direction" required:"true"`
	Ethertype       string  `json:"ethertype,omitempty"`
	PortRangeMax    *int    `json:"port_range_max,omitempty"`
	PortRangeMin    *int    `json:"port_range_min,omitempty"`
	Protocol        string  `json:"protocol,omitempty"`
	RemoteGroupID   *string `json:"remote_group_id,omitempty"`
	RemoteIPPrefix  *string `json:"remote_ip_prefix,omitempty"`
	SecurityGroupID string  `json:"security_group_id" required:"true"`
	TenantID        string  `json:"tenant_id,omitempty"`
}

// ToSecurityGroupRuleCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToSecurityGroupRuleCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "security_group_rule")
}

// Create accepts a CreateOpts struct and creates a new security group rule.
func Create(c *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToSecurityGroupRuleCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(createURL(c), b, &r.Body, nil)
	return
}

// Delete accepts a unique ID and deletes the security group rule associated with it.
func Delete(c *eclcloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, id), nil)
	return
}
