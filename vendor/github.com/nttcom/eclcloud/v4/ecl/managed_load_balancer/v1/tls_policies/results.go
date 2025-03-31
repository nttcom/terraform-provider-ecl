package tls_policies

import (
	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// ShowResult represents the result of a Show operation.
// Call its Extract method to interpret it as a TLSPolicy.
type ShowResult struct {
	commonResult
}

// TLSPolicy represents a tls policy.
type TLSPolicy struct {

	// - ID of the TLS policy
	ID string `json:"id"`

	// - Name of the TLS policy
	Name string `json:"name"`

	// - Description of the TLS policy
	Description string `json:"description"`

	// - Whether the TLS policy will be set `policy.tls_policy_id` when that is not specified
	Default bool `json:"default"`

	// - The list of acceptable TLS protocols in the policy that specifed this TLS policty
	TLSProtocols []string `json:"tls_protocols"`

	// - The list of acceptable TLS ciphers in the policy that specifed this TLS policty
	TLSCiphers []string `json:"tls_ciphers"`
}

// ExtractInto interprets any commonResult as a tls policy, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "tls_policy")
}

// Extract is a function that accepts a result and extracts a TLSPolicy resource.
func (r commonResult) Extract() (*TLSPolicy, error) {
	var tLSPolicy TLSPolicy

	err := r.ExtractInto(&tLSPolicy)

	return &tLSPolicy, err
}

// TLSPolicyPage is the page returned by a pager when traversing over a collection of tls policy.
type TLSPolicyPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a TLSPolicyPage struct is empty.
func (r TLSPolicyPage) IsEmpty() (bool, error) {
	is, err := ExtractTLSPolicies(r)

	return len(is) == 0, err
}

// ExtractTLSPoliciesInto interprets the results of a single page from a List() call, producing a slice of tls policy entities.
func ExtractTLSPoliciesInto(r pagination.Page, v interface{}) error {
	return r.(TLSPolicyPage).Result.ExtractIntoSlicePtr(v, "tls_policies")
}

// ExtractTLSPolicies accepts a Page struct, specifically a NetworkPage struct, and extracts the elements into a slice of TLSPolicy structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractTLSPolicies(r pagination.Page) ([]TLSPolicy, error) {
	var s []TLSPolicy

	err := ExtractTLSPoliciesInto(r, &s)

	return s, err
}
