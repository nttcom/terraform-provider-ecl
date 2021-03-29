package virtualstorages

import (
	"encoding/json"
	"time"

	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/pagination"
)

// IPAddressPool is struct which corresponds to ip_addr_pool object.
type IPAddressPool struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// HostRoute is struct which corresponds to host_routes object.
type HostRoute struct {
	Destination string `json:"destination"`
	Nexthop     string `json:"nexthop"`
}

// VirtualStorage contains all the information associated with a Virtual Storage.
type VirtualStorage struct {
	// API error in virtual storage creation.
	APIErrorMessage string `json:"api_error_message"`
	// Unique identifier for the virtual storage.
	ID string `json:"id"`
	// network_id which this virtual storage is connected.
	NetworkID string `json:"network_id"`
	// subnet_id which this virtual storage is connected.
	SubnetID string `json:"subnet_id"`
	// ip_address_pool object for virtual storage.
	IPAddrPool IPAddressPool `json:"ip_addr_pool"`
	// List of host routes of virtual storage.
	HostRoutes []HostRoute `json:"host_routes"`
	// volume_type_id of virtual storage
	VolumeTypeID string `json:"volume_type_id"`
	// Human-readable display name for the virtual storage.
	Name string `json:"name"`
	// Human-readable description for the virtual storage.
	Description string `json:"description"`
	// Current status of the virtual storage.
	Status string `json:"status"`
	// The date when this volume was created.
	CreatedAt time.Time `json:"-"`
	// The date when this volume was last updated
	UpdatedAt time.Time `json:"-"`
	// Error in virtual storage creation.
	ErrorMessage string `json:"error_message"`
}

// UnmarshalJSON creates JSON format of virtual storage
func (r *VirtualStorage) UnmarshalJSON(b []byte) error {
	type tmp VirtualStorage
	var s struct {
		tmp
		CreatedAt eclcloud.JSONISO8601 `json:"created_at"`
		UpdatedAt eclcloud.JSONISO8601 `json:"updated_at"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = VirtualStorage(s.tmp)

	r.CreatedAt = time.Time(s.CreatedAt)
	r.UpdatedAt = time.Time(s.UpdatedAt)

	return err
}

// VirtualStoragePage is a pagination.pager that is returned from a call to the List function.
type VirtualStoragePage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if a ListResult contains no VirtualStorages.
func (r VirtualStoragePage) IsEmpty() (bool, error) {
	vss, err := ExtractVirtualStorages(r)
	return len(vss) == 0, err
}

// ExtractVirtualStorages extracts and returns VirtualStorages.
// It is used while iterating over a virtualstorages.List call.
func ExtractVirtualStorages(r pagination.Page) ([]VirtualStorage, error) {
	var s []VirtualStorage
	err := ExtractVirtualStoragesInto(r, &s)
	return s, err
}

type commonResult struct {
	eclcloud.Result
}

// Extract will get the VirtualStorage object out of the commonResult object.
func (r commonResult) Extract() (*VirtualStorage, error) {
	var s VirtualStorage
	err := r.ExtractInto(&s)
	return &s, err
}

func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "virtual_storage")
}

// ExtractVirtualStoragesInto is information expander for virtual storage
func ExtractVirtualStoragesInto(r pagination.Page, v interface{}) error {
	return r.(VirtualStoragePage).Result.ExtractIntoSlicePtr(v, "virtual_storages")
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
