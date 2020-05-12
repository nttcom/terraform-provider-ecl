package ecl

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/ecl/dns/v2/recordsets"
)

func randomZoneName() string {
	return fmt.Sprintf("ACPTTEST-zone-%s.com.", acctest.RandString(5))
}

func TestAccDNSV2RecordSet_basic(t *testing.T) {
	var recordset recordsets.RecordSet
	zoneName := randomZoneName()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDNSV2RecordSetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDNSV2RecordSetBasic(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDNSV2RecordSetExists("ecl_dns_recordset_v2.recordset_1", &recordset),
					// for recordset_1
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_1", "name", zoneName),
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_1", "description", "a record set 1"),
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_1", "ttl", "3000"),
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_1", "record", "10.1.0.0"),
					// for recordset_2
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_2", "name", zoneName),
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_2", "description", "a record set 2"),
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_2", "ttl", "3000"),
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_2", "record", "20.1.0.0"),
				),
			},
			resource.TestStep{
				Config: testAccDNSV2RecordSetUpdate(zoneName),
				Check: resource.ComposeTestCheckFunc(
					// for recordset_1
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_1", "name", zoneName),
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_1", "description", "a record set 1-update"),
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_1", "ttl", "3000"),
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_1", "record", "10.1.0.0"),
					// for recordset_2
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_2", "name", zoneName),
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_2", "description", ""),
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_2", "ttl", "86400"),
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_2", "record", "20.1.0.1"),
				),
			},
			resource.TestStep{
				Config: testAccDNSV2RecordSetUpdate2(zoneName),
				Check: resource.ComposeTestCheckFunc(
					// for recordset_1
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_1", "name", zoneName),
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_1", "description", ""),
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_1", "ttl", "3000"),
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_1", "record", "10.1.0.0"),
					// for recordset_2
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_2", "name", maxLengthDomainNameFromZoneName(zoneName)),
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_2", "description", maxLengthDescription()),
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_2", "ttl", "0"),
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_2", "record", "20.1.0.1"),
				),
			},
		},
	})
}

func TestAccDNSV2RecordSet_ipv6(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var recordset recordsets.RecordSet
	zoneName := randomZoneName()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDNSV2RecordSetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDNSV2RecordSetIPv6(zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDNSV2RecordSetExists("ecl_dns_recordset_v2.recordset_1", &recordset),
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_1", "description", "a record set"),
					resource.TestCheckResourceAttr(
						"ecl_dns_recordset_v2.recordset_1", "record", "fd2b:db7f:6ae:dd8d::1"),
					// in ECL2.0 dns, even multiple record is requested, API only returns one record as response.
					// resource.TestCheckResourceAttr(
					// 	"ecl_dns_recordset_v2.recordset_1", "records.1", "fd2b:db7f:6ae:dd8d::2"),
				),
			},
		},
	})
}

func testAccCheckDNSV2RecordSetDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	dnsClient, err := config.dnsV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL DNS client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_dns_recordset_v2" {
			continue
		}

		zoneID, recordsetID, err := parseDNSV2RecordSetID(rs.Primary.ID)
		if err != nil {
			return err
		}

		_, err = recordsets.Get(dnsClient, zoneID, recordsetID).Extract()
		if err == nil {
			return fmt.Errorf("Record set still exists")
		}
	}

	return nil
}

func testAccCheckDNSV2RecordSetExists(n string, recordset *recordsets.RecordSet) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		dnsClient, err := config.dnsV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating ECL DNS client: %s", err)
		}

		zoneID, recordsetID, err := parseDNSV2RecordSetID(rs.Primary.ID)
		if err != nil {
			return err
		}

		found, err := recordsets.Get(dnsClient, zoneID, recordsetID).Extract()
		if err != nil {
			return err
		}

		if found.ID != recordsetID {
			return fmt.Errorf("Record set not found")
		}

		*recordset = *found

		return nil
	}
}

func testAccDNSV2RecordSetBasic(zoneName string) string {
	return fmt.Sprintf(`
		resource "ecl_dns_zone_v2" "zone_1" {
			name = "%s"
			email = "email2@example.com"
			description = "a zone"
			ttl = 6000
			type = "PRIMARY"
		}

		resource "ecl_dns_recordset_v2" "recordset_1" {
			zone_id = "${ecl_dns_zone_v2.zone_1.id}"
			type = "A"
			name = "%s"
			description = "a record set 1"
			ttl = 3000
			record = "10.1.0.0"
		}

		resource "ecl_dns_recordset_v2" "recordset_2" {
			zone_id = "${ecl_dns_zone_v2.zone_1.id}"
			type = "A"
			name = "%s"
			description = "a record set 2"
			ttl = 3000
			record = "20.1.0.0"
		}
	`, zoneName, zoneName, zoneName)
}

func testAccDNSV2RecordSetUpdate(zoneName string) string {
	return fmt.Sprintf(`
		resource "ecl_dns_zone_v2" "zone_1" {
			name = "%s"
			email = "email2@example.com"
			description = "a zone"
			ttl = 6000
			type = "PRIMARY"
		}

		resource "ecl_dns_recordset_v2" "recordset_1" {
			zone_id = "${ecl_dns_zone_v2.zone_1.id}"
			type = "A"
			name = "%s"
			description = "a record set 1-update"
			ttl = 3000
			record = "10.1.0.0"
		}

		resource "ecl_dns_recordset_v2" "recordset_2" {
			zone_id = "${ecl_dns_zone_v2.zone_1.id}"
			type = "A"
			name = "%s"
			description = ""
			ttl = 86400
			record = "20.1.0.1"
		}
	`, zoneName, zoneName, zoneName)
}

func testAccDNSV2RecordSetUpdate2(zoneName string) string {
	return fmt.Sprintf(`
		resource "ecl_dns_zone_v2" "zone_1" {
			name = "%s"
			email = "email2@example.com"
			description = "a zone"
			ttl = 6000
			type = "PRIMARY"
		}

		resource "ecl_dns_recordset_v2" "recordset_1" {
			zone_id = "${ecl_dns_zone_v2.zone_1.id}"
			type = "A"
			name = "%s"
			description = ""
			ttl = 3000
			record = "10.1.0.0"
		}

		resource "ecl_dns_recordset_v2" "recordset_2" {
			zone_id = "${ecl_dns_zone_v2.zone_1.id}"
			type = "A"
			name = "%s"
			description = "%s"
			ttl = 0
			record = "20.1.0.1"
		}
	`,
		zoneName,
		zoneName,
		maxLengthDomainNameFromZoneName(zoneName),
		maxLengthDescription(),
	)
}

func testAccDNSV2RecordSetIPv6(zoneName string) string {
	return fmt.Sprintf(`
		resource "ecl_dns_zone_v2" "zone_1" {
			name = "%s"
			email = "email2@example.com"
			description = "a zone"
			ttl = 6000
			type = "PRIMARY"
		}

		resource "ecl_dns_recordset_v2" "recordset_1" {
			zone_id = "${ecl_dns_zone_v2.zone_1.id}"
			name = "%s"
			type = "AAAA"
			description = "a record set"
			ttl = 3000
			record = "fd2b:db7f:6ae:dd8d::1"
		}
	`, zoneName, zoneName)
}

func maxLengthDomainNameFromZoneName(zoneName string) string {
	return strings.Repeat("a", 63) + "." + zoneName
}

func maxLengthDescription() string {
	return strings.Repeat("a", 255)
}
