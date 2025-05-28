package policies

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

/*
List Policies
*/

// ListOpts allows the filtering and sorting of paginated collections through the API.
// Filtering is achieved by passing in struct field values that map to the policy attributes you want to see returned.
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

	// - Load balancing algorithm (method) of the policy
	Algorithm string `q:"algorithm"`

	// - Persistence setting of the policy
	Persistence string `q:"persistence"`

	// - If `persistence` is `"source-ip"`
	//   - The timeout (in minutes) during which the persistence remain after the latest traffic from the client is sent to the load balancer
	// - If `persistence` is `"cookie"`
	//   - The expiration (in minutes) of the persistence set in the cookie that the load balancer returns to the client
	PersistenceTimeout int `q:"persistence_timeout"`

	// - The timeout (in seconds) during which a session is allowed to remain inactive
	IdleTimeout int `q:"idle_timeout"`

	// - URL of the sorry page to which accesses are redirected if all members in the target group are down
	SorryPageUrl string `q:"sorry_page_url"`

	// - Source NAT setting of the policy
	SourceNat string `q:"source_nat"`

	// - ID of the certificate that assigned to the policy
	// - Also includes certificate contained in `server_name_indications`
	CertificateID string `q:"certificate_id"`

	// - ID of the health monitor that assigned to the policy
	HealthMonitorID string `q:"health_monitor_id"`

	// - ID of the listener that assigned to the policy
	ListenerID string `q:"listener_id"`

	// - ID of the default target group that assigned to the policy
	DefaultTargetGroupID string `q:"default_target_group_id"`

	// - ID of the backup target group that assigned to the policy
	// - If all members of the default target group are down, traffic is routed to the backup target group
	BackupTargetGroupID string `q:"backup_target_group_id"`

	// - ID of the TLS policy that assigned to the policy
	TLSPolicyID string `q:"tls_policy_id"`

	// - ID of the load balancer which the resource belongs to
	LoadBalancerID string `q:"load_balancer_id"`

	// - ID of the owner tenant of the resource
	TenantID string `q:"tenant_id"`
}

// ToPolicyListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToPolicyListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)

	return q.String(), err
}

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToPolicyListQuery() (string, error)
}

// List returns a Pager which allows you to iterate over a collection of policies.
// It accepts a ListOpts struct, which allows you to filter and sort the returned collection for greater efficiency.
func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)

	if opts != nil {
		query, err := opts.ToPolicyListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}

		url += query
	}

	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return PolicyPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

/*
Create Policy
*/

// CreateOptsServerNameIndication represents server_name_indication information in the policy creation.
type CreateOptsServerNameIndication struct {

	// - The server name of Server Name Indication (SNI)
	// - Must be unique in a policy
	// - If `input_type` is `"fixed"` , the following restrictions apply
	//   - Only `a-z A-Z 0-9 - . *` are allowed
	//   - `"*"` and `"."` are count as double (2 characters)
	ServerName string `json:"server_name"`

	// - You can choice the input type of the server name
	InputType string `json:"input_type,omitempty"`

	// - Priority of Server Name Indication (SNI)
	// - Must be unique in a policy
	Priority int `json:"priority"`

	// - ID of the certificate that assigned to Server Name Indication (SNI)
	// - The certificate need to be in `"UPLOADED"` state before used in a policy
	// - The load balancer can be configured with up to 50 unique certificates, combining `policy.certificate_id` and `policy.server_name_indications.certificate_id`
	CertificateID string `json:"certificate_id"`
}

