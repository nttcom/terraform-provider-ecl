package target_groups

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

/*
List Target Groups
*/

// ListOpts allows the filtering and sorting of paginated collections through the API.
// Filtering is achieved by passing in struct field values that map to the target group attributes you want to see returned.
type ListOpts struct {

	// - ID of the resource
	ID string `q:"id"`

	// - Name of the resource
	// - This field accepts UTF-8 characters up to 3 bytes
	Name string `q:"name"`

	// - Description of the resource
	// - This field accepts UTF-8 characters up to 3 bytes
	Description string `q:"description"`

	// - Configuration status of the resource
	ConfigurationStatus string `q:"configuration_status"`

	// - Operation status of the resource
	OperationStatus string `q:"operation_status"`

	// - ID of the load balancer which the resource belongs to
	LoadBalancerID string `q:"load_balancer_id"`

	// - ID of the owner tenant of the resource
	TenantID string `q:"tenant_id"`
}

// ToTargetGroupListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToTargetGroupListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)

	return q.String(), err
}

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToTargetGroupListQuery() (string, error)
}

// List returns a Pager which allows you to iterate over a collection of target groups.
// It accepts a ListOpts struct, which allows you to filter and sort the returned collection for greater efficiency.
func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)

	if opts != nil {
		query, err := opts.ToTargetGroupListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}

		url += query
	}

	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return TargetGroupPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

/*
Create Target Group
*/

// CreateOptsMember represents member information in the target group creation.
type CreateOptsMember struct {

	// - IP address of the member (real server)
	// - Set a unique combination of IP address and port in all members which belong to the same target group
	// - Must not set a IP address which is included in `virtual_ip_address` and `reserved_fixed_ips` of load balancer interfaces that the target group belongs to
	// - Must not set a IP address of listeners which belong to the same load balancer as the target group
	// - Cannot use a IP address in the following networks
	//   - This host on this network (0.0.0.0/8)
	//   - Shared Address Space (100.64.0.0/10)
	//   - Loopback (127.0.0.0/8)
	//   - Link Local (169.254.0.0/16)
	//   - Multicast (224.0.0.0/4)
	//   - Reserved (240.0.0.0/4)
	//   - Limited Broadcast (255.255.255.255/32)
	IPAddress string `json:"ip_address"`

	// - Port number of the member (real server)
	// - Set a unique combination of IP address and port in all members which belong to the same target group
	Port int `json:"port"`

	// - Weight for the member (real server)
	// - If `policy.algorithm` is `"weighted-round-robin"` or `"weighted-least-connection"`, use this parameter
	// - Set same weight for the combination of IP address and port in all members which belong to the same load balancer
	Weight int `json:"weight,omitempty"`
}

// CreateOpts represents options used to create a new target group.
type CreateOpts struct {

	// - Name of the target group
	// - This field accepts UTF-8 characters up to 3 bytes
	Name string `json:"name,omitempty"`

	// - Description of the target group
	// - This field accepts UTF-8 characters up to 3 bytes
	Description string `json:"description,omitempty"`

	// - Tags of the target group
	// - Set JSON object up to 32,767 characters
	//   - Nested structure is permitted
	//   - The whitespace around separators ( `","` and `":"` ) are ignored
	// - This field accepts UTF-8 characters up to 3 bytes
	Tags map[string]interface{} `json:"tags,omitempty"`

	// - ID of the load balancer which the target group belongs to
	LoadBalancerID string `json:"load_balancer_id"`

	// - Members (real servers) of the target group
	Members *[]CreateOptsMember `json:"members"`
}

// ToTargetGroupCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToTargetGroupCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "target_group")
}

// CreateOptsBuilder allows extensions to add additional parameters to the Create request.
type CreateOptsBuilder interface {
	ToTargetGroupCreateMap() (map[string]interface{}, error)
}

// Create accepts a CreateOpts struct and creates a new target group using the values provided.
func Create(c *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToTargetGroupCreateMap()
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
Show Target Group
*/

// ShowOpts represents options used to show a target group.
type ShowOpts struct {

	// - If `true` is set, `current` and `staged` are returned in response body
	Changes bool `q:"changes"`
}

// ToTargetGroupShowQuery formats a ShowOpts into a query string.
func (opts ShowOpts) ToTargetGroupShowQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)

	return q.String(), err
}

// ShowOptsBuilder allows extensions to add additional parameters to the Show request.
type ShowOptsBuilder interface {
	ToTargetGroupShowQuery() (string, error)
}

// Show retrieves a specific target group based on its unique ID.
func Show(c *eclcloud.ServiceClient, id string, opts ShowOptsBuilder) (r ShowResult) {
	url := showURL(c, id)

	if opts != nil {
		query, _ := opts.ToTargetGroupShowQuery()
		url += query
	}

	_, r.Err = c.Get(url, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Update Target Group Attributes
*/

// UpdateOpts represents options used to update a existing target group.
type UpdateOpts struct {

	// - Name of the target group
	// - This field accepts UTF-8 characters up to 3 bytes
	Name *string `json:"name,omitempty"`

	// - Description of the target group
	// - This field accepts UTF-8 characters up to 3 bytes
	Description *string `json:"description,omitempty"`

	// - Tags of the target group
	// - Set JSON object up to 32,767 characters
	//   - Nested structure is permitted
	//   - The whitespace around separators ( `","` and `":"` ) are ignored
	// - This field accepts UTF-8 characters up to 3 bytes
	Tags *map[string]interface{} `json:"tags,omitempty"`
}

// ToTargetGroupUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToTargetGroupUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "target_group")
}

// UpdateOptsBuilder allows extensions to add additional parameters to the Update request.
type UpdateOptsBuilder interface {
	ToTargetGroupUpdateMap() (map[string]interface{}, error)
}

// Update accepts a UpdateOpts struct and updates a existing target group using the values provided.
func Update(c *eclcloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToTargetGroupUpdateMap()
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
Delete Target Group
*/

// Delete accepts a unique ID and deletes the target group associated with it.
func Delete(c *eclcloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, id), &eclcloud.RequestOpts{
		OkCodes: []int{204},
	})

	return
}

