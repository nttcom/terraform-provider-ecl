package servers

import (
	"encoding/base64"

	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

// Get retrieves the server with the provided ID.
// To extract the Server object from the response,
// call the Extract method on the GetResult.
func Get(client *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToServerListQuery() (string, error)
}

// ListOpts holds options for listing servers.
// It is passed to the servers.List function.
type ListOpts struct {
	// ChangesSince is a time/date stamp for when the server last changed status.
	ChangesSince string `q:"changes-since"`

	// Image is the name of the image in URL format.
	Image string `q:"image"`

	// Flavor is the name of the flavor in URL format.
	Flavor string `q:"flavor"`

	// Name of the server as a string.
	Name string `q:"name"`

	// Status is the value of the status of the server so that you can filter on
	// "ACTIVE" for example.
	Status string `q:"status"`

	// Marker is a UUID of the server at which you want to set a marker.
	Marker string `q:"marker"`

	// Limit is an integer value for the limit of values to return.
	Limit int `q:"limit"`
}

// ToServerListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToServerListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns Server optionally limited by the conditions provided in ListOpts.
func List(client *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToServerListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}

	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return ServerPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToServerCreateMap() (map[string]interface{}, error)
}

// CreateOptsNetwork represents networks information in server creation.
type CreateOptsNetwork struct {
	UUID    string `json:"uuid,omitempty"`
	Port    string `json:"port,omitempty"`
	FixedIP string `json:"fixed_ip,omitempty"`
	Plane   string `json:"plane,omitempty"`
}

// CreateOptsRaidArray represents raid configuration for the server resource.
type CreateOptsRaidArray struct {
	PrimaryStorage     bool                  `json:"primary_storage,omitempty"`
	RaidCardHardwareID string                `json:"raid_card_hardware_id,omitempty"`
	DiskHardwareIDs    []string              `json:"disk_hardware_ids,omitempty"`
	RaidLevel          int                   `json:"raid_level,omitempty"`
	Partitions         []CreateOptsPartition `json:"partitions,omitempty"`
}

// CreateOptsPartition represents partition configuration for the server resource.
type CreateOptsPartition struct {
	LVM            bool   `json:"lvm,omitempty"`
	Size           string `json:"size,omitempty"`
	PartitionLabel string `json:"partition_label,omitempty"`
}

// CreateOptsLVMVolumeGroup represents LVM volume group configuration for the server resource.
type CreateOptsLVMVolumeGroup struct {
	VGLabel                       string                    `json:"vg_label,omitempty"`
	PhysicalVolumePartitionLabels []string                  `json:"physical_volume_partition_labels,omitempty"`
	LogicalVolumes                []CreateOptsLogicalVolume `json:"logical_volumes,omitempty"`
}

// CreateOptsLogicalVolume represents logical volume configuration for the server resource.
type CreateOptsLogicalVolume struct {
	LVLabel string `json:"lv_label,omitempty"`
	Size    string `json:"size,omitempty"`
}

// CreateOptsFilesystem represents file system configuration for the server resource.
type CreateOptsFilesystem struct {
	Label      string `json:"label,omitempty"`
	FSType     string `json:"fs_type,omitempty"`
	MountPoint string `json:"mount_point,omitempty"`
}

// CreateOptsPersonality represents personal files configuration for the server resource.
type CreateOptsPersonality struct {
	Path     string `json:"path,omitempty"`
	Contents string `json:"contents,omitempty"`
}

// CreateOpts represents options used to create a server.
type CreateOpts struct {
	Name             string                     `json:"name" required:"true"`
	Networks         []CreateOptsNetwork        `json:"networks" required:"true"`
	AdminPass        string                     `json:"adminPass,omitempty"`
	ImageRef         string                     `json:"imageRef,omitempty"`
	FlavorRef        string                     `json:"flavorRef" required:"true"`
	AvailabilityZone string                     `json:"availability_zone,omitempty"`
	KeyName          string                     `json:"key_name,omitempty"`
	UserData         []byte                     `json:"-"`
	RaidArrays       []CreateOptsRaidArray      `json:"raid_arrays,omitempty"`
	LVMVolumeGroups  []CreateOptsLVMVolumeGroup `json:"lvm_volume_groups,omitempty"`
	Filesystems      []CreateOptsFilesystem     `json:"filesystems,omitempty"`
	Metadata         map[string]string          `json:"metadata,omitempty"`
	Personality      []CreateOptsPersonality    `json:"personality,omitempty"`
}

// ToServerCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToServerCreateMap() (map[string]interface{}, error) {
	b, err := eclcloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	if opts.UserData != nil {
		var userData string
		if _, err := base64.StdEncoding.DecodeString(string(opts.UserData)); err != nil {
			userData = base64.StdEncoding.EncodeToString(opts.UserData)
		} else {
			userData = string(opts.UserData)
		}
		b["user_data"] = &userData
	}

	return map[string]interface{}{"server": b}, nil
}

// Create accepts a CreateOpts struct and creates a new server
// using the values provided.
// This operation does not actually require a request body, i.e. the
// CreateOpts struct argument can be empty.
func Create(c *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToServerCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(createURL(c), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// Delete requests that a server previously provisioned be removed from your
// account.
func Delete(client *eclcloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, id), nil)
	return
}
