package load_balancer_syslog_servers

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// Extract is a function that accepts a result and extracts a Load Balancer Syslog Server resource.
func (r commonResult) Extract() (*LoadBalancerSyslogServer, error) {
	var s struct {
		LoadBalancerSyslogServer *LoadBalancerSyslogServer `json:"load_balancer_syslog_server"`
	}
	err := r.ExtractInto(&s)
	return s.LoadBalancerSyslogServer, err
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Load Balancer Syslog Server.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Load Balancer Syslog Server.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Load Balancer Syslog Server.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// LoadBalancerSyslogServer represents a Load Balancer Syslog Server. See package documentation for a top-level
// description of what this is.
type LoadBalancerSyslogServer struct {

	// should syslog record acl info
	AclLogging string `json:"acl_logging"`

	// should syslog record appflow info
	AppflowLogging string `json:"appflow_logging"`

	// date format utilized by syslog
	DateFormat string `json:"date_format"`

	// Description is description
	Description string `json:"description"`

	// UUID representing the Load Balancer Syslog Server.
	ID string `json:"id"`

	// Ip address of syslog server
	IPAddress string `json:"ip_address"`

	// The ID of load_balancer this load_balancer_syslog_server belongs to.
	LoadBalancerID string `json:"load_balancer_id"`

	// 	Log facility for syslog
	LogFacility string `json:"log_facility"`

	// Log level for syslog
	LogLevel string `json:"log_level"`

	// Name of the syslog resource
	Name string `json:"name"`

	// Port number of syslog server
	PortNumber int `json:"port_number"`

	// priority (0-255)
	Priority int `json:"priority"`

	// Load balancer syslog server status
	Status string `json:"status"`

	// should syslog record tcp protocol info
	TcpLogging string `json:"tcp_logging"`

	// TenantID is the project owner of the Load Balancer Syslog Server.
	TenantID string `json:"tenant_id"`

	// time zone utilized by syslog
	TimeZone string `json:"time_zone"`

	// protocol for syslog transport
	TransportType string `json:"transport_type"`

	// can user configure log messages
	UserConfigurableLogMessages string `json:"user_configurable_log_messages"`
}

// LoadBalancerSyslogServerPage is the page returned by a pager when traversing over a collection
// of load balancer Syslog Servers.
type LoadBalancerSyslogServerPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a LoadBalancerSyslogServerPage struct is empty.
func (r LoadBalancerSyslogServerPage) IsEmpty() (bool, error) {
	is, err := ExtractLoadBalancerSyslogServers(r)
	return len(is) == 0, err
}

// ExtractLoadBalancerSyslogServers accepts a Page struct, specifically a LoadBalancerSyslogServerPage struct,
// and extracts the elements into a slice of Load Balancer Syslog Server structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractLoadBalancerSyslogServers(r pagination.Page) ([]LoadBalancerSyslogServer, error) {
	var s struct {
		LoadBalancerSyslogServers []LoadBalancerSyslogServer `json:"load_balancer_syslog_servers"`
	}
	err := (r.(LoadBalancerSyslogServerPage)).ExtractInto(&s)
	return s.LoadBalancerSyslogServers, err
}