/*
Create Staged Target Group Configurations
*/

// CreateStagedOptsMember represents member information in the target group configurations creation.
type CreateStagedOptsMember struct {

	// - IP address of the member (real server)
	// - Set a unique combination of IP address and port in all members which belong to the same target group
	// - Must not set a IP address which is included in `virtual_ip_address` and `reserved_fixed_ips` of load balancer interfaces that the target group belongs to
	// - Must not set a IP address of listeners which belong to the same load balancer as the target group
	// - Cannot use a IP address in the following networks
	//   - This host on this network (0.0.0.0/8)
	//   - Shared Address Space (100.64.0.0/10)
	//   - Loopback (127.0.0.0/8)
	//   - Link Local (169.254.0.0/16)
	//   - Multicast (224.0.0.0/4)
	//   - Reserved (240.0.0.0/4)
	//   - Limited Broadcast (255.255.255.255/32)
	IPAddress string `json:"ip_address"`

	// - Port number of the member (real server)
	// - Set a unique combination of IP address and port in all members which belong to the same target group
	Port int `json:"port"`

	// - Weight for the member (real server)
	// - If `policy.algorithm` is `"weighted-round-robin"` or `"weighted-least-connection"`, use this parameter
	// - Set same weight for the combination of IP address and port in all members which belong to the same load balancer
	Weight int `json:"weight,omitempty"`
}

// CreateStagedOpts represents options used to create new target group configurations.
type CreateStagedOpts struct {

	// - Members (real servers) of the target group
	Members *[]CreateStagedOptsMember `json:"members,omitempty"`
}

// ToTargetGroupCreateStagedMap builds a request body from CreateStagedOpts.
func (opts CreateStagedOpts) ToTargetGroupCreateStagedMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "target_group")
}

// CreateStagedOptsBuilder allows extensions to add additional parameters to the CreateStaged request.
type CreateStagedOptsBuilder interface {
	ToTargetGroupCreateStagedMap() (map[string]interface{}, error)
}

// CreateStaged accepts a CreateStagedOpts struct and creates new target group configurations using the values provided.
func CreateStaged(c *eclcloud.ServiceClient, id string, opts CreateStagedOptsBuilder) (r CreateStagedResult) {
	b, err := opts.ToTargetGroupCreateStagedMap()
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
Show Staged Target Group Configurations
*/

// ShowStaged retrieves specific target group configurations based on its unique ID.
func ShowStaged(c *eclcloud.ServiceClient, id string) (r ShowStagedResult) {
	_, r.Err = c.Get(showStagedURL(c, id), &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Update Staged Target Group Configurations
*/

// UpdateStagedOptsMember represents member information in target group configurations updation.
type UpdateStagedOptsMember struct {

	// - IP address of the member (real server)
	// - Set a unique combination of IP address and port in all members which belong to the same target group
	// - Must not set a IP address which is included in `virtual_ip_address` and `reserved_fixed_ips` of load balancer interfaces that the target group belongs to
	// - Must not set a IP address of listeners which belong to the same load balancer as the target group
	// - Cannot use a IP address in the following networks
	//   - This host on this network (0.0.0.0/8)
	//   - Shared Address Space (100.64.0.0/10)
	//   - Loopback (127.0.0.0/8)
	//   - Link Local (169.254.0.0/16)
	//   - Multicast (224.0.0.0/4)
	//   - Reserved (240.0.0.0/4)
	//   - Limited Broadcast (255.255.255.255/32)
	IPAddress *string `json:"ip_address"`

	// - Port number of the member (real server)
	// - Set a unique combination of IP address and port in all members which belong to the same target group
	Port *int `json:"port"`

	// - Weight for the member (real server)
	// - If `policy.algorithm` is `"weighted-round-robin"` or `"weighted-least-connection"`, use this parameter
	// - Set same weight for the combination of IP address and port in all members which belong to the same load balancer
	Weight *int `json:"weight,omitempty"`
}

// UpdateStagedOpts represents options used to update existing Target Group configurations.
type UpdateStagedOpts struct {

	// - Members (real servers) of the target group
	Members *[]UpdateStagedOptsMember `json:"members,omitempty"`
}

// ToTargetGroupUpdateStagedMap builds a request body from UpdateStagedOpts.
func (opts UpdateStagedOpts) ToTargetGroupUpdateStagedMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "target_group")
}

// UpdateStagedOptsBuilder allows extensions to add additional parameters to the UpdateStaged request.
type UpdateStagedOptsBuilder interface {
	ToTargetGroupUpdateStagedMap() (map[string]interface{}, error)
}

// UpdateStaged accepts a UpdateStagedOpts struct and updates existing Target Group configurations using the values provided.
func UpdateStaged(c *eclcloud.ServiceClient, id string, opts UpdateStagedOptsBuilder) (r UpdateStagedResult) {
	b, err := opts.ToTargetGroupUpdateStagedMap()
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
Cancel Staged Target Group Configurations
*/

// CancelStaged accepts a unique ID and deletes target group configurations associated with it.
func CancelStaged(c *eclcloud.ServiceClient, id string) (r CancelStagedResult) {
	_, r.Err = c.Delete(cancelStagedURL(c, id), &eclcloud.RequestOpts{
		OkCodes: []int{204},
	})

	return
}
