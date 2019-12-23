package volumetypes

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

// ExtraSpec is struct which corresponds to extra_specs object.
type ExtraSpec struct {
	AvailableVolumeSize       []int    `json:"available_volume_size"`
	AvailableVolumeThroughput []string `json:"available_volume_throughput"`
	AvailableIOPSPerGB        []string `json:"available_iops_per_gb"`
}

// VolumeType contains all the information associated with a Virtual Storage.
type VolumeType struct {
	// API error in virtual storage creation.
	APIErrorMessage string `json:"api_error_message"`
	// Unique identifier for the volume type.
	ID string `json:"id"`
	// Human-readable display name for the volume type.
	Name string `json:"name"`
	// Extra specification of volume type.
	// This includes available_volume_size, and available_iops_per_gb,
	// or available_throughput depending on storage service type.
	ExtraSpecs ExtraSpec `json:"extra_specs"`
}

// VolumeTypePage is a pagination.pager that is returned from a call to the List function.
type VolumeTypePage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if a ListResult contains no VirtualStorages.
func (r VolumeTypePage) IsEmpty() (bool, error) {
	vtypes, err := ExtractVolumeTypes(r)
	return len(vtypes) == 0, err
}

// ExtractVolumeTypes extracts and returns VolumeTypes.
// It is used while iterating over a volumetypes.List call.
func ExtractVolumeTypes(r pagination.Page) ([]VolumeType, error) {
	var s []VolumeType
	err := ExtractVolumeTypesInto(r, &s)
	return s, err
}

type commonResult struct {
	eclcloud.Result
}

// Extract will get the VolumeType object out of the commonResult object.
func (r commonResult) Extract() (*VolumeType, error) {
	var s VolumeType
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "volume_type")
}

// ExtractVolumeTypesInto is information expander for volume types
func ExtractVolumeTypesInto(r pagination.Page, v interface{}) error {
	return r.(VolumeTypePage).Result.ExtractIntoSlicePtr(v, "volume_types")
}

// GetResult contains the response body and error from a Get request.
type GetResult struct {
	commonResult
}
