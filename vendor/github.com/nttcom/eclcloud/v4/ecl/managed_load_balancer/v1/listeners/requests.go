package listeners

import (
	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/pagination"
)

/*
List Listeners
*/

// ListOpts allows the filtering and sorting of paginated collections through the API.
// Filtering is achieved by passing in struct field values that map to the listener attributes you want to see returned.
type ListOpts struct {

	// - ID of the resource
	ID string `q:"id"`

	// - Name of the resource
	// - This field accepts single-byte characters only
	Name string `q:"name"`

	// - Description of the resource
	// - This field accepts single-byte characters only
	Description string `q:"description"`

	// - Configuration status of the resource
	ConfigurationStatus string `q:"configuration_status"`

	// - Operation status of the resource
	OperationStatus string `q:"operation_status"`

	// - IP address of the resource for listening
	IPAddress string `q:"ip_address"`

	// - Port number of the resource for healthchecking or listening
	Port int `q:"port"`

	// - Protocol of the resource for healthchecking or listening
	Protocol string `q:"protocol"`

	// - ID of the load balancer which the resource belongs to
	LoadBalancerID string `q:"load_balancer_id"`

	// - ID of the owner tenant of the resource
	TenantID string `q:"tenant_id"`
}

// ToListenerListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToListenerListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)

	return q.String(), err
}

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToListenerListQuery() (string, error)
}

// List returns a Pager which allows you to iterate over a collection of listeners.
// It accepts a ListOpts struct, which allows you to filter and sort the returned collection for greater efficiency.
func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)

	if opts != nil {
		query, err := opts.ToListenerListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}

		url += query
	}

	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return ListenerPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

/*
Create Listener
*/

// CreateOpts represents options used to create a new listener.
type CreateOpts struct {

	// - Name of the listener
	// - This field accepts single-byte characters only
	Name string `json:"name,omitempty"`

	// - Description of the listener
	// - This field accepts single-byte characters only
	Description string `json:"description,omitempty"`

	// - Tags of the listener
	// - Set JSON object up to 32,768 characters
	//   - Nested structure is permitted
	// - This field accepts single-byte characters only
	Tags map[string]interface{} `json:"tags,omitempty"`

	// - IP address of the listener for listening
	// - Set an unique combination of IP address and port in all listeners which belong to the same load balancer
	// - Must not set a IP address which is included in `virtual_ip_address` and `reserved_fixed_ips` of load balancer interfaces that the listener belongs to
	// - Must not set a link-local IP address (RFC 3927) which includes Common Function Gateway
	IPAddress string `json:"ip_address"`

	// - Port number of the listener for listening
	// - Combination of IP address and port must be unique for all listeners which belong to the same load balancer
	Port int `json:"port"`

	// - Protocol of the listener for listening
	Protocol string `json:"protocol"`

	// - ID of the load balancer which the listener belongs to
	LoadBalancerID string `json:"load_balancer_id"`
}

// ToListenerCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToListenerCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "listener")
}

// CreateOptsBuilder allows extensions to add additional parameters to the Create request.
type CreateOptsBuilder interface {
	ToListenerCreateMap() (map[string]interface{}, error)
}

// Create accepts a CreateOpts struct and creates a new listener using the values provided.
func Create(c *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToListenerCreateMap()
	if err != nil {
		r.Err = err

		return
	}

	_, r.Err = c.Post(createURL(c), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Show Listener
*/

// ShowOpts represents options used to show a listener.
type ShowOpts struct {

	// - If `true` is set, `current` and `staged` are returned in response body
	Changes bool `q:"changes"`
}

// ToListenerShowQuery formats a ShowOpts into a query string.
func (opts ShowOpts) ToListenerShowQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)

	return q.String(), err
}

// ShowOptsBuilder allows extensions to add additional parameters to the Show request.
type ShowOptsBuilder interface {
	ToListenerShowQuery() (string, error)
}

