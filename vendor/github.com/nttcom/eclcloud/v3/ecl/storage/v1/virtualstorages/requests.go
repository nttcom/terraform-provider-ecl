package virtualstorages

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToVirtualStorageCreateMap() (map[string]interface{}, error)
}

// CreateOpts contains options for creating a VirtualStorage. This object is passed to
// the virtualstorages.Create function. For more information about these parameters,
// see the VirtualStorage object.
type CreateOpts struct {
	// The virtual storage name
	Name string `json:"name" required:"true"`
	// The virtual storage description
	Description string `json:"description,omitempty"`
	// The network_id to connect virtual storage
	NetworkID string `json:"network_id" required:"true"`
	// The subnet_id to connect virtual storage
	SubnetID string `json:"subnet_id" required:"true"`
	// The virtual storage volume_type_id
	VolumeTypeID string `json:"volume_type_id" required:"true"`
	// The ip address pool of virtual storage
	IPAddrPool IPAddressPool `json:"ip_addr_pool" required:"true"`
	// The virtual storage host_routes
	HostRoutes []HostRoute `json:"host_routes,omitempty"`
}

// ToVirtualStorageCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToVirtualStorageCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "virtual_storage")
}

// Create will create a new VirtualStorage based on the values in CreateOpts.
// To extract the VirtualStorage object from the response, call the Extract method on the
// CreateResult.
func Create(client *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToVirtualStorageCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{202},
	})
	return
}

// Delete will delete the existing VirtualStorage with the provided ID.
func Delete(client *eclcloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(
		deleteURL(client, id),
		&eclcloud.RequestOpts{
			OkCodes: []int{200},
		})
	return
}

// Get retrieves the VirtualStorage with the provided ID.
// To extract the VirtualStorage object from the response,
// call the Extract method on the GetResult.
func Get(client *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToVirtualStorageListQuery() (string, error)
}

// ListOpts holds options for listing VirtualStorages.
// It is passed to the virtualstorages.List function.
type ListOpts struct {
	// Now there are no definiton as query params in API specification
	// But do not remove this struct in future specification change.
}

// ToVirtualStorageListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToVirtualStorageListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns VirtualStorage optionally limited by the conditions provided in ListOpts.
func List(client *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToVirtualStorageListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return VirtualStoragePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToVirtualStorageUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts contain options for updating an existing VirtualStorage.
// This object is passed to the virtual_storage.Update function.
// For more information about the parameters, see the VirtualStorage object.
type UpdateOpts struct {
	Name        *string        `json:"name,omitempty"`
	Description *string        `json:"description,omitempty"`
	IPAddrPool  *IPAddressPool `json:"ip_addr_pool,omitempty"`
	HostRoutes  *[]HostRoute   `json:"host_routes,omitempty"`
}

// ToVirtualStorageUpdateMap assembles a request body based on the contents of an
// UpdateOpts.
func (opts UpdateOpts) ToVirtualStorageUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "virtual_storage")
}

// Update will update the VirtualStorage with provided information.
// To extract the updated VirtualStorage from the response,
// call the Extract method on the UpdateResult.
func Update(client *eclcloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToVirtualStorageUpdateMap()
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

	all, err := ExtractVirtualStorages(pages)
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
		return "", eclcloud.ErrResourceNotFound{Name: name, ResourceType: "virtual_storage"}
	case 1:
		return id, nil
	default:
		return "", eclcloud.ErrMultipleResourcesFound{Name: name, Count: count, ResourceType: "virtual_storage"}
	}
}
