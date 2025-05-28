package plans

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// ShowResult represents the result of a Show operation.
// Call its Extract method to interpret it as a Plan.
type ShowResult struct {
	commonResult
}

// Plan represents a plan.
type Plan struct {

	// - ID of the plan
	ID string `json:"id"`

	// - Name of the plan
	Name string `json:"name"`

	// - Description of the plan
	Description string `json:"description"`

	// - Bandwidth of the load balancer
	Bandwidth string `json:"bandwidth"`

	// - Redundancy of the load balancer
	Redundancy string `json:"redundancy"`

	// - Maximum number of interfaces for the load balancer
	MaxNumberOfInterfaces int `json:"max_number_of_interfaces"`

	// - Maximum number of health monitors for the load balancer
	MaxNumberOfHealthMonitors int `json:"max_number_of_health_monitors"`

	// - Maximum number of listeners for the load balancer
	MaxNumberOfListeners int `json:"max_number_of_listeners"`

	// - Maximum number of policies for the load balancer
	MaxNumberOfPolicies int `json:"max_number_of_policies"`

	// - Maximum number of routes for the load balancer
	MaxNumberOfRoutes int `json:"max_number_of_routes"`

	// - Maximum number of target groups for the load balancer
	MaxNumberOfTargetGroups int `json:"max_number_of_target_groups"`

	// - Maximum number of members for a target group
	MaxNumberOfMembers int `json:"max_number_of_members"`

	// - Maximum number of rules for a policy
	MaxNumberOfRules int `json:"max_number_of_rules"`

	// - Maximum number of conditions in a rule
	MaxNumberOfConditions int `json:"max_number_of_conditions"`

	// - Maximum number of Server Name Indications (SNIs) in a policy
	MaxNumberOfServerNameIndications int `json:"max_number_of_server_name_indications"`

	// - Whether a new load balancer can be created with this plan
	Enabled bool `json:"enabled"`
}

// ExtractInto interprets any commonResult as a plan, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "plan")
}

// Extract is a function that accepts a result and extracts a Plan resource.
func (r commonResult) Extract() (*Plan, error) {
	var plan Plan

	err := r.ExtractInto(&plan)

	return &plan, err
}

// PlanPage is the page returned by a pager when traversing over a collection of plan.
type PlanPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a PlanPage struct is empty.
func (r PlanPage) IsEmpty() (bool, error) {
	is, err := ExtractPlans(r)

	return len(is) == 0, err
}

// ExtractPlansInto interprets the results of a single page from a List() call, producing a slice of plan entities.
func ExtractPlansInto(r pagination.Page, v interface{}) error {
	return r.(PlanPage).Result.ExtractIntoSlicePtr(v, "plans")
}

// ExtractPlans accepts a Page struct, specifically a NetworkPage struct, and extracts the elements into a slice of Plan structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractPlans(r pagination.Page) ([]Plan, error) {
	var s []Plan

	err := ExtractPlansInto(r, &s)

	return s, err
}