// Show retrieves a specific listener based on its unique ID.
func Show(c *eclcloud.ServiceClient, id string, opts ShowOptsBuilder) (r ShowResult) {
	url := showURL(c, id)

	if opts != nil {
		query, _ := opts.ToListenerShowQuery()
		url += query
	}

	_, r.Err = c.Get(url, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Update Listener Attribute
*/

// UpdateOpts represents options used to update a existing listener.
type UpdateOpts struct {

	// - Name of the listener
	// - This field accepts single-byte characters only
	Name *string `json:"name,omitempty"`

	// - Description of the listener
	// - This field accepts single-byte characters only
	Description *string `json:"description,omitempty"`

	// - Tags of the listener
	// - Set JSON object up to 32,768 characters
	//   - Nested structure is permitted
	// - This field accepts single-byte characters only
	Tags *map[string]interface{} `json:"tags,omitempty"`
}

// ToListenerUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToListenerUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "listener")
}

// UpdateOptsBuilder allows extensions to add additional parameters to the Update request.
type UpdateOptsBuilder interface {
	ToListenerUpdateMap() (map[string]interface{}, error)
}

// Update accepts a UpdateOpts struct and updates a existing listener using the values provided.
func Update(c *eclcloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToListenerUpdateMap()
	if err != nil {
		r.Err = err

		return
	}

	_, r.Err = c.Patch(updateURL(c, id), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Delete Listener
*/

// Delete accepts a unique ID and deletes the listener associated with it.
func Delete(c *eclcloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, id), &eclcloud.RequestOpts{
		OkCodes: []int{204},
	})

	return
}

/*
Create Staged Listener Configurations
*/

// CreateStagedOpts represents options used to create new listener configurations.
type CreateStagedOpts struct {

	// - IP address of the listener for listening
	// - Set an unique combination of IP address and port in all listeners which belong to the same load balancer
	// - Must not set a IP address which is included in `virtual_ip_address` and `reserved_fixed_ips` of load balancer interfaces that the listener belongs to
	// - Must not set a link-local IP address (RFC 3927) which includes Common Function Gateway
	IPAddress string `json:"ip_address,omitempty"`

	// - Port number of the listener for listening
	// - Combination of IP address and port must be unique for all listeners which belong to the same load balancer
	Port int `json:"port,omitempty"`

	// - Protocol of the listener for listening
	Protocol string `json:"protocol,omitempty"`
}

// ToListenerCreateStagedMap builds a request body from CreateStagedOpts.
func (opts CreateStagedOpts) ToListenerCreateStagedMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "listener")
}

// CreateStagedOptsBuilder allows extensions to add additional parameters to the CreateStaged request.
type CreateStagedOptsBuilder interface {
	ToListenerCreateStagedMap() (map[string]interface{}, error)
}

// CreateStaged accepts a CreateStagedOpts struct and creates new listener configurations using the values provided.
func CreateStaged(c *eclcloud.ServiceClient, id string, opts CreateStagedOptsBuilder) (r CreateStagedResult) {
	b, err := opts.ToListenerCreateStagedMap()
	if err != nil {
		r.Err = err

		return
	}

	_, r.Err = c.Post(createStagedURL(c, id), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Show Staged Listener Configurations
*/

// ShowStaged retrieves specific listener configurations based on its unique ID.
func ShowStaged(c *eclcloud.ServiceClient, id string) (r ShowStagedResult) {
	_, r.Err = c.Get(showStagedURL(c, id), &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Update Staged Listener Configurations
*/

// UpdateStagedOpts represents options used to update existing Listener configurations.
type UpdateStagedOpts struct {

	// - IP address of the listener for listening
	// - Set an unique combination of IP address and port in all listeners which belong to the same load balancer
	// - Must not set a IP address which is included in `virtual_ip_address` and `reserved_fixed_ips` of load balancer interfaces that the listener belongs to
	// - Must not set a link-local IP address (RFC 3927) which includes Common Function Gateway
	IPAddress *string `json:"ip_address,omitempty"`

	// - Port number of the listener for listening
	// - Combination of IP address and port must be unique for all listeners which belong to the same load balancer
	Port *int `json:"port,omitempty"`

	// - Protocol of the listener for listening
	Protocol *string `json:"protocol,omitempty"`
}

// ToListenerUpdateStagedMap builds a request body from UpdateStagedOpts.
func (opts UpdateStagedOpts) ToListenerUpdateStagedMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "listener")
}

// UpdateStagedOptsBuilder allows extensions to add additional parameters to the UpdateStaged request.
type UpdateStagedOptsBuilder interface {
	ToListenerUpdateStagedMap() (map[string]interface{}, error)
}

// UpdateStaged accepts a UpdateStagedOpts struct and updates existing Listener configurations using the values provided.
func UpdateStaged(c *eclcloud.ServiceClient, id string, opts UpdateStagedOptsBuilder) (r UpdateStagedResult) {
	b, err := opts.ToListenerUpdateStagedMap()
	if err != nil {
		r.Err = err

		return
	}

	_, r.Err = c.Patch(updateStagedURL(c, id), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Cancel Staged Listener Configurations
*/

// CancelStaged accepts a unique ID and deletes listener configurations associated with it.
func CancelStaged(c *eclcloud.ServiceClient, id string) (r CancelStagedResult) {
	_, r.Err = c.Delete(cancelStagedURL(c, id), &eclcloud.RequestOpts{
		OkCodes: []int{204},
	})

	return
}