// CreateOpts represents options used to create a new policy.
type CreateOpts struct {

	// - Name of the policy
	// - This field accepts UTF-8 characters up to 3 bytes
	Name string `json:"name,omitempty"`

	// - Description of the policy
	// - This field accepts UTF-8 characters up to 3 bytes
	Description string `json:"description,omitempty"`

	// - Tags of the policy
	// - Set JSON object up to 32,767 characters
	//   - Nested structure is permitted
	//   - The whitespace around separators ( `","` and `":"` ) are ignored
	// - This field accepts UTF-8 characters up to 3 bytes
	Tags map[string]interface{} `json:"tags,omitempty"`

	// - Load balancing algorithm (method) of the policy
	Algorithm string `json:"algorithm,omitempty"`

	// - Persistence setting of the policy
	// - If `listener.protocol` is `"http"` or `"https"`, `"cookie"` is available
	Persistence string `json:"persistence,omitempty"`

	// - If `persistence` is `"none"`
	//   - Must not set this parameter or set `0`
	// - If `persistence` is `"source-ip"`
	//   - The timeout (in minutes) during which the persistence remain after the latest traffic from the client is sent to the load balancer
	//   - Default value is `5`
	//   - This parameter can be set between `1` to `2000`
	// - If `persistence` is `"cookie"`
	//   - The expiration (in minutes) of the persistence set in the cookie that the load balancer returns to the client
	//     - If you specify `0` , the cookie persists only for the current session
	//   - Default value is `525600`
	//   - This parameter can be set between `0` to `525600`
	PersistenceTimeout int `json:"persistence_timeout,omitempty"`

	// - The timeout (in seconds) during which a session is allowed to remain inactive
	// - There may be a time difference up to 60 seconds, between the set value and the actual timeout
	// - If `listener.protocol` is `"tcp"` or `"udp"`
	//   - Default value is `120`
	// - If `listener.protocol` is `"http"` or `"https"`
	//   - Default value is `600`
	//   - On session timeout, the load balancer sends TCP RST packets to both the client and the real server
	IdleTimeout int `json:"idle_timeout,omitempty"`

	// - URL of the sorry page to which accesses are redirected if all members in the target groups are down
	// - If `listener.protocol` is `"http"` or `"https"`, this parameter can be set
	// - If `listener.protocol` is neither `"http"` nor `"https"`, must not set this parameter or set `""`
	SorryPageUrl string `json:"sorry_page_url,omitempty"`

	// - Source NAT setting of the policy
	// - If `source_nat` is `"enable"` and `listener.protocol` is `"http"` or `"https"`
	//   - The source IP address of the request is replaced with `virtual_ip_address` which is assigned to the interface from which the request was sent
	//   - `X-Forwarded-For` header with the IP address of the client is added
	SourceNat string `json:"source_nat,omitempty"`

	// - The list of Server Name Indications (SNIs) allows the policy to presents multiple certificates on the same listener
	// - The SNI with the highest priority value will be used when multiple SNIs match
	// - If `listener.protocol` is not `"https"`, must not set this parameter or set `[]`
	//   - If you change `listener.protocol` from `"https"` to others, set `[]`
	ServerNameIndications *[]CreateOptsServerNameIndication `json:"server_name_indications,omitempty"`

	// - ID of the certificate that assigned to the policy
	// - The certificate need to be in `"UPLOADED"` state before used in a policy
	// - If `listener.protocol` is `"https"`, set `certificate.id`
	// - If `listener.protocol` is not `"https"`, must not set this parameter or set `""`
	// - The load balancer can be configured with up to 50 unique certificates, combining `policy.certificate_id` and `policy.server_name_indications.certificate_id`
	CertificateID string `json:"certificate_id,omitempty"`

	// - ID of the health monitor that assigned to the policy
	// - Must not set ID of the health monitor that `configuration_status` is `"DELETE_STAGED"`
	HealthMonitorID string `json:"health_monitor_id"`

	// - ID of the listener that assigned to the policy
	// - Must not set ID of the listener that `configuration_status` is `"DELETE_STAGED"`
	// - Must not set ID of the listener that already assigned to the other policy
	ListenerID string `json:"listener_id"`

	// - ID of the default target group that assigned to the policy
	// - If all members of the default target group are down:
	//   - When `backup_target_group_id` is set, traffic is routed to it
	//   - When `sorry_page_url` is set, accesses are redirected to URL of the sorry page
	//   - When both `backup_target_group_id` and `sorry_page_url` are not set, the load balancer does not respond
	// - The same member cannot be specified for the default target group and the backup target group
	// - Must not set ID of the target group that `configuration_status` is `"DELETE_STAGED"`
	DefaultTargetGroupID string `json:"default_target_group_id"`

	// - ID of the backup target group that assigned to the policy
	// - If all members of the default target group are down, traffic is routed to the backup target group
	// - If all members of the backup target group are down:
	//   - When `sorry_page_url` is set, accesses are redirected to URL of the sorry page
	//   - When `sorry_page_url` is not set, the load balancer does not respond
	// - Set a different ID of the target group from `default_target_group_id`
	// - The same member cannot be specified for the default target group and the backup target group
	// - Must not set ID of the target group that `configuration_status` is `"DELETE_STAGED"`
	BackupTargetGroupID string `json:"backup_target_group_id,omitempty"`

	// - ID of the TLS policy that assigned to the policy
	// - If `listener.protocol` is `"https"`, you can set this parameter explicitly
	//   - If not set this parameter, the ID of the `tls_policy` with `default: true` will be automatically set
	// - If `listener.protocol` is not `"https"`, must not set this parameter or set `""`
	TLSPolicyID string `json:"tls_policy_id,omitempty"`

	// - ID of the load balancer which the policy belongs to
	LoadBalancerID string `json:"load_balancer_id"`
}

// ToPolicyCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToPolicyCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "policy")
}

// CreateOptsBuilder allows extensions to add additional parameters to the Create request.
type CreateOptsBuilder interface {
	ToPolicyCreateMap() (map[string]interface{}, error)
}

// Create accepts a CreateOpts struct and creates a new policy using the values provided.
func Create(c *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToPolicyCreateMap()
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
Show Policy
*/

// ShowOpts represents options used to show a policy.
type ShowOpts struct {

	// - If `true` is set, `current` and `staged` are returned in response body
	Changes bool `q:"changes"`
}

// ToPolicyShowQuery formats a ShowOpts into a query string.
func (opts ShowOpts) ToPolicyShowQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)

	return q.String(), err
}

// ShowOptsBuilder allows extensions to add additional parameters to the Show request.
type ShowOptsBuilder interface {
	ToPolicyShowQuery() (string, error)
}

// Show retrieves a specific policy based on its unique ID.
func Show(c *eclcloud.ServiceClient, id string, opts ShowOptsBuilder) (r ShowResult) {
	url := showURL(c, id)

	if opts != nil {
		query, _ := opts.ToPolicyShowQuery()
		url += query
	}

	_, r.Err = c.Get(url, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Update Policy Attributes
*/

// UpdateOpts represents options used to update a existing policy.
type UpdateOpts struct {

	// - Name of the policy
	// - This field accepts UTF-8 characters up to 3 bytes
	Name *string `json:"name,omitempty"`

	// - Description of the policy
	// - This field accepts UTF-8 characters up to 3 bytes
	Description *string `json:"description,omitempty"`

	// - Tags of the policy
	// - Set JSON object up to 32,767 characters
	//   - Nested structure is permitted
	//   - The whitespace around separators ( `","` and `":"` ) are ignored
	// - This field accepts UTF-8 characters up to 3 bytes
	Tags *map[string]interface{} `json:"tags,omitempty"`
}

// ToPolicyUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToPolicyUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "policy")
}

// UpdateOptsBuilder allows extensions to add additional parameters to the Update request.
type UpdateOptsBuilder interface {
	ToPolicyUpdateMap() (map[string]interface{}, error)
}

// Update accepts a UpdateOpts struct and updates a existing policy using the values provided.
func Update(c *eclcloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToPolicyUpdateMap()
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
Delete Policy
*/

// Delete accepts a unique ID and deletes the policy associated with it.
func Delete(c *eclcloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, id), &eclcloud.RequestOpts{
		OkCodes: []int{204},
	})

	return
}

/*
Create Staged Policy Configurations
*/

// CreateStagedOptsServerNameIndication represents server_name_indication information in the policy configurations creation.
type CreateStagedOptsServerNameIndication struct {

	// - The server name of Server Name Indication (SNI)
	// - Must be unique in a policy
	// - If `input_type` is `"fixed"` , the following restrictions apply
	//   - Only `a-z A-Z 0-9 - . *` are allowed
	//   - `"*"` and `"."` are count as double (2 characters)
	ServerName string `json:"server_name"`

	// - You can choice the input type of the server name
	InputType string `json:"input_type,omitempty"`

	// - Priority of Server Name Indication (SNI)
	// - Must be unique in a policy
	Priority int `json:"priority"`

	// - ID of the certificate that assigned to Server Name Indication (SNI)
	// - The certificate need to be in `"UPLOADED"` state before used in a policy
	// - The load balancer can be configured with up to 50 unique certificates, combining `policy.certificate_id` and `policy.server_name_indications.certificate_id`
	CertificateID string `json:"certificate_id"`
}

