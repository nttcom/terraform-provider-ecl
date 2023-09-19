package volumes

import (
	"encoding/json"
	"time"

	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/pagination"
)

// Volume contains all the information associated with a Volume.
type Volume struct {
	// API error in volume creation.
	APIErrorMessage string `json:"api_error_message"`
	// Unique identifier for the volume.
	ID string `json:"id"`
	// Current status of the volume.
	Status string `json:"status"`
	// Human-readable display name for the volume.
	Name string `json:"name"`
	// Human-readable description for the volume.
	Description string `json:"description"`
	// The volume size
	Size int `json:"size"`
	// The volume IOPS GB
	IOPSPerGB string `json:"iops_per_gb"`
	// The volume Throughput
	Throughput string `json:"throughput"`
	// The initiator_iqns for volume (in case ISCSI)
	InitiatorIQNs []string `json:"initiator_iqns"`
	// Relevant snapshot's IDs of this volume
	SnapshotIDs []string `json:"snapshot_ids"`
	// IP Addresses to connect this volume as target device.
	TargetIPs []string `json:"target_ips"`
	// The metadata of volume
	Metadata map[string]string `json:"metadata"`
	// The parent virtual storage ID to connect volume
	VirtualStorageID string `json:"virtual_storage_id"`
	// The availability zone of volume
	AvailabilityZone string `json:"availability_zone"`
	// The date when this volume was created.
	CreatedAt time.Time `json:"-"`
	// The date when this volume was last updated
	UpdatedAt time.Time `json:"-"`
	// Export rule of the volum
	ExportRules []string `json:"export_rules"`
	// Reservation parcentage about snapshot reservation capacity of the volume
	PercentSnapshotReserveUsed int `json:"percent_snapshot_reserve_used"`
	// Error in volume creation.
	ErrorMessage string `json:"error_message"`
}

// UnmarshalJSON creates JSON format of volume
func (r *Volume) UnmarshalJSON(b []byte) error {
	type tmp Volume
	var s struct {
		tmp
		CreatedAt eclcloud.JSONISO8601 `json:"created_at"`
		UpdatedAt eclcloud.JSONISO8601 `json:"updated_at"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Volume(s.tmp)

	r.CreatedAt = time.Time(s.CreatedAt)
	r.UpdatedAt = time.Time(s.UpdatedAt)

	return err
}

// VolumePage is a pagination.pager that is returned from a call to the List function.
type VolumePage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if a ListResult contains no Volumes.
func (r VolumePage) IsEmpty() (bool, error) {
	vss, err := ExtractVolumes(r)
	return len(vss) == 0, err
}

// ExtractVolumes extracts and returns Volumes.
// It is used while iterating over a Volumes.List call.
func ExtractVolumes(r pagination.Page) ([]Volume, error) {
	var s []Volume
	err := ExtractVolumesInto(r, &s)
	return s, err
}

type commonResult struct {
	eclcloud.Result
}

// Extract will get the Volume object out of the commonResult object.
func (r commonResult) Extract() (*Volume, error) {
	var s Volume
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "volume")
}

// ExtractVolumesInto is information expander for volume
func ExtractVolumesInto(r pagination.Page, v interface{}) error {
	return r.(VolumePage).Result.ExtractIntoSlicePtr(v, "volumes")
}

// CreateResult contains the response body and error from a Create request.
type CreateResult struct {
	commonResult
}

// GetResult contains the response body and error from a Get request.
type GetResult struct {
	commonResult
}

// UpdateResult contains the response body and error from an Update request.
type UpdateResult struct {
	commonResult
}

// DeleteResult contains the response body and error from a Delete request.
type DeleteResult struct {
	eclcloud.ErrResult
}
