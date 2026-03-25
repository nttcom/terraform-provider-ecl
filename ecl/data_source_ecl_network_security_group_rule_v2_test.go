package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccNetworkV2SecurityGroupRuleDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2SecurityGroupRuleDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SecurityGroupRuleDataSourceID("data.ecl_network_security_group_rule_v2.rule_1"),
					resource.TestCheckResourceAttr("data.ecl_network_security_group_rule_v2.rule_1", "direction", "ingress"),
					resource.TestCheckResourceAttr("data.ecl_network_security_group_rule_v2.rule_1", "ethertype", "IPv4"),
					resource.TestCheckResourceAttr("data.ecl_network_security_group_rule_v2.rule_1", "protocol", "tcp"),
					resource.TestCheckResourceAttr("data.ecl_network_security_group_rule_v2.rule_1", "port_range_min", "22"),
					resource.TestCheckResourceAttr("data.ecl_network_security_group_rule_v2.rule_1", "port_range_max", "22"),
				),
			},
		},
	})
}

func TestAccNetworkV2SecurityGroupRuleDataSource_withRemoteGroupId(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2SecurityGroupRuleDataSourceWithRemoteGroupId,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SecurityGroupRuleDataSourceID("data.ecl_network_security_group_rule_v2.rule_1"),
					resource.TestCheckResourceAttr("data.ecl_network_security_group_rule_v2.rule_1", "direction", "ingress"),
					resource.TestCheckResourceAttr("data.ecl_network_security_group_rule_v2.rule_1", "ethertype", "IPv4"),
					resource.TestCheckResourceAttr("data.ecl_network_security_group_rule_v2.rule_1", "protocol", "tcp"),
					resource.TestCheckResourceAttrSet("data.ecl_network_security_group_rule_v2.rule_1", "remote_group_id"),
				),
			},
		},
	})
}

func testAccCheckNetworkV2SecurityGroupRuleDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find security group rule data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Security group rule data source ID not set")
		}

		return nil
	}
}

const testAccNetworkV2SecurityGroupRuleDataSourceBasic = `
resource "ecl_network_security_group_v2" "secgroup_1" {
  name        = "secgroup_1"
  description = "security group for rule testing"
}

resource "ecl_network_security_group_rule_v2" "rule_1" {
  security_group_id = ecl_network_security_group_v2.secgroup_1.id
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 22
  port_range_max    = 22
  remote_ip_prefix  = "0.0.0.0/0"
  description       = "Allow SSH"
}

data "ecl_network_security_group_rule_v2" "rule_1" {
  security_group_rule_id = ecl_network_security_group_rule_v2.rule_1.id
}
`

const testAccNetworkV2SecurityGroupRuleDataSourceWithRemoteGroupId = `
resource "ecl_network_security_group_v2" "secgroup_1" {
  name        = "secgroup_1"
  description = "security group 1"
}

resource "ecl_network_security_group_v2" "secgroup_2" {
  name        = "secgroup_2"
  description = "security group 2"
}

resource "ecl_network_security_group_rule_v2" "rule_1" {
  security_group_id = ecl_network_security_group_v2.secgroup_1.id
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "tcp"
  port_range_min    = 443
  port_range_max    = 443
  remote_group_id   = ecl_network_security_group_v2.secgroup_2.id
  description       = "Allow HTTPS from secgroup_2"
}

data "ecl_network_security_group_rule_v2" "rule_1" {
  security_group_rule_id = ecl_network_security_group_rule_v2.rule_1.id
}
`
