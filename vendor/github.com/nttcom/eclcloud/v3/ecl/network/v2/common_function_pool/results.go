package common_function_pool

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

// CommonFunctionPool represents a Common Function Pool. See package documentation for a top-level
// description of what this is.
type CommonFunctionPool struct {

	// Description is description
	Description string `json:"description"`

	// UUID representing the Common Function Pool.
	ID string `json:"id"`

	// Name of Common Function Pool
	Name string `json:"name"`
}

type commonResult struct {
	eclcloud.Result
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Common Function Pool.
type GetResult struct {
	commonResult
}

// CommonFunctionPoolPage is the page returned by a pager when traversing over a collection
// of common function pools.
type CommonFunctionPoolPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a CommonFunctionPoolPage struct is empty.
func (r CommonFunctionPoolPage) IsEmpty() (bool, error) {
	is, err := ExtractCommonFunctionPools(r)
	return len(is) == 0, err
}

// ExtractCommonFunctionPools accepts a Page struct, specifically a CommonFunctionPoolPage struct,
// and extracts the elements into a slice of Common Function Pool structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractCommonFunctionPools(r pagination.Page) ([]CommonFunctionPool, error) {
	var s struct {
		CommonFunctionPools []CommonFunctionPool `json:"common_function_pools"`
	}
	err := (r.(CommonFunctionPoolPage)).ExtractInto(&s)
	return s.CommonFunctionPools, err
}

// Extract is a function that accepts a result and extracts a Common Function Pool resource.
func (r commonResult) Extract() (*CommonFunctionPool, error) {
	var s struct {
		CommonFunctionPool *CommonFunctionPool `json:"common_function_pool"`
	}
	err := r.ExtractInto(&s)
	return s.CommonFunctionPool, err
}
