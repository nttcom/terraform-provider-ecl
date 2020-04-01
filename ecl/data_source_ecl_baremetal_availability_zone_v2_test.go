package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccBaremetalV2AvailabilityZoneDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckBaremetal(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccBaremetalV2AvailabilityZoneDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBaremetalV2AvailabilityZoneDataSourceID("data.ecl_baremetal_availability_zone_v2.zone_groupa"),
					resource.TestCheckResourceAttr(
						"data.ecl_baremetal_availability_zone_v2.zone_groupa", "zone_name", OS_BAREMETAL_AVAILABLE_ZONE),
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

var testAccBaremetalV2AvailabilityZoneDataSourceBasic = fmt.Sprintf(`
data "ecl_baremetal_availability_zone_v2" "zone_groupa" {
  zone_name = "%s"
}
`,
	OS_BAREMETAL_AVAILABLE_ZONE,
)
