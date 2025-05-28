package rules

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

/*
List Rules
*/

// ListOpts allows the filtering and sorting of paginated collections through the API.
// Filtering is achieved by passing in struct field values that map to the rule attributes you want to see returned.
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

	// - Priority of the rule
	Priority int `q:"priority"`

	// - ID of the target group that assigned to the rule
	// - If all members of the target groups specified in the rule are down, traffic is routed to the target groups specified in the policy
	TargetGroupID string `q:"target_group_id"`

	// - ID of the backup target group that assigned to the rule
	// - If all members of the target group are down, traffic is routed to the backup target group
	// - If all members of the target groups specified in the rule are down, traffic is routed to the target groups specified in the policy
	BackupTargetGroupID string `q:"backup_target_group_id"`

	// - ID of the policy which the rule belongs to
	PolicyID string `q:"policy_id"`

	// - ID of the load balancer which the resource belongs to
	LoadBalancerID string `q:"load_balancer_id"`

	// - ID of the owner tenant of the resource
	TenantID string `q:"tenant_id"`
}

// ToRuleListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToRuleListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)

	return q.String(), err
}

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToRuleListQuery() (string, error)
}

// List returns a Pager which allows you to iterate over a collection of rules.
// It accepts a ListOpts struct, which allows you to filter and sort the returned collection for greater efficiency.
func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)

	if opts != nil {
		query, err := opts.ToRuleListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}

		url += query
	}

	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return RulePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

/*
Create Rule
*/

// CreateOptsCondition represents condition information in the rule creation.
type CreateOptsCondition struct {

	// - URL path patterns (regular expressions) of the condition
	// - Set a path pattern as unique string in all path patterns which belong to the same policy
	// - Set a path pattern in PCRE (Perl Compatible Regular Expressions) format
	//   - Capturing groups and backreferences are not supported
	PathPatterns []string `json:"path_patterns,omitempty"`
}

// CreateOpts represents options used to create a new rule.
type CreateOpts struct {

	// - Name of the rule
	// - This field accepts UTF-8 characters up to 3 bytes
	Name string `json:"name,omitempty"`

	// - Description of the rule
	// - This field accepts UTF-8 characters up to 3 bytes
	Description string `json:"description,omitempty"`

	// - Tags of the rule
	// - Set JSON object up to 32,767 characters
	//   - Nested structure is permitted
	//   - The whitespace around separators ( `","` and `":"` ) are ignored
	// - This field accepts UTF-8 characters up to 3 bytes
	Tags map[string]interface{} `json:"tags,omitempty"`

	// - Priority of the rule
	// - Set a unique number in all rules which belong to the same policy
	Priority int `json:"priority,omitempty"`

	// - ID of the target group that assigned to the rule
	// - If all members of the target group specified in the rule are down:
	//   - When `backup_target_group_id` of the rule is set, traffic is routed to it
	//   - When `backup_target_group_id` of the rule is not set, traffic is routed to the target groups specified in the policy
	// - The same member cannot be specified for the target group and the backup target group
	// - Must not set ID of the target group that `configuration_status` is `"DELETE_STAGED"`
	TargetGroupID string `json:"target_group_id"`

	// - ID of the backup target group that assigned to the rule
	// - If all members of the target group specified in the rule are down, traffic is routed to the backup target group specified in the rule
	// - If all members of the backup target group specified in the rule are down, traffic is routed to the target groups specified in the policy
	// - Set a different ID of the target group from `target_group_id`
	// - The same member cannot be specified for the target group and the backup target group
	// - Must not set ID of the target group that `configuration_status` is `"DELETE_STAGED"`
	BackupTargetGroupID string `json:"backup_target_group_id,omitempty"`

	// - ID of the policy which the rule belongs to
	// - Set ID of the policy which has a listener in which protocol is either `"http"` or `"https"`
	PolicyID string `json:"policy_id"`

	// - Conditions of the rules to distribute accesses to the target groups
	// - Set one or more condition
	Conditions *CreateOptsCondition `json:"conditions"`
}

// ToRuleCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToRuleCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "rule")
}