// CreateStagedOpts represents options used to create new policy configurations.
type CreateStagedOpts struct {

	// - Load balancing algorithm (method) of the policy
	Algorithm string `json:"algorithm,omitempty"`

	// - Persistence setting of the policy
	// - If `listener.protocol` is `"http"` or `"https"`, `"cookie"` is available
	Persistence string `json:"persistence,omitempty"`

	// - If `persistence` is `"none"`
	//   - Must not set this parameter or set `0`
	// - If `persistence` is `"source-ip"`
	//   - The timeout (in minutes) during which the persistence remain after the latest traffic from the client is sent to the load balancer
	//   - Default value is `5`
	//   - This parameter can be set between `1` to `2000`
	// - If `persistence` is `"cookie"`
	//   - The expiration (in minutes) of the persistence set in the cookie that the load balancer returns to the client
	//     - If you specify `0` , the cookie persists only for the current session
	//   - Default value is `525600`
	//   - This parameter can be set between `0` to `525600`
	PersistenceTimeout int `json:"persistence_timeout,omitempty"`

	// - The timeout (in seconds) during which a session is allowed to remain inactive
	// - There may be a time difference up to 60 seconds, between the set value and the actual timeout
	// - If `listener.protocol` is `"tcp"` or `"udp"`
	//   - Default value is `120`
	// - If `listener.protocol` is `"http"` or `"https"`
	//   - Default value is `600`
	//   - On session timeout, the load balancer sends TCP RST packets to both the client and the real server
	IdleTimeout int `json:"idle_timeout,omitempty"`

	// - URL of the sorry page to which accesses are redirected if all members in the target groups are down
	// - If `listener.protocol` is `"http"` or `"https"`, this parameter can be set
	// - If `listener.protocol` is neither `"http"` nor `"https"`, must not set this parameter or set `""`
	//   - If you change `listener.protocol` from `"http"` or `"https"` to others, set `""`
	SorryPageUrl string `json:"sorry_page_url,omitempty"`

	// - Source NAT setting of the policy
	// - If `source_nat` is `"enable"` and `listener.protocol` is `"http"` or `"https"`
	//   - The source IP address of the request is replaced with `virtual_ip_address` which is assigned to the interface from which the request was sent
	//   - `X-Forwarded-For` header with the IP address of the client is added
	SourceNat string `json:"source_nat,omitempty"`

	// - The list of Server Name Indications (SNIs) allows the policy to presents multiple certificates on the same listener
	// - The SNI with the highest priority value will be used when multiple SNIs match
	// - If `listener.protocol` is not `"https"`, must not set this parameter or set `[]`
	//   - If you change `listener.protocol` from `"https"` to others, set `[]`
	ServerNameIndications *[]CreateStagedOptsServerNameIndication `json:"server_name_indications,omitempty"`

	// - ID of the certificate that assigned to the policy
	// - The certificate need to be in `"UPLOADED"` state before used in a policy
	// - If `listener.protocol` is `"https"`, set `certificate.id`
	// - If `listener.protocol` is not `"https"`, must not set this parameter or set `""`
	//   - If you change `listener.protocol` from `"https"` to others, set `""`
	// - The load balancer can be configured with up to 50 unique certificates, combining `policy.certificate_id` and `policy.server_name_indications.certificate_id`
	CertificateID string `json:"certificate_id,omitempty"`

	// - ID of the health monitor that assigned to the policy
	// - Must not set ID of the health monitor that `configuration_status` is `"DELETE_STAGED"`
	HealthMonitorID string `json:"health_monitor_id,omitempty"`

	// - ID of the listener that assigned to the policy
	// - Must not set ID of the listener that `configuration_status` is `"DELETE_STAGED"`
	// - Must not set ID of the listener that already assigned to the other policy
	ListenerID string `json:"listener_id,omitempty"`

	// - ID of the default target group that assigned to the policy
	// - If all members of the default target group are down:
	//   - When `backup_target_group_id` is set, traffic is routed to it
	//   - When `sorry_page_url` is set, accesses are redirected to URL of the sorry page
	//   - When both `backup_target_group_id` and `sorry_page_url` are not set, the load balancer does not respond
	// - The same member cannot be specified for the default target group and the backup target group
	// - Must not set ID of the target group that `configuration_status` is `"DELETE_STAGED"`
	DefaultTargetGroupID string `json:"default_target_group_id,omitempty"`

	// - ID of the backup target group that assigned to the policy
	// - If all members of the default target group are down, traffic is routed to the backup target group
	// - If all members of the backup target group are down:
	//   - When `sorry_page_url` is set, accesses are redirected to URL of the sorry page
	//   - When `sorry_page_url` is not set, the load balancer does not respond
	// - Set a different ID of the target group from `default_target_group_id`
	// - The same member cannot be specified for the default target group and the backup target group
	// - Must not set ID of the target group that `configuration_status` is `"DELETE_STAGED"`
	BackupTargetGroupID string `json:"backup_target_group_id,omitempty"`

	// - ID of the TLS policy that assigned to the policy
	// - If `listener.protocol` is `"https"`, you can set this parameter explicitly
	//   - If not set this parameter, the ID of the `tls_policy` with `default: true` will be automatically set
	// - If `listener.protocol` is not `"https"`, must not set this parameter or set `""`
	//   - If you change `listener.protocol` from `"https"` to others, set `""`
	TLSPolicyID string `json:"tls_policy_id,omitempty"`
}

