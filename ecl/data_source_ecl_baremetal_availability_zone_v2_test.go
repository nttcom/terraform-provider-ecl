package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccBaremetalV2AvailabilityZoneDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccBaremetalV2AvailabilityZoneDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBaremetalV2AvailabilityZoneDataSourceID("data.ecl_baremetal_availability_zone_v2.zone_groupa"),
					resource.TestCheckResourceAttr(
						"data.ecl_baremetal_availability_zone_v2.zone_groupa", "zone_name", "groupa"),
				),
			},
		},
	})
}

func testAccCheckBaremetalV2AvailabilityZoneDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find availability zone data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Availability zone data source ID not set")
		}

		return nil
	}
}

const testAccBaremetalV2AvailabilityZoneDataSourceBasic = `
data "ecl_baremetal_availability_zone_v2" "zone_groupa" {
  zone_name = "groupa"
}
`
