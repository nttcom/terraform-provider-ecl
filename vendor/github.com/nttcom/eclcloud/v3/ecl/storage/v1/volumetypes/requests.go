package volumetypes

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

// Get retrieves the VolumeType with the provided ID.
// To extract the VolumeType object from the response,
// call the Extract method on the GetResult.
func Get(client *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToVolumeTypeListQuery() (string, error)
}

// ListOpts holds options for listing ToVolumeTypes.
// It is passed to the volumetypes.List function.
type ListOpts struct {
	// Now there are no definiton as query params in API specification
	// But do not remove this struct in future specification change.
}

// ToVolumeTypeListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToVolumeTypeListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns VolumeType optionally limited by the conditions provided in ListOpts.
func List(client *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToVolumeTypeListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return VolumeTypePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// IDFromName is a convienience function that returns a server's ID given its name.
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

	all, err := ExtractVolumeTypes(pages)
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
		return "", eclcloud.ErrResourceNotFound{Name: name, ResourceType: "volume_type"}
	case 1:
		return id, nil
	default:
		return "", eclcloud.ErrMultipleResourcesFound{Name: name, Count: count, ResourceType: "virtual_storage"}
	}
}
