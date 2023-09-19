package servers

import (
	"encoding/json"
	"time"

	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// GetResult is the result of Get operations. Call its Extract method to
// interpret it as a Server.
type GetResult struct {
	commonResult
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Server.
type CreateResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Server.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// Extract provides access to the individual Server returned by
// the Get and functions.
func (r commonResult) Extract() (*Server, error) {
	var s struct {
		Server *Server `json:"server"`
	}
	err := r.ExtractInto(&s)
	return s.Server, err
}

// RaidArray represents raid configuration for the server resource.
type RaidArray struct {
	PrimaryStorage     bool        `json:"primary_storage"`
	RaidCardHardwareID string      `json:"raid_card_hardware_id"`
	DiskHardwareIDs    []string    `json:"disk_hardware_ids"`
	RaidLevel          int         `json:"raid_level"`
	Partitions         []Partition `json:"partitions"`
}

// Partition represents partition configuration for the server resource.
type Partition struct {
	LVM            bool   `json:"lvm"`
	Size           int    `json:"size"`
	PartitionLabel string `json:"partition_label"`
}

// LVMVolumeGroup represents LVM volume group configuration for the server resource.
type LVMVolumeGroup struct {
	VGLabel                       string          `json:"vg_label"`
	PhysicalVolumePartitionLabels []string        `json:"physical_volume_partition_labels"`
	LogicalVolumes                []LogicalVolume `json:"logical_volumes"`
}

// LogicalVolume represents logical volume configuration for the server resource.
type LogicalVolume struct {
	LVLabel string `json:"lv_label"`
	Size    int    `json:"size"`
}

// Filesystem represents file system configuration for the server resource.
type Filesystem struct {
	Label      string `json:"label"`
	FSType     string `json:"fs_type"`
	MountPoint string `json:"mount_point"`
}

// NICPhysicalPort represents port configuraion for the server resource.
type NICPhysicalPort struct {
	ID                    string         `json:"id"`
	MacAddr               string         `json:"mac_addr"`
	NetworkPhysicalPortID string         `json:"network_physical_port_id"`
	Plane                 string         `json:"plane"`
	AttachedPorts         []AttachedPort `json:"attached_ports"`
	HardwareID            string         `json:"hardware_id"`
}

// AttachedPort represents attached port configuration for the server resource.
type AttachedPort struct {
	PortID    string    `json:"port_id"`
	NetworkID string    `json:"network_id"`
	FixedIPs  []FixedIP `json:"fixed_ips"`
}

// FixedIP represents fixed IP configuration for the server resource.
type FixedIP struct {
	SubnetID  string `json:"subnet_id"`
	IPAddress string `json:"ip_address"`
}

// ChassisStatus represents chassis status for the server resource
type ChassisStatus struct {
	ChassisPower bool `json:"chassis-power"`
	PowerSupply  bool `json:"power-supply"`
	CPU          bool `json:"cpu"`
	Memory       bool `json:"memory"`
	Fan          bool `json:"fan"`
	Disk         int  `json:"disk"`
	NIC          bool `json:"nic"`
	SystemBoard  bool `json:"system-board"`
	Etc          bool `json:"etc"`
}

// Personality represents personal files configuration for the server resource.
type Personality struct {
	Path     string `json:"path"`
	Contents string `json:"contents"`
}

// Server represents hardware configurations for server resources
// in a region.
type Server struct {
	ID               string                   `json:"id"`
	TenantID         string                   `json:"tenant_id"`
	UserID           string                   `json:"user_id"`
	Name             string                   `json:"name"`
	Updated          time.Time                `json:"-"`
	Created          time.Time                `json:"-"`
	Status           string                   `json:"status"`
	AdminPass        string                   `json:"adminPass"`
	PowerState       string                   `json:"OS-EXT-STS:power_state"`
	TaskState        string                   `json:"OS-EXT-STS:task_state"`
	VMState          string                   `json:"OS-EXT-STS:vm_state"`
	AvailabilityZone string                   `json:"OS-EXT-AZ:availability_zone"`
	Progress         int                      `json:"progress"`
	Image            map[string]interface{}   `json:"image"`
	Flavor           map[string]interface{}   `json:"flavor"`
	Metadata         map[string]string        `json:"metadata"`
	Links            []eclcloud.Link          `json:"links"`
	RaidArrays       []RaidArray              `json:"raid_arrays"`
	LVMVolumeGroups  []LVMVolumeGroup         `json:"lvm_volume_groups"`
	Filesystems      []Filesystem             `json:"filesystems"`
	NICPhysicalPorts []NICPhysicalPort        `json:"nic_physical_ports"`
	ChassisStatus    ChassisStatus            `json:"chassis-status"`
	MediaAttachments []map[string]interface{} `json:"media_attachments"`
	Personality      []Personality            `json:"personality"`
}

// UnmarshalJSON to override default
func (r *Server) UnmarshalJSON(b []byte) error {
	type tmp Server
	var s struct {
		tmp
		Created eclcloud.JSONRFC3339Milli `json:"created"`
		Updated eclcloud.JSONRFC3339Milli `json:"updated"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Server(s.tmp)

	r.Created = time.Time(s.Created)
	r.Updated = time.Time(s.Updated)

	return err
}

// ServerPage contains a single page of all servers from a ListDetails call.
type ServerPage struct {
	pagination.LinkedPageBase
}

// NextPageURL uses the response's embedded link reference to navigate to the
// next page of results.
func (page ServerPage) NextPageURL() (string, error) {
	var s struct {
		Links []eclcloud.Link `json:"servers_links"`
	}
	err := page.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return eclcloud.ExtractNextURL(s.Links)
}

// IsEmpty determines if a FlavorPage contains any results.
func (page ServerPage) IsEmpty() (bool, error) {
	flavors, err := ExtractServers(page)
	return len(flavors) == 0, err
}

// ExtractServers provides access to the list of flavors in a page acquired
// from the ListDetail operation.
func ExtractServers(r pagination.Page) ([]Server, error) {
	var s struct {
		Servers []Server `json:"servers"`
	}
	err := (r.(ServerPage)).ExtractInto(&s)
	return s.Servers, err
}
