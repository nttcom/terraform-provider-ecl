package availabilityzones

import (
	"github.com/nttcom/eclcloud/v3/pagination"
)

// ZoneState represents the current state of the availability zone.
type ZoneState struct {
	// Returns true if the availability zone is available
	Available bool `json:"available"`
}

// AvailabilityZone contains all the information associated with an ECL
// AvailabilityZone.
type AvailabilityZone struct {
	ZoneName  string      `json:"zoneName"`
	ZoneState ZoneState   `json:"zoneState"`
	Hosts     interface{} `json:"hosts"`
}

// AvailabilityZonePage stores a single page of all AvailabilityZone results
// from a List call.
// Use the ExtractAvailabilityZones function to convert the results to a slice of
// AvailabilityZones.
type AvailabilityZonePage struct {
	pagination.SinglePageBase
}

// ExtractAvailabilityZones returns a slice of AvailabilityZones contained in a
// single page of results.
func ExtractAvailabilityZones(r pagination.Page) ([]AvailabilityZone, error) {
	var s struct {
		AvailabilityZoneInfo []AvailabilityZone `json:"availabilityZoneInfo"`
	}
	err := (r.(AvailabilityZonePage)).ExtractInto(&s)
	return s.AvailabilityZoneInfo, err
}
