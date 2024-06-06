package flavors

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

// Get retrieves the flavor with the provided ID.
// To extract the Flavor object from the response,
// call the Extract method on the GetResult.
func Get(client *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToFlavorListQuery() (string, error)
}

// ListOpts holds options for listing flavors.
// It is passed to the flavors.List function.
type ListOpts struct {
	// Now there are no definition as query params in API specification
	// But do not remove this struct in future specification change.
}

// ToFlavorListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToFlavorListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns Flavor optionally limited by the conditions provided in ListOpts.
func List(client *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToFlavorListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return FlavorPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// IDFromName is a convenience function that returns a flavor's ID given its
// name.
func IDFromName(client *eclcloud.ServiceClient, name string) (string, error) {
	count := 0
	id := ""
	allPages, err := List(client, nil).AllPages()
	if err != nil {
		return "", err
	}

	all, err := ExtractFlavors(allPages)
	if err != nil {
		return "", err
	}

	for _, f := range all {
		if f.Name == name {
			count++
			id = f.ID
		}
	}

	switch count {
	case 0:
		err := &eclcloud.ErrResourceNotFound{}
		err.ResourceType = "flavor"
		err.Name = name
		return "", err
	case 1:
		return id, nil
	default:
		err := &eclcloud.ErrMultipleResourcesFound{}
		err.ResourceType = "flavor"
		err.Name = name
		err.Count = count
		return "", err
	}
}
