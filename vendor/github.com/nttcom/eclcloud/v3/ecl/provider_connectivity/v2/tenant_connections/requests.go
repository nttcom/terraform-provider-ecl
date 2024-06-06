package tenant_connections

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToTenantConnectionListQuery() (string, error)
}

// ListOpts provides options to filter the List results.
type ListOpts struct {
	TenantConnectionRequestID string `q:"tenant_connection_request_id"`
	Status                    string `q:"status"`
	Name                      string `q:"name"`
	TenantID                  string `q:"tenant_id"`
	NameOther                 string `q:"name_other"`
	TenantIDOther             string `q:"tenant_id_other"`
	NetworkID                 string `q:"network_id"`
	DeviceType                string `q:"device_type"`
	DeviceID                  string `q:"device_id"`
	DeviceInterfaceID         string `q:"device_interface_id"`
	PortID                    string `q:"port_id"`
}

// ToTenantConnectionListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToTenantConnectionListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

// List retrieves a list of Tenant Connection.
func List(client *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToTenantConnectionListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return TenantConnectionPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves details of an Tenant Connection.
func Get(client *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToTenantConnectionCreateMap() (map[string]interface{}, error)
}

// ServerFixedIPs contains the IP Address and the SubnetID.
type ServerFixedIPs struct {
	SubnetID  string `json:"subnet_id,omitempty"`
	IPAddress string `json:"ip_address,omitempty"`
}

// AddressPair contains the IP Address and the MAC address.
type AddressPair struct {
	IPAddress  string `json:"ip_address,omitempty"`
	MACAddress string `json:"mac_address,omitempty"`
}

// VnaFixedIPs represents ip address part of virtual network appliance.
type VnaFixedIPs struct {
	IPAddress string `json:"ip_address,omitempty"`
}

// Vna represents the parameter when device_type is VSRX.
type Vna struct {
	FixedIPs []VnaFixedIPs `json:"fixed_ips,omitempty"`
}

// ComputeServer represents the parameter when device_type is Compute Server.
type ComputeServer struct {
	AllowedAddressPairs []AddressPair    `json:"allowed_address_pairs,omitempty"`
	FixedIPs            []ServerFixedIPs `json:"fixed_ips,omitempty"`
}

// BaremetalServer represents the parameter when device_type is Baremetal Server.
type BaremetalServer struct {
	AllowedAddressPairs []AddressPair    `json:"allowed_address_pairs,omitempty"`
	FixedIPs            []ServerFixedIPs `json:"fixed_ips,omitempty"`
	SegmentationID      int              `json:"segmentation_id,omitempty"`
	SegmentationType    string           `json:"segmentation_type,omitempty"`
}

// CreateOpts provides options used to create a Tenant Connection.
type CreateOpts struct {
	Name                      string            `json:"name,omitempty"`
	Description               string            `json:"description,omitempty"`
	Tags                      map[string]string `json:"tags,omitempty"`
	TenantConnectionRequestID string            `json:"tenant_connection_request_id" required:"true"`
	DeviceType                string            `json:"device_type" required:"true"`
	DeviceID                  string            `json:"device_id" required:"true"`
	DeviceInterfaceID         string            `json:"device_interface_id,omitempty"`
	AttachmentOpts            interface{}       `json:"attachment_opts,omitempty"`
}

// ToTenantConnectionCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToTenantConnectionCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "tenant_connection")
}

// Create creates a new Tenant Connection.
func Create(client *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToTenantConnectionCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), &b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete deletes a Tenant Connection.
func Delete(client *eclcloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, id), &eclcloud.RequestOpts{
		OkCodes: []int{204},
	})
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateOptsBuilder interface {
	ToTenantConnectionUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents parameters to update a Tenant Connection.
type UpdateOpts struct {
	Name             *string            `json:"name,omitempty"`
	Description      *string            `json:"description,omitempty"`
	Tags             *map[string]string `json:"tags,omitempty"`
	NameOther        *string            `json:"name_other,omitempty"`
	DescriptionOther *string            `json:"description_other,omitempty"`
	TagsOther        *map[string]string `json:"tags_other,omitempty"`
}

// ToTenantConnectionUpdateMap formats a UpdateOpts into an update request.
func (opts UpdateOpts) ToTenantConnectionUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "tenant_connection")
}

// Update modifies the attributes of a Tenant Connection.
func Update(client *eclcloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToTenantConnectionUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(updateURL(client, id), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