// CreateOptsBuilder allows extensions to add additional parameters to the Create request.
type CreateOptsBuilder interface {
	ToRuleCreateMap() (map[string]interface{}, error)
}

// Create accepts a CreateOpts struct and creates a new rule using the values provided.
func Create(c *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToRuleCreateMap()
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
Show Rule
*/

// ShowOpts represents options used to show a rule.
type ShowOpts struct {

	// - If `true` is set, `current` and `staged` are returned in response body
	Changes bool `q:"changes"`
}

// ToRuleShowQuery formats a ShowOpts into a query string.
func (opts ShowOpts) ToRuleShowQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)

	return q.String(), err
}

// ShowOptsBuilder allows extensions to add additional parameters to the Show request.
type ShowOptsBuilder interface {
	ToRuleShowQuery() (string, error)
}

// Show retrieves a specific rule based on its unique ID.
func Show(c *eclcloud.ServiceClient, id string, opts ShowOptsBuilder) (r ShowResult) {
	url := showURL(c, id)

	if opts != nil {
		query, _ := opts.ToRuleShowQuery()
		url += query
	}

	_, r.Err = c.Get(url, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Update Rule Attributes
*/

// UpdateOpts represents options used to update a existing rule.
type UpdateOpts struct {

	// - Name of the rule
	// - This field accepts UTF-8 characters up to 3 bytes
	Name *string `json:"name,omitempty"`

	// - Description of the rule
	// - This field accepts UTF-8 characters up to 3 bytes
	Description *string `json:"description,omitempty"`

	// - Tags of the rule
	// - Set JSON object up to 32,767 characters
	//   - Nested structure is permitted
	//   - The whitespace around separators ( `","` and `":"` ) are ignored
	// - This field accepts UTF-8 characters up to 3 bytes
	Tags *map[string]interface{} `json:"tags,omitempty"`
}

// ToRuleUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToRuleUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "rule")
}

// UpdateOptsBuilder allows extensions to add additional parameters to the Update request.
type UpdateOptsBuilder interface {
	ToRuleUpdateMap() (map[string]interface{}, error)
}

// Update accepts a UpdateOpts struct and updates a existing rule using the values provided.
func Update(c *eclcloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToRuleUpdateMap()
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
Delete Rule
*/

// Delete accepts a unique ID and deletes the rule associated with it.
func Delete(c *eclcloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, id), &eclcloud.RequestOpts{
		OkCodes: []int{204},
	})

	return
}

/*
Create Staged Rule Configurations
*/

// CreateStagedOptsCondition represents condition information in the rule configurations creation.
type CreateStagedOptsCondition struct {

	// - URL path patterns (regular expressions) of the condition
	// - Set a path pattern as unique string in all path patterns which belong to the same policy
	// - Set a path pattern in PCRE (Perl Compatible Regular Expressions) format
	//   - Capturing groups and backreferences are not supported
	PathPatterns []string `json:"path_patterns,omitempty"`
}

// CreateStagedOpts represents options used to create new rule configurations.
type CreateStagedOpts struct {

	// - Priority of the rule
	// - Set a unique number in all rules which belong to the same policy
	Priority int `json:"priority,omitempty"`

	// - ID of the target group that assigned to the rule
	// - If all members of the target group specified in the rule are down:
	//   - When `backup_target_group_id` of the rule is set, traffic is routed to it
	//   - When `backup_target_group_id` of the rule is not set, traffic is routed to the target groups specified in the policy
	// - The same member cannot be specified for the target group and the backup target group
	// - Must not set ID of the target group that `configuration_status` is `"DELETE_STAGED"`
	TargetGroupID string `json:"target_group_id,omitempty"`

	// - ID of the backup target group that assigned to the rule
	// - If all members of the target group specified in the rule are down, traffic is routed to the backup target group specified in the rule
	// - If all members of the backup target group specified in the rule are down, traffic is routed to the target groups specified in the policy
	// - Set a different ID of the target group from `target_group_id`
	// - The same member cannot be specified for the target group and the backup target group
	// - Must not set ID of the target group that `configuration_status` is `"DELETE_STAGED"`
	BackupTargetGroupID string `json:"backup_target_group_id,omitempty"`

	// - Conditions of the rules to distribute accesses to the target groups
	// - Set one or more condition
	Conditions *CreateStagedOptsCondition `json:"conditions,omitempty"`
}

