package servers

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToResourceListQuery() (string, error)
}

// ListOpts provides options to filter the List results.
type ListOpts struct {
	ChangesSince string `q:"changes-since"`
	Marker       string `q:"marker"`
	Limit        int    `q:"limit"`
	Name         string `q:"name"`
	Image        string `q:"image"`
	Flavor       string `q:"flavor"`
	Status       string `q:"status"`
}

// ToResourceListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToResourceListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

// List retrieves a list of Servers.
func List(client *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToResourceListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ServerPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// ListDetails retrieves a list of Servers in details.
func ListDetails(client *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listDetailsURL(client)
	if opts != nil {
		query, err := opts.ToResourceListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ServerPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves details of a Server.
func Get(client *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToResourceCreateMap() (map[string]interface{}, error)
}

// CreateOpts provides options used to create a Server.
type CreateOpts struct {
	Name             string            `json:"name"`
	Description      string            `json:"description,omitempty"`
	Networks         []Network         `json:"networks"`
	AdminPass        string            `json:"adminPass,omitempty"`
	ImageRef         string            `json:"imageRef"`
	FlavorRef        string            `json:"flavorRef"`
	AvailabilityZone string            `json:"availability_zone,omitempty"`
	Metadata         map[string]string `json:"metadata,omitempty"`
}

type Network struct {
	UUID           string `json:"uuid"`
	Port           string `json:"port,omitempty"`
	FixedIP        string `json:"fixed_ip,omitempty"`
	Plane          string `json:"plane"`
	SegmentationID int    `json:"segmentation_id"`
}

// ToResourceCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToResourceCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "server")
}

// Create creates a new Server.
func Create(client *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToResourceCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), &b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete deletes a Server.
func Delete(client *eclcloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, id), nil)
	return
}

type AddLicenseOpts struct {
	VmName       string   `json:"vm_name,omitempty"`
	VmID         string   `json:"vm_id,omitempty"`
	LicenseTypes []string `json:"license_types"`
}

func (opts AddLicenseOpts) ToResourceCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "add-license-to-vm")
}

func AddLicense(client *eclcloud.ServiceClient, serverID string, opts CreateOptsBuilder) (r AddLicenseResult) {
	b, err := opts.ToResourceCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(actionURL(client, serverID), &b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type GetAddLicenseResultOpts struct {
	JobID string `json:"job_id"`
}

func (opts GetAddLicenseResultOpts) ToResourceCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "get-result-for-add-license-to-vm")
}

func GetAddLicenseResult(client *eclcloud.ServiceClient, serverID string, opts CreateOptsBuilder) (r AddLicenseResult) {
	b, err := opts.ToResourceCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(actionURL(client, serverID), &b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
