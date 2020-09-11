package volumes

import (
	"github.com/nttcom/eclcloud"
	// "github.com/nttcom/eclcloud/ecl/storage/v1/virtualstorages"
	// "github.com/nttcom/eclcloud/ecl/storage/v1/volumetypes"
	"github.com/nttcom/eclcloud/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToVolumeCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains options for creating a Volume. This object is passed to
// the Volumes.Create function. For more information about these parameters,
// see the Volume object.
type CreateOpts struct {
	// The volume name
	Name string `json:"name" required:"true"`
	// The volume description
	Description string `json:"description,omitempty"`
	// The volume size
	Size int `json:"size" required:"true"`
	// The volume IOPS as IOPS/GB
	IOPSPerGB string `json:"iops_per_gb,omitempty"`
	// The volume Throughput
	Throughput string `json:"throughput,omitempty"`
	// The initiator_iqns for volume (in case ISCSI)
	InitiatorIQNs []string `json:"initiator_iqns,omitempty"`
	// The availability zone of volume
	AvailabilityZone string `json:"availability_zone,omitempty"`
	// The parent virtual storage ID to connect volume
	VirtualStorageID string `json:"virtual_storage_id" required:"true"`
}

// ToVolumeCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToVolumeCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "volume")
}

// Create will create a new Volume based on the values in CreateOpts.
// To extract the Volume object from the response, call the Extract method on the
// CreateResult.
func Create(client *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToVolumeCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

// Delete will delete the existing Volume with the provided ID.
func Delete(client *eclcloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(
		deleteURL(client, id),
		&eclcloud.RequestOpts{
			OkCodes: []int{200},
		})
	return
}

// Get retrieves the Volume with the provided ID.
// To extract the Volume object from the response,
// call the Extract method on the GetResult.
func Get(client *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToVolumeListQuery() (string, error)
}

// ListOpts holds options for listing Volumes.
// It is passed to the Volumes.List function.
type ListOpts struct {
	// Now there are no definitions as query params in API specification
	// But do not remove this struct in future specification change.
}

// ToVolumeListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToVolumeListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns Volume optionally limited by the conditions provided in ListOpts.
func List(client *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToVolumeListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return VolumePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToVolumeUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contain options for updating an existing Volume.
// This object is passed to the volume.Update function.
// For more information about the parameters, see the Volume object.
type UpdateOpts struct {
	Name          *string   `json:"name,omitempty"`
	Description   *string   `json:"description,omitempty"`
	InitiatorIQNs *[]string `json:"initiator_iqns,omitempty"`
}

// ToVolumeUpdateMap assembles a request body based on the contents of an
// UpdateOpts.
// Volume of Storage SDP only allows to send "initiator_iqns" when
// the service type is "File Storage" type
// So in "ToVolumeUpdateMap" function, check volume type of virtual storage
// related to volume first
// And if service type is "File Storage's one", add initiator_iqns as request parameter
func (opts UpdateOpts) ToVolumeUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "volume")
}

// Update will update the Volume with provided information.
// To extract the updated Volume from the response,
// call the Extract method on the UpdateResult.
func Update(client *eclcloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToVolumeUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(updateURL(client, id), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

// IDFromName is a convenience function that returns a server's ID given its name.
func IDFromName(client *eclcloud.ServiceClient, name string) (string, error) {
	count := 0
	id := ""

	listOpts := ListOpts{
		// Name: name,
	}

	pages, err := List(client, listOpts).AllPages()
	if err != nil {
		return "", err
	}

	all, err := ExtractVolumes(pages)
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
		return "", eclcloud.ErrResourceNotFound{Name: name, ResourceType: "volume"}
	case 1:
		return id, nil
	default:
		return "", eclcloud.ErrMultipleResourcesFound{Name: name, Count: count, ResourceType: "volume"}
	}
}
