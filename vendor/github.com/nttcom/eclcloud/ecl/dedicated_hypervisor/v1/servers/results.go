package servers

import (
	"time"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

// Server represents dedicated hypervisor server information.
type Server struct {
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	ImageRef        string          `json:"imageRef"`
	Description     *string         `json:"description"`
	Status          string          `json:"status"`
	HypervisorType  string          `json:"hypervisor_type"`
	BaremetalServer BaremetalServer `json:"baremetal_server"`
	Links           []Link          `json:"links"`
	AdminPass       string          `json:"adminPass"`
}

type BaremetalServer struct {
	PowerState               string            `json:"OS-EXT-STS:power_state"`
	TaskState                string            `json:"OS-EXT-STS:task_state"`
	VMState                  string            `json:"OS-EXT-STS:vm_state"`
	AvailabilityZone         string            `json:"OS-EXT-AZ:availability_zone"`
	Created                  time.Time         `json:"created"`
	Flavor                   Flavor            `json:"flavor"`
	ID                       string            `json:"id"`
	Image                    Image             `json:"image"`
	Metadata                 map[string]string `json:"metadata"`
	Name                     string            `json:"name"`
	Progress                 int               `json:"progress"`
	Status                   string            `json:"status"`
	TenantID                 string            `json:"tenant_id"`
	Updated                  time.Time         `json:"updated"`
	UserID                   string            `json:"user_id"`
	NicPhysicalPorts         []NicPhysicalPort `json:"nic_physical_ports"`
	ChassisStatus            ChassisStatus     `json:"chassis-status"`
	Links                    []Link            `json:"links"`
	RaidArrays               []RaidArray       `json:"raid_arrays"`
	LvmVolumeGroups          []LvmVolumeGroup  `json:"lvm_volume_groups"`
	Filesystems              []Filesystem      `json:"filesystems"`
	MediaAttachments         []MediaAttachment `json:"media_attachments"`
	ManagedByService         string            `json:"managed_by_service"`
	ManagedServiceResourceID string            `json:"managed_service_resource_id"`
}

type Flavor struct {
	ID    string `json:"id"`
	Links []Link `json:"links"`
}

type Image struct {
	ID    string `json:"id"`
	Links []Link `json:"links"`
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

type NicPhysicalPort struct {
	ID                    string `json:"id"`
	MACAddr               string `json:"mac_addr"`
	NetworkPhysicalPortID string `json:"network_physical_port_id"`
	Plane                 string `json:"plane"`
	AttachedPorts         []Port `json:"attached_ports"`
	HardwareID            string `json:"hardware_id"`
}

type Port struct {
	PortID    string    `json:"port_id"`
	NetworkID string    `json:"network_id"`
	FixedIPs  []FixedIP `json:"fixed_ips"`
}

type FixedIP struct {
	SubnetID  string `json:"subnet_id"`
	IPAddress string `json:"ip_address"`
}

type ChassisStatus struct {
	ChassisPower bool `json:"chassis-power"`
	PowerSupply  bool `json:"power-supply"`
	CPU          bool `json:"cpu"`
	Memory       bool `json:"memory"`
	Fan          bool `json:"fan"`
	Disk         int  `json:"disk"`
	Nic          bool `json:"nic"`
	SystemBoard  bool `json:"system-board"`
	Etc          bool `json:"etc"`
	Console      bool `json:"console"`
}

type RaidArray struct {
	PrimaryStorage     bool        `json:"primary_storage"`
	Partitions         []Partition `json:"partitions"`
	RaidCardHardwareID string      `json:"raid_card_hardware_id"`
	DiskHardwareIDs    []string    `json:"disk_hardware_ids"`
}

type Partition struct {
	Lvm            bool   `json:"lvm"`
	Size           string `json:"size"`
	PartitionLabel string `json:"partition_label"`
}

type LvmVolumeGroup struct {
	VgLabel                       string          `json:"vg_label"`
	PhysicalVolumePartitionLabels []string        `json:"physical_volume_partition_labels"`
	LogicalVolumes                []LogicalVolume `json:"logical_volumes"`
}

type LogicalVolume struct {
	LvLabel string `json:"lv_label"`
	Size    string `json:"size"`
}

type Filesystem struct {
	Label      string `json:"label"`
	MountPoint string `json:"mount_point"`
	FsType     string `json:"fs_type"`
}

type MediaAttachment struct {
	Image Image `json:"image"`
}

type commonResult struct {
	eclcloud.Result
}

// GetResult is the response from a Get operation. Call its Extract method
// to interpret it as a Server.
type GetResult struct {
	commonResult
}

// CreateResult is the response from a Create operation. Call its Extract method
// to interpret it as a Server.
type CreateResult struct {
	commonResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr to
// determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// ServerPage is a single page of Server results.
type ServerPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of Servers contains any results.
func (r ServerPage) IsEmpty() (bool, error) {
	servers, err := ExtractServers(r)
	return len(servers) == 0, err
}

// ExtractServers returns a slice of Servers contained in a single page of
// results.
func ExtractServers(r pagination.Page) ([]Server, error) {
	var s struct {
		Servers []Server `json:"servers"`
	}
	err := (r.(ServerPage)).ExtractInto(&s)
	return s.Servers, err
}

// Extract interprets any commonResult as a Server.
func (r commonResult) Extract() (*Server, error) {
	var s struct {
		Server *Server `json:"server"`
	}
	err := r.ExtractInto(&s)
	return s.Server, err
}

type Job struct {
	JobID          string         `json:"job_id"`
	Status         string         `json:"status"`
	RequestedParam RequestedParam `json:"requested_param"`
}

type RequestedParam struct {
	VmName       string   `json:"vm_name"`
	LicenseTypes []string `json:"license_types"`
}

type AddLicenseResult struct {
	eclcloud.Result
}

func (r AddLicenseResult) Extract() (*Job, error) {
	var job Job
	err := r.ExtractInto(&job)
	return &job, err
}
