package network_based_device_ha

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToHADeviceQuery() (string, error)
}

// ListOpts enables filtering of a list request.
type ListOpts struct {
	TenantID string `q:"tenant_id"`
	Locale   string `q:"locale"`
}

// ToHADeviceQuery formats a ListOpts into a query string.
func (opts ListOpts) ToHADeviceQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List enumerates the Devices to which the current token has access.
func List(client *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToHADeviceQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return HADevicePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToHADeviceCreateMap() (map[string]interface{}, error)
}

// GtHostInCreate represents parameters used to create a HA Device.
type GtHostInCreate struct {
	OperatingMode string `json:"operatingmode" required:"true"`
	LicenseKind   string `json:"licensekind" required:"true"`
	AZGroup       string `json:"azgroup" required:"true"`

	HALink1NetworkID string `json:"halink1networkid" required:"true"`
	HALink1SubnetID  string `json:"halink1subnetid" required:"true"`
	HALink1IPAddress string `json:"halink1ipaddress" required:"true"`

	HALink2NetworkID string `json:"halink2networkid" required:"true"`
	HALink2SubnetID  string `json:"halink2subnetid" required:"true"`
	HALink2IPAddress string `json:"halink2ipaddress" required:"true"`
}

// CreateOpts represents parameters used to create a device.
type CreateOpts struct {
	SOKind   string            `json:"sokind" required:"true"`
	TenantID string            `json:"tenant_id" required:"true"`
	Locale   string            `json:"locale,omitempty"`
	GtHost   [2]GtHostInCreate `json:"gt_host" required:"true"`
}

// ToHADeviceCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToHADeviceCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "")
}

// Create creates a new device.
func Create(client *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToHADeviceCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), &b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// DeleteOptsBuilder allows extensions to add additional parameters to
// the Delete request.
type DeleteOptsBuilder interface {
	ToHADeviceDeleteMap() (map[string]interface{}, error)
}

// GtHostInDelete represents parameters used to delete a HA Device.
type GtHostInDelete struct {
	HostName string `json:"hostname" required:"true"`
}

// DeleteOpts represents parameters used to delete a device.
type DeleteOpts struct {
	SOKind   string            `json:"sokind" required:"true"`
	TenantID string            `json:"tenant_id" required:"true"`
	GtHost   [2]GtHostInDelete `json:"gt_host" required:"true"`
}

// ToHADeviceDeleteMap formats a DeleteOpts into a delete request.
func (opts DeleteOpts) ToHADeviceDeleteMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "")
}

// Delete deletes a device.
func Delete(client *eclcloud.ServiceClient, deviceType string, opts DeleteOptsBuilder) (r DeleteResult) {
	b, err := opts.ToHADeviceDeleteMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), &b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return

}

// UpdateOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateOptsBuilder interface {
	ToHADeviceUpdateMap() (map[string]interface{}, error)
}

// GtHostInUpdate represents parameters used to update a HA Device.
type GtHostInUpdate struct {
	OperatingMode string `json:"operatingmode" required:"true"`
	LicenseKind   string `json:"licensekind" required:"true"`
	HostName      string `json:"hostname" required:"true"`
}

// UpdateOpts represents parameters to update a HA Device.
type UpdateOpts struct {
	SOKind   string            `json:"sokind" required:"true"`
	Locale   string            `json:"locale,omitempty"`
	TenantID string            `json:"tenant_id" required:"true"`
	GtHost   [2]GtHostInUpdate `json:"gt_host" required:"true"`
}

// ToHADeviceUpdateMap formats a UpdateOpts into an update request.
func (opts UpdateOpts) ToHADeviceUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "")
}

// Update modifies the attributes of a device.
func Update(client *eclcloud.ServiceClient, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToHADeviceUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(updateURL(client), &b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
