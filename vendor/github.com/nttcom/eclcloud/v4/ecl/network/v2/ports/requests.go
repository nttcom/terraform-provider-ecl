package ports

import (
	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToPortListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the port attributes you want to see returned. SortKey allows you to sort
// by a particular port attribute. SortDir sets the direction, and is either
// `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	Description      string `q:"description"`
	DeviceID         string `q:"device_id"`
	DeviceOwner      string `q:"device_owner"`
	ID               string `q:"id"`
	MACAddress       string `q:"mac_address"`
	Name             string `q:"name"`
	NetworkID        string `q:"network_id"`
	SegmentationID   int    `q:"segmentation_id"`
	SegmentationType string `q:"segmentation_type"`
	Status           string `q:"status"`
	TenantID         string `q:"tenant_id"`
}

// ToPortListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToPortListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of
// ports. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those ports that are owned by the tenant
// who submits the request, unless the request is submitted by a user with
// administrative rights.
func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToPortListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return PortPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves a specific port based on its unique ID.
func Get(c *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(getURL(c, id), &r.Body, nil)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToPortCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents the attributes used when creating a new port.
type CreateOpts struct {
	AdminStateUp        *bool             `json:"admin_state_up,omitempty"`
	AllowedAddressPairs []AddressPair     `json:"allowed_address_pairs,omitempty"`
	Description         string            `json:"description,omitempty"`
	DeviceID            string            `json:"device_id,omitempty"`
	DeviceOwner         string            `json:"device_owner,omitempty"`
	FixedIPs            interface{}       `json:"fixed_ips,omitempty"`
	MACAddress          string            `json:"mac_address,omitempty"`
	Name                string            `json:"name,omitempty"`
	NetworkID           string            `json:"network_id"`
	SegmentationID      int               `json:"segmentation_id,omitempty"`
	SegmentationType    string            `json:"segmentation_type,omitempty"`
	Tags                map[string]string `json:"tags,omitempty"`
	TenantID            string            `json:"tenant_id,omitempty"`
}

// ToPortCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToPortCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "port")
}

// Create accepts a CreateOpts struct and creates a new network using the values
// provided. You must remember to provide a NetworkID value.
func Create(c *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToPortCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(createURL(c), b, &r.Body, nil)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToPortUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents the attributes used when updating an existing port.
type UpdateOpts struct {
	AdminStateUp        *bool              `json:"admin_state_up,omitempty"`
	AllowedAddressPairs *[]AddressPair     `json:"allowed_address_pairs,omitempty"`
	Description         *string            `json:"description,omitempty"`
	DeviceID            *string            `json:"device_id,omitempty"`
	DeviceOwner         *string            `json:"device_owner,omitempty"`
	FixedIPs            interface{}        `json:"fixed_ips,omitempty"`
	Name                *string            `json:"name,omitempty"`
	SegmentationID      *int               `json:"segmentation_id,omitempty"`
	SegmentationType    *string            `json:"segmentation_type,omitempty"`
	Tags                *map[string]string `json:"tags,omitempty"`
}

// ToPortUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToPortUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "port")
}

// Update accepts a UpdateOpts struct and updates an existing port using the
// values provided.
func Update(c *eclcloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToPortUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Put(updateURL(c, id), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200, 201},
	})
	return
}

// Delete accepts a unique ID and deletes the port associated with it.
func Delete(c *eclcloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, id), nil)
	return
}

// IDFromName is a convenience function that returns a port's ID,
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

	all, err := ExtractPorts(pages)
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
		return "", eclcloud.ErrResourceNotFound{Name: name, ResourceType: "port"}
	case 1:
		return id, nil
	default:
		return "", eclcloud.ErrMultipleResourcesFound{Name: name, Count: count, ResourceType: "port"}
	}
}