// ToPolicyCreateStagedMap builds a request body from CreateStagedOpts.
func (opts CreateStagedOpts) ToPolicyCreateStagedMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "policy")
}

// CreateStagedOptsBuilder allows extensions to add additional parameters to the CreateStaged request.
type CreateStagedOptsBuilder interface {
	ToPolicyCreateStagedMap() (map[string]interface{}, error)
}

// CreateStaged accepts a CreateStagedOpts struct and creates new policy configurations using the values provided.
func CreateStaged(c *eclcloud.ServiceClient, id string, opts CreateStagedOptsBuilder) (r CreateStagedResult) {
	b, err := opts.ToPolicyCreateStagedMap()
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
Show Staged Policy Configurations
*/

// ShowStaged retrieves specific policy configurations based on its unique ID.
func ShowStaged(c *eclcloud.ServiceClient, id string) (r ShowStagedResult) {
	_, r.Err = c.Get(showStagedURL(c, id), &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Update Staged Policy Configurations
*/

// UpdateStagedOptsServerNameIndication represents server_name_indication information in policy configurations updation.
type UpdateStagedOptsServerNameIndication struct {

	// - The server name of Server Name Indication (SNI)
	// - Must be unique in a policy
	// - If `input_type` is `"fixed"` , the following restrictions apply
	//   - Only `a-z A-Z 0-9 - . *` are allowed
	//   - `"*"` and `"."` are count as double (2 characters)
	ServerName *string `json:"server_name"`

	// - You can choice the input type of the server name
	InputType *string `json:"input_type,omitempty"`

	// - Priority of Server Name Indication (SNI)
	// - Must be unique in a policy
	Priority *int `json:"priority"`

	// - ID of the certificate that assigned to Server Name Indication (SNI)
	// - The certificate need to be in `"UPLOADED"` state before used in a policy
	// - The load balancer can be configured with up to 50 unique certificates, combining `policy.certificate_id` and `policy.server_name_indications.certificate_id`
	CertificateID *string `json:"certificate_id"`
}

// UpdateStagedOpts represents options used to update existing Policy configurations.
type UpdateStagedOpts struct {

	// - Load balancing algorithm (method) of the policy
	Algorithm *string `json:"algorithm,omitempty"`

	// - Persistence setting of the policy
	// - If `listener.protocol` is `"http"` or `"https"`, `"cookie"` is available
	Persistence *string `json:"persistence,omitempty"`

	// - If `persistence` is `"none"`
	//   - Must not set this parameter or set `0`
	// - If `persistence` is `"source-ip"`
	//   - The timeout (in minutes) during which the persistence remain after the latest traffic from the client is sent to the load balancer
	//   - Default value is `5`
	//   - This parameter can be set between `1` to `2000`
	// - If `persistence` is `"cookie"`
	//   - The expiration (in minutes) of the persistence set in the cookie that the load balancer returns to the client
	//     - If you specify `0` , the cookie persists only for the current session
	//   - Default value is `525600`
	//   - This parameter can be set between `0` to `525600`
	PersistenceTimeout *int `json:"persistence_timeout,omitempty"`

	// - The timeout (in seconds) during which a session is allowed to remain inactive
	// - There may be a time difference up to 60 seconds, between the set value and the actual timeout
	// - If `listener.protocol` is `"tcp"` or `"udp"`
	//   - Default value is `120`
	// - If `listener.protocol` is `"http"` or `"https"`
	//   - Default value is `600`
	//   - On session timeout, the load balancer sends TCP RST packets to both the client and the real server
	IdleTimeout *int `json:"idle_timeout,omitempty"`

	// - URL of the sorry page to which accesses are redirected if all members in the target groups are down
	// - If `listener.protocol` is `"http"` or `"https"`, this parameter can be set
	// - If `listener.protocol` is neither `"http"` nor `"https"`, must not set this parameter or set `""`
	//   - If you change `listener.protocol` from `"http"` or `"https"` to others, set `""`
	SorryPageUrl *string `json:"sorry_page_url,omitempty"`

	// - Source NAT setting of the policy
	// - If `source_nat` is `"enable"` and `listener.protocol` is `"http"` or `"https"`
	//   - The source IP address of the request is replaced with `virtual_ip_address` which is assigned to the interface from which the request was sent
	//   - `X-Forwarded-For` header with the IP address of the client is added
	SourceNat *string `json:"source_nat,omitempty"`

	// - The list of Server Name Indications (SNIs) allows the policy to presents multiple certificates on the same listener
	// - The SNI with the highest priority value will be used when multiple SNIs match
	// - If `listener.protocol` is not `"https"`, must not set this parameter or set `[]`
	//   - If you change `listener.protocol` from `"https"` to others, set `[]`
	ServerNameIndications *[]UpdateStagedOptsServerNameIndication `json:"server_name_indications,omitempty"`

	// - ID of the certificate that assigned to the policy
	// - The certificate need to be in `"UPLOADED"` state before used in a policy
	// - If `listener.protocol` is `"https"`, set `certificate.id`
	// - If `listener.protocol` is not `"https"`, must not set this parameter or set `""`
	//   - If you change `listener.protocol` from `"https"` to others, set `""`
	// - The load balancer can be configured with up to 50 unique certificates, combining `policy.certificate_id` and `policy.server_name_indications.certificate_id`
	CertificateID *string `json:"certificate_id,omitempty"`

	// - ID of the health monitor that assigned to the policy
	// - Must not set ID of the health monitor that `configuration_status` is `"DELETE_STAGED"`
	HealthMonitorID *string `json:"health_monitor_id,omitempty"`

	// - ID of the listener that assigned to the policy
	// - Must not set ID of the listener that `configuration_status` is `"DELETE_STAGED"`
	// - Must not set ID of the listener that already assigned to the other policy
	ListenerID *string `json:"listener_id,omitempty"`

	// - ID of the default target group that assigned to the policy
	// - If all members of the default target group are down:
	//   - When `backup_target_group_id` is set, traffic is routed to it
	//   - When `sorry_page_url` is set, accesses are redirected to URL of the sorry page
	//   - When both `backup_target_group_id` and `sorry_page_url` are not set, the load balancer does not respond
	// - The same member cannot be specified for the default target group and the backup target group
	// - Must not set ID of the target group that `configuration_status` is `"DELETE_STAGED"`
	DefaultTargetGroupID *string `json:"default_target_group_id,omitempty"`

	// - ID of the backup target group that assigned to the policy
	// - If all members of the default target group are down, traffic is routed to the backup target group
	// - If all members of the backup target group are down:
	//   - When `sorry_page_url` is set, accesses are redirected to URL of the sorry page
	//   - When `sorry_page_url` is not set, the load balancer does not respond
	// - Set a different ID of the target group from `default_target_group_id`
	// - The same member cannot be specified for the default target group and the backup target group
	// - Must not set ID of the target group that `configuration_status` is `"DELETE_STAGED"`
	BackupTargetGroupID *string `json:"backup_target_group_id,omitempty"`

	// - ID of the TLS policy that assigned to the policy
	// - If `listener.protocol` is `"https"`, you can set this parameter explicitly
	//   - If not set this parameter, the ID of the `tls_policy` with `default: true` will be automatically set
	// - If `listener.protocol` is not `"https"`, must not set this parameter or set `""`
	//   - If you change `listener.protocol` from `"https"` to others, set `""`
	TLSPolicyID *string `json:"tls_policy_id,omitempty"`
}

// ToPolicyUpdateStagedMap builds a request body from UpdateStagedOpts.
func (opts UpdateStagedOpts) ToPolicyUpdateStagedMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "policy")
}

// UpdateStagedOptsBuilder allows extensions to add additional parameters to the UpdateStaged request.
type UpdateStagedOptsBuilder interface {
	ToPolicyUpdateStagedMap() (map[string]interface{}, error)
}

// UpdateStaged accepts a UpdateStagedOpts struct and updates existing Policy configurations using the values provided.
func UpdateStaged(c *eclcloud.ServiceClient, id string, opts UpdateStagedOptsBuilder) (r UpdateStagedResult) {
	b, err := opts.ToPolicyUpdateStagedMap()
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
Cancel Staged Policy Configurations
*/

// CancelStaged accepts a unique ID and deletes policy configurations associated with it.
func CancelStaged(c *eclcloud.ServiceClient, id string) (r CancelStagedResult) {
	_, r.Err = c.Delete(cancelStagedURL(c, id), &eclcloud.RequestOpts{
		OkCodes: []int{204},
	})

	return
}
