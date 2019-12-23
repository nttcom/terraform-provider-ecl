package zones

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

// ListOptsBuilder allows extensions to add parameters to the List request.
type ListOptsBuilder interface {
	ToZoneListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the server attributes you want to see returned. Marker and Limit are used
// for pagination.
type ListOpts struct {
	// Domain name of zone for partial-match search.
	DomainName string `q:"domain_name"`
	// Sorts the response by the attribute value. A valid value is only domain_name.
	SortKey string `q:"sort_key"`
	// Sorts the response by the requested sort direction.
	// A valid value is asc (ascending) or desc (descending). Default is asc.
	SortDir string `q:"sort_dir"`
	// UUID of the zone at which you want to set a marker.
	Marker string `q:"marker"`
	// Integer value for the limit of values to return.
	Limit int `q:"limit"`

	// Following are original designate parameters.
	// But can not be used in ECL2.0
	// TODO: Remove them at last of development.
	//
	// Description string `q:"description"`
	// Email       string `q:"email"`
	// Name        string `q:"name"`
	// Status      string `q:"status"`
	// TTL         int    `q:"ttl"`
	// Type        string `q:"type"`
}

// ToZoneListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToZoneListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List implements a zone List request.
func List(client *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := baseURL(client)
	if opts != nil {
		query, err := opts.ToZoneListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ZonePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get returns information about a zone, given its ID.
func Get(client *eclcloud.ServiceClient, zoneID string) (r GetResult) {
	_, r.Err = client.Get(zoneURL(client, zoneID), &r.Body, nil)
	return
}

// CreateOptsBuilder allows extensions to add additional attributes to the
// Create request.
type CreateOptsBuilder interface {
	ToZoneCreateMap() (map[string]interface{}, error)
}

// CreateOpts specifies the attributes used to create a zone.
type CreateOpts struct {
	// Description of the zone.
	Description string `json:"description,omitempty"`

	// Email contact of the zone.
	Email string `json:"email,omitempty"`

	// Name of the zone.
	Name string `json:"name" required:"true"`

	// Masters specifies zone masters if this is a secondary zone.
	Masters []string `json:"masters,omitempty"`

	// TTL is the time to live of the zone.
	TTL int `json:"-"`

	// Type specifies if this is a primary or secondary zone.
	Type string `json:"type,omitempty"`
}

// ToZoneCreateMap formats an CreateOpts structure into a request body.
func (opts CreateOpts) ToZoneCreateMap() (map[string]interface{}, error) {
	b, err := eclcloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	if opts.TTL > 0 {
		b["ttl"] = opts.TTL
	}

	return b, nil
}

// Create implements a zone create request.
func Create(client *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToZoneCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(baseURL(client), &b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{201, 202},
	})
	return
}

// UpdateOptsBuilder allows extensions to add additional attributes to the
// Update request.
type UpdateOptsBuilder interface {
	ToZoneUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts specifies the attributes to update a zone.
type UpdateOpts struct {
	// Description of the zone.
	Description *string `json:"description,omitempty"`

	// TTL is the time to live of the zone.
	TTL *int `json:"ttl,omitempty"`

	// Masters specifies zone masters if this is a secondary zone.
	Masters *[]string `json:"masters,omitempty"`

	// Email contact of the zone.
	Email *string `json:"email,omitempty"`
}

// ToZoneUpdateMap formats an UpdateOpts structure into a request body.
func (opts UpdateOpts) ToZoneUpdateMap() (map[string]interface{}, error) {
	b, err := eclcloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Update implements a zone update request.
func Update(client *eclcloud.ServiceClient, zoneID string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToZoneUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Patch(zoneURL(client, zoneID), &b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200, 202},
	})
	return
}

// Delete implements a zone delete request.
func Delete(client *eclcloud.ServiceClient, zoneID string) (r DeleteResult) {
	_, r.Err = client.Delete(zoneURL(client, zoneID), &eclcloud.RequestOpts{
		OkCodes: []int{202},
	})
	return
}
