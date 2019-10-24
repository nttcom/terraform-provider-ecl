package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccBaremetalV2FlavorDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccBaremetalV2FlavorDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBaremetalV2FlavorDataSourceID("data.ecl_baremetal_flavor_v2.flavor_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_baremetal_flavor_v2.flavor_1", "name", "General Purpose 1 v1"),
					resource.TestCheckResourceAttr(
						"data.ecl_baremetal_flavor_v2.flavor_1", "ram", "32768"),
					resource.TestCheckResourceAttr(
						"data.ecl_baremetal_flavor_v2.flavor_1", "disk", "550"),
					resource.TestCheckResourceAttr(
						"data.ecl_baremetal_flavor_v2.flavor_1", "vcpus", "4"),
				),
			},
		},
	})
}

func testAccCheckBaremetalV2FlavorDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find flavor data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Flavor data source ID not set")
		}

		return nil
	}
}

const testAccBaremetalV2FlavorDataSourceBasic = `
data "ecl_baremetal_flavor_v2" "flavor_1" {
  name = "General Purpose 1 v1"
}
`