// ToRuleCreateStagedMap builds a request body from CreateStagedOpts.
func (opts CreateStagedOpts) ToRuleCreateStagedMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "rule")
}

// CreateStagedOptsBuilder allows extensions to add additional parameters to the CreateStaged request.
type CreateStagedOptsBuilder interface {
	ToRuleCreateStagedMap() (map[string]interface{}, error)
}

// CreateStaged accepts a CreateStagedOpts struct and creates new rule configurations using the values provided.
func CreateStaged(c *eclcloud.ServiceClient, id string, opts CreateStagedOptsBuilder) (r CreateStagedResult) {
	b, err := opts.ToRuleCreateStagedMap()
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
Show Staged Rule Configurations
*/

// ShowStaged retrieves specific rule configurations based on its unique ID.
func ShowStaged(c *eclcloud.ServiceClient, id string) (r ShowStagedResult) {
	_, r.Err = c.Get(showStagedURL(c, id), &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Update Staged Rule Configurations
*/

// UpdateStagedOptsCondition represents condition information in rule configurations updation.
type UpdateStagedOptsCondition struct {

	// - URL path patterns (regular expressions) of the condition
	// - Set a path pattern as unique string in all path patterns which belong to the same policy
	// - Set a path pattern in PCRE (Perl Compatible Regular Expressions) format
	//   - Capturing groups and backreferences are not supported
	PathPatterns *[]string `json:"path_patterns,omitempty"`
}

// UpdateStagedOpts represents options used to update existing Rule configurations.
type UpdateStagedOpts struct {

	// - Priority of the rule
	// - Set a unique number in all rules which belong to the same policy
	Priority *int `json:"priority,omitempty"`

	// - ID of the target group that assigned to the rule
	// - If all members of the target group specified in the rule are down:
	//   - When `backup_target_group_id` of the rule is set, traffic is routed to it
	//   - When `backup_target_group_id` of the rule is not set, traffic is routed to the target groups specified in the policy
	// - The same member cannot be specified for the target group and the backup target group
	// - Must not set ID of the target group that `configuration_status` is `"DELETE_STAGED"`
	TargetGroupID *string `json:"target_group_id,omitempty"`

	// - ID of the backup target group that assigned to the rule
	// - If all members of the target group specified in the rule are down, traffic is routed to the backup target group specified in the rule
	// - If all members of the backup target group specified in the rule are down, traffic is routed to the target groups specified in the policy
	// - Set a different ID of the target group from `target_group_id`
	// - The same member cannot be specified for the target group and the backup target group
	// - Must not set ID of the target group that `configuration_status` is `"DELETE_STAGED"`
	BackupTargetGroupID *string `json:"backup_target_group_id,omitempty"`

	// - Conditions of the rules to distribute accesses to the target groups
	// - Set one or more condition
	Conditions *UpdateStagedOptsCondition `json:"conditions,omitempty"`
}

// ToRuleUpdateStagedMap builds a request body from UpdateStagedOpts.
func (opts UpdateStagedOpts) ToRuleUpdateStagedMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "rule")
}

// UpdateStagedOptsBuilder allows extensions to add additional parameters to the UpdateStaged request.
type UpdateStagedOptsBuilder interface {
	ToRuleUpdateStagedMap() (map[string]interface{}, error)
}

// UpdateStaged accepts a UpdateStagedOpts struct and updates existing Rule configurations using the values provided.
func UpdateStaged(c *eclcloud.ServiceClient, id string, opts UpdateStagedOptsBuilder) (r UpdateStagedResult) {
	b, err := opts.ToRuleUpdateStagedMap()
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
Cancel Staged Rule Configurations
*/

// CancelStaged accepts a unique ID and deletes rule configurations associated with it.
func CancelStaged(c *eclcloud.ServiceClient, id string) (r CancelStagedResult) {
	_, r.Err = c.Delete(cancelStagedURL(c, id), &eclcloud.RequestOpts{
		OkCodes: []int{204},
	})

	return
}
