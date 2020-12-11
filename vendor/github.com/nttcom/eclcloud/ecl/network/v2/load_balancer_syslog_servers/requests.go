package load_balancer_syslog_servers

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the Load Balancer Syslog Server attributes you want to see returned. SortKey allows you to sort
// by a particular Load Balancer Syslog Server attribute. SortDir sets the direction, and is either
// `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	Description    string `q:"description"`
	ID             string `q:"id"`
	IPAddress      string `q:"ip_address"`
	LoadBalancerID string `q:"load_balancer_id"`
	LogFacility    string `q:"log_facility"`
	LogLevel       string `q:"log_level"`
	Name           string `q:"name"`
	PortNumber     int    `q:"port_number"`
	Status         string `q:"status"`
	TransportType  string `q:"transport_type"`
}

// ToLoadBalancerSyslogServersListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToLoadBalancerSyslogServersListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// Load Balancer Syslog Servers. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those Load Balancer Syslog Servers that are owned by the tenant
// who submits the request, unless the request is submitted by a user with
// administrative rights.
func List(c *eclcloud.ServiceClient, opts ListOpts) pagination.Pager {
	url := listURL(c)
	query, err := opts.ToLoadBalancerSyslogServersListQuery()
	if err != nil {
		return pagination.Pager{Err: err}
	}
	url += query
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return LoadBalancerSyslogServerPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific Load Balancer Syslog Server based on its unique ID.
func Get(c *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, id), &r.Body, nil)
	return
}

// CreateOpts represents the attributes used when creating a new Load Balancer Syslog Server.
type CreateOpts struct {

	// should syslog record acl info
	AclLogging string `json:"acl_logging,omitempty"`

	// should syslog record appflow info
	AppflowLogging string `json:"appflow_logging,omitempty"`

	// date format utilized by syslog
	DateFormat string `json:"date_format,omitempty"`

	// Description is description
	Description string `json:"description,omitempty"`

	// Ip address of syslog server
	IPAddress string `json:"ip_address" required:"true"`

	// The ID of load_balancer this load_balancer_syslog_server belongs to.
	LoadBalancerID string `json:"load_balancer_id" required:"true"`

	// 	Log facility for syslog
	LogFacility string `json:"log_facility,omitempty"`

	// Log level for syslog
	LogLevel string `json:"log_level,omitempty"`

	// Name is a human-readable name of the Load Balancer Syslog Server.
	Name string `json:"name" required:"true"`

	// Port number of syslog server
	PortNumber int `json:"port_number,omitempty"`

	// priority (0-255)
	Priority *int `json:"priority,omitempty"`

	// should syslog record tcp protocol info
	TcpLogging string `json:"tcp_logging,omitempty"`

	// The UUID of the project who owns the Load Balancer Syslog Server. Only administrative users
	// can specify a project UUID other than their own.
	TenantID string `json:"tenant_id,omitempty"`

	// time zone utilized by syslog
	TimeZone string `json:"time_zone,omitempty"`

	// protocol for syslog transport
	TransportType string `json:"transport_type,omitempty"`

	// can user configure log messages
	UserConfigurableLogMessages string `json:"user_configurable_log_messages,omitempty"`
}

// ToLoadBalancerSyslogServerCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToLoadBalancerSyslogServerCreateMap() (map[string]interface{}, error) {
	b, err := eclcloud.BuildRequestBody(opts, "load_balancer_syslog_server")
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Create accepts a CreateOpts struct and creates a new Load Balancer Syslog Server using the values
// provided. You must remember to provide a valid LoadBalancerPlanID.
func Create(c *eclcloud.ServiceClient, opts CreateOpts) (r CreateResult) {
	b, err := opts.ToLoadBalancerSyslogServerCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(createURL(c), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{201},
	})
	return
}

// UpdateOpts represents the attributes used when updating an existing Load Balancer Syslog Server.
type UpdateOpts struct {

	// should syslog record acl info
	AclLogging string `json:"acl_logging,omitempty"`

	// should syslog record appflow info
	AppflowLogging string `json:"appflow_logging,omitempty"`

	// date format utilized by syslog
	DateFormat string `json:"date_format,omitempty"`

	// Description is description
	Description *string `json:"description,omitempty"`

	// 	Log facility for syslog
	LogFacility string `json:"log_facility,omitempty"`

	// Log level for syslog
	LogLevel string `json:"log_level,omitempty"`

	// priority (0-255)
	Priority *int `json:"priority,omitempty"`

	// should syslog record tcp protocol info
	TcpLogging string `json:"tcp_logging,omitempty"`

	// time zone utilized by syslog
	TimeZone string `json:"time_zone,omitempty"`

	// can user configure log messages
	UserConfigurableLogMessages string `json:"user_configurable_log_messages,omitempty"`
}

// ToLoadBalancerSyslogServerUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToLoadBalancerSyslogServerUpdateMap() (map[string]interface{}, error) {
	b, err := eclcloud.BuildRequestBody(opts, "load_balancer_syslog_server")
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Update accepts a UpdateOpts struct and updates an existing Load Balancer Syslog Server using the
// values provided.
func Update(c *eclcloud.ServiceClient, id string, opts UpdateOpts) (r UpdateResult) {
	b, err := opts.ToLoadBalancerSyslogServerUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(updateURL(c, id), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete accepts a unique ID and deletes the Load Balancer Syslog Server associated with it.
func Delete(c *eclcloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, id), nil)
	return
}

// IDFromName is a convenience function that returns a Load Balancer Syslog Server's ID,
// given its name.
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

	all, err := ExtractLoadBalancerSyslogServers(pages)
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
		return "", eclcloud.ErrResourceNotFound{Name: name, ResourceType: "load_balancer_syslog_server"}
	case 1:
		return id, nil
	default:
		return "", eclcloud.ErrMultipleResourcesFound{Name: name, Count: count, ResourceType: "load_balancer_syslog_server"}
	}
}
