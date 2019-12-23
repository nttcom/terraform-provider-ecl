package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var zoneName = fmt.Sprintf("ACPTTEST%s.com.", acctest.RandString(5))

func TestAccDNSV2ZoneDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckDNS(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDNSV2ZoneDataSourceZone,
			},
			resource.TestStep{
				Config: testAccDNSV2ZoneDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDNSV2ZoneDataSourceID("data.ecl_dns_zone_v2.z1"),
					resource.TestCheckResourceAttr(
						"data.ecl_dns_zone_v2.z1", "name", zoneName),
				),
			},
		},
	})
}
func TestAccDNSV2ZoneDataSource_queryDomainName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckDNS(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDNSV2ZoneDataSourceZone,
			},
			resource.TestStep{
				Config: testAccDNSV2ZoneDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDNSV2ZoneDataSourceID("data.ecl_dns_zone_v2.z1"),
					resource.TestCheckResourceAttr(
						"data.ecl_dns_zone_v2.z1", "name", zoneName),
				),
			},
		},
	})
}

func testAccCheckDNSV2ZoneDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find DNS Zone data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("DNS Zone data source ID not set")
		}

		return nil
	}
}

var testAccDNSV2ZoneDataSourceZone = fmt.Sprintf(`
resource "ecl_dns_zone_v2" "z1" {
  name = "%s"
  email = "terraform-dns-zone-v2-test-name@example.com"
  type = "PRIMARY"
  ttl = 7200
}`, zoneName)

var testAccDNSV2ZoneDataSourceBasic = fmt.Sprintf(`
%s
data "ecl_dns_zone_v2" "z1" {
  domain_name = "${ecl_dns_zone_v2.z1.name}"
}
`, testAccDNSV2ZoneDataSourceZone)

var testAccDNSV2ZoneDataSourceQueryDomainName = fmt.Sprintf(`
%s
data "ecl_dns_zone_v2" "z1" {
  domain_name = "ACPTTEST"
}
`, testAccDNSV2ZoneDataSourceZone)
