package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccComputeV2FlavorDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2FlavorDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2FlavorDataSourceID("data.ecl_compute_flavor_v2.flavor_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_compute_flavor_v2.flavor_1", "name", "1CPU-2GB"),
					resource.TestCheckResourceAttr(
						"data.ecl_compute_flavor_v2.flavor_1", "ram", "2048"),
					resource.TestCheckResourceAttr(
						"data.ecl_compute_flavor_v2.flavor_1", "disk", "0"),
					resource.TestCheckResourceAttr(
						"data.ecl_compute_flavor_v2.flavor_1", "vcpus", "1"),
					resource.TestCheckResourceAttr(
						"data.ecl_compute_flavor_v2.flavor_1", "rx_tx_factor", "1"),
					resource.TestCheckResourceAttr(
						"data.ecl_compute_flavor_v2.flavor_1", "is_public", "true"),
				),
			},
		},
	})
}

func TestAccComputeV2FlavorDataSourceTestQueries(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccComputeV2FlavorDataSourceQueryMinRAM,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2FlavorDataSourceID("data.ecl_compute_flavor_v2.flavor_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_compute_flavor_v2.flavor_1", "name", "1CPU-2GB"),
					resource.TestCheckResourceAttr(
						"data.ecl_compute_flavor_v2.flavor_1", "ram", "2048"),
					resource.TestCheckResourceAttr(
						"data.ecl_compute_flavor_v2.flavor_1", "disk", "0"),
					resource.TestCheckResourceAttr(
						"data.ecl_compute_flavor_v2.flavor_1", "vcpus", "1"),
					resource.TestCheckResourceAttr(
						"data.ecl_compute_flavor_v2.flavor_1", "rx_tx_factor", "1"),
					resource.TestCheckResourceAttr(
						"data.ecl_compute_flavor_v2.flavor_1", "is_public", "true"),
				),
			},
			resource.TestStep{
				Config: testAccComputeV2FlavorDataSourceQueryVCPUs,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckComputeV2FlavorDataSourceID("data.ecl_compute_flavor_v2.flavor_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_compute_flavor_v2.flavor_1", "name", "1CPU-2GB"),
					resource.TestCheckResourceAttr(
						"data.ecl_compute_flavor_v2.flavor_1", "ram", "2048"),
					resource.TestCheckResourceAttr(
						"data.ecl_compute_flavor_v2.flavor_1", "disk", "0"),
					resource.TestCheckResourceAttr(
						"data.ecl_compute_flavor_v2.flavor_1", "vcpus", "1"),
					resource.TestCheckResourceAttr(
						"data.ecl_compute_flavor_v2.flavor_1", "rx_tx_factor", "1"),
					resource.TestCheckResourceAttr(
						"data.ecl_compute_flavor_v2.flavor_1", "is_public", "true"),
				),
			},
		},
	})
}

func testAccCheckComputeV2FlavorDataSourceID(n string) resource.TestCheckFunc {
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

const testAccComputeV2FlavorDataSourceBasic = `
data "ecl_compute_flavor_v2" "flavor_1" {
  name = "1CPU-2GB"
}
`

const testAccComputeV2FlavorDataSourceQueryMinRAM = `
data "ecl_compute_flavor_v2" "flavor_1" {
  name = "1CPU-2GB"
  min_ram = 2048
}
`

const testAccComputeV2FlavorDataSourceQueryVCPUs = `
data "ecl_compute_flavor_v2" "flavor_1" {
  name = "1CPU-2GB"
  vcpus = 1
}
`
