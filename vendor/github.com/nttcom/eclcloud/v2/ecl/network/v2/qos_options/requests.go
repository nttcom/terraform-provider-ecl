package qos_options

import (
	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToQosOptionsListQuery() (string, error)
}

// ListOpts allows the filtering of paginated collections through the API.
// Filtering is achieved by passing in struct field values that map to
// the QoS option attributes you want to see returned. Marker and Limit are used
// for pagination.
type ListOpts struct {
	// Unique ID for the AWSService.
	AWSServiceID string `q:"aws_service_id"`

	// Unique ID for the AzureService.
	AzureServiceID string `q:"azure_service_id"`

	// Bandwidth assigned with this QoS Option
	Bandwidth string `q:"bandwidth"`

	// Description is the description of the QoS Policy.
	Description string `q:"description"`

	// Unique ID for the FICService.
	FICServiceID string `q:"fic_service_id"`

	// Unique ID for the GCPService.
	GCPServiceID string `q:"gcp_service_id"`

	// Unique ID for the QoS option.
	ID string `q:"id"`

	// Unique ID for the InterDCService.
	InterDCServiceID string `q:"interdc_service_id"`

	// Unique ID for the InternetService.
	InternetServiceID string `q:"internet_service_id"`

	// Name is the name of the QoS option.
	Name string `q:"name"`

	// Type of the QoS option.(guarantee or besteffort)
	QoSType string `q:"qos_type"`

	// Service type of the QoS option.(aws, azure, fic, gcp, vpn, internet, interdc)
	ServiceType string `q:"service_type"`

	// Indicates whether QoS option is currently operational.
	Status string `q:"status"`

	// Unique ID for the VPNService.
	VPNServiceID string `q:"vpn_service_id"`
}

// ToQosOptionsListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToQosOptionsListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List makes a request against the API to list QoS options accessible to you.
func List(client *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToQosOptionsListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return QosOptionPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific QoS option based on its unique ID.
func Get(client *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}
