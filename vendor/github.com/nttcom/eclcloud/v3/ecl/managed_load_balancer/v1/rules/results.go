package rules

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// CreateResult represents the result of a Create operation.
// Call its Extract method to interpret it as a Rule.
type CreateResult struct {
	commonResult
}

// ShowResult represents the result of a Show operation.
// Call its Extract method to interpret it as a Rule.
type ShowResult struct {
	commonResult
}

// UpdateResult represents the result of a Update operation.
// Call its Extract method to interpret it as a Rule.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a Delete operation.
// Call its ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// CreateStagedResult represents the result of a CreateStaged operation.
// Call its Extract method to interpret it as a Rule.
type CreateStagedResult struct {
	commonResult
}

// ShowStagedResult represents the result of a ShowStaged operation.
// Call its Extract method to interpret it as a Rule.
type ShowStagedResult struct {
	commonResult
}

// UpdateStagedResult represents the result of a UpdateStaged operation.
// Call its Extract method to interpret it as a Rule.
type UpdateStagedResult struct {
	commonResult
}

// CancelStagedResult represents the result of a CancelStaged operation.
// Call its ExtractErr method to determine if the request succeeded or failed.
type CancelStagedResult struct {
	eclcloud.ErrResult
}

// ConfigurationInResponse represents a configuration in a rule.
type ConfigurationInResponse struct {

	// - Priority of the rule
	Priority int `json:"priority,omitempty"`

	// - ID of the target group that assigned to the rule
	TargetGroupID string `json:"target_group_id,omitempty"`

	// - Conditions of the rules to distribute accesses to the target groups
	Conditions ConditionInResponse `json:"conditions,omitempty"`
}

// ConditionInResponse represents a condition in a rule.
type ConditionInResponse struct {

	// - URL path patterns (regular expressions) of the condition
	PathPatterns []string `json:"path_patterns"`
}

// Rule represents a rule.
type Rule struct {

	// - ID of the rule
	ID string `json:"id"`

	// - Name of the rule
	Name string `json:"name"`

	// - Description of the rule
	Description string `json:"description"`

	// - Tags of the rule (JSON object format)
	Tags map[string]interface{} `json:"tags"`

	// - Configuration status of the rule
	//   - `"ACTIVE"`
	//     - There are no configurations of the rule that waiting to be applied
	//   - `"CREATE_STAGED"`
	//     - The rule has been added and waiting to be applied
	//   - `"UPDATE_STAGED"`
	//     - Changed configurations of the rule exists that waiting to be applied
	//   - `"DELETE_STAGED"`
	//     - The rule has been removed and waiting to be applied
	ConfigurationStatus string `json:"configuration_status"`

	// - Operation status of the load balancer which the rule belongs to
	//   - `"NONE"` :
	//     - There are no operations of the load balancer
	//     - The load balancer and related resources can be operated
	//   - `"PROCESSING"`
	//     - The latest operation of the load balancer is processing
	//     - The load balancer and related resources cannot be operated
	//   - `"COMPLETE"`
	//     - The latest operation of the load balancer has been succeeded
	//     - The load balancer and related resources can be operated
	//   - `"STUCK"`
	//     - The latest operation of the load balancer has been stopped
	//     - Operators of NTT Communications will investigate the operation
	//     - The load balancer and related resources cannot be operated
	//   - `"ERROR"`
	//     - The latest operation of the load balancer has been failed
	//     - The operation was roll backed normally
	//     - The load balancer and related resources can be operated
	OperationStatus string `json:"operation_status"`

	// - ID of the policy which the rule belongs to
	PolicyID string `json:"policy_id"`

	// - ID of the load balancer which the rule belongs to
	LoadBalancerID string `json:"load_balancer_id"`

	// - ID of the owner tenant of the rule
	TenantID string `json:"tenant_id"`

	// - Priority of the rule
	Priority int `json:"priority,omitempty"`

	// - ID of the target group that assigned to the rule
	TargetGroupID string `json:"target_group_id,omitempty"`

	// - Conditions of the rules to distribute accesses to the target groups
	Conditions ConditionInResponse `json:"conditions,omitempty"`

	// - Running configurations of the rule
	// - If `changes` is `true`, return object
	// - If current configuration does not exist, return `null`
	Current ConfigurationInResponse `json:"current,omitempty"`

	// - Added or changed configurations of the rule that waiting to be applied
	// - If `changes` is `true`, return object
	// - If staged configuration does not exist, return `null`
	Staged ConfigurationInResponse `json:"staged,omitempty"`
}

// ExtractInto interprets any commonResult as a rule, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "rule")
}

// Extract is a function that accepts a result and extracts a Rule resource.
func (r commonResult) Extract() (*Rule, error) {
	var rule Rule

	err := r.ExtractInto(&rule)

	return &rule, err
}

// RulePage is the page returned by a pager when traversing over a collection of rule.
type RulePage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a RulePage struct is empty.
func (r RulePage) IsEmpty() (bool, error) {
	is, err := ExtractRules(r)

	return len(is) == 0, err
}

// ExtractRulesInto interprets the results of a single page from a List() call, producing a slice of rule entities.
func ExtractRulesInto(r pagination.Page, v interface{}) error {
	return r.(RulePage).Result.ExtractIntoSlicePtr(v, "rules")
}

// ExtractRules accepts a Page struct, specifically a NetworkPage struct, and extracts the elements into a slice of Rule structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractRules(r pagination.Page) ([]Rule, error) {
	var s []Rule

	err := ExtractRulesInto(r, &s)

	return s, err
}
