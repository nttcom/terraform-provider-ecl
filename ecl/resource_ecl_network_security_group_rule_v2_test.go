package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/v3/ecl/network/v2/security_group_rules"
	"github.com/nttcom/eclcloud/v3/ecl/network/v2/security_groups"
)

func TestAccNetworkV2SecurityGroupRule_basic(t *testing.T) {
	var securityGroup security_groups.SecurityGroup
	var securityGroupRule security_group_rules.SecurityGroupRule

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2SecurityGroupRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2SecurityGroupRuleBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SecurityGroupExists("ecl_network_security_group_v2.secgroup_1", &securityGroup),
					testAccCheckNetworkV2SecurityGroupRuleExists("ecl_network_security_group_rule_v2.rule_1", &securityGroupRule),
					resource.TestCheckResourceAttr("ecl_network_security_group_rule_v2.rule_1", "direction", "ingress"),
					resource.TestCheckResourceAttr("ecl_network_security_group_rule_v2.rule_1", "ethertype", "IPv4"),
					resource.TestCheckResourceAttr("ecl_network_security_group_rule_v2.rule_1", "protocol", "tcp"),
					resource.TestCheckResourceAttr("ecl_network_security_group_rule_v2.rule_1", "port_range_min", "22"),
					resource.TestCheckResourceAttr("ecl_network_security_group_rule_v2.rule_1", "port_range_max", "22"),
					resource.TestCheckResourceAttr("ecl_network_security_group_rule_v2.rule_1", "remote_ip_prefix", "0.0.0.0/0"),
				),
			},
		},
	})
}

func TestAccNetworkV2SecurityGroupRule_defaultValues(t *testing.T) {
	var securityGroup security_groups.SecurityGroup
	var securityGroupRule security_group_rules.SecurityGroupRule

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2SecurityGroupRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2SecurityGroupRuleDefaultValues,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SecurityGroupExists("ecl_network_security_group_v2.secgroup_1", &securityGroup),
					testAccCheckNetworkV2SecurityGroupRuleExists("ecl_network_security_group_rule_v2.rule_1", &securityGroupRule),
					resource.TestCheckResourceAttr("ecl_network_security_group_rule_v2.rule_1", "direction", "ingress"),
					resource.TestCheckResourceAttr("ecl_network_security_group_rule_v2.rule_1", "ethertype", "IPv4"),
					resource.TestCheckResourceAttr("ecl_network_security_group_rule_v2.rule_1", "protocol", "any"),
					resource.TestCheckResourceAttr("ecl_network_security_group_rule_v2.rule_1", "port_range_min", "0"),
					resource.TestCheckResourceAttr("ecl_network_security_group_rule_v2.rule_1", "port_range_max", "65535"),
				),
			},
		},
	})
}

func TestAccNetworkV2SecurityGroupRule_icmpProtocol(t *testing.T) {
	var securityGroup security_groups.SecurityGroup
	var securityGroupRule security_group_rules.SecurityGroupRule

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2SecurityGroupRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2SecurityGroupRuleICMP,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SecurityGroupExists("ecl_network_security_group_v2.secgroup_1", &securityGroup),
					testAccCheckNetworkV2SecurityGroupRuleExists("ecl_network_security_group_rule_v2.rule_1", &securityGroupRule),
					resource.TestCheckResourceAttr("ecl_network_security_group_rule_v2.rule_1", "direction", "ingress"),
					resource.TestCheckResourceAttr("ecl_network_security_group_rule_v2.rule_1", "ethertype", "IPv4"),
					resource.TestCheckResourceAttr("ecl_network_security_group_rule_v2.rule_1", "protocol", "icmp"),
				),
			},
		},
	})
}

func TestAccNetworkV2SecurityGroupRule_egress(t *testing.T) {
	var securityGroup security_groups.SecurityGroup
	var securityGroupRule security_group_rules.SecurityGroupRule

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2SecurityGroupRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2SecurityGroupRuleEgress,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SecurityGroupExists("ecl_network_security_group_v2.secgroup_1", &securityGroup),
					testAccCheckNetworkV2SecurityGroupRuleExists("ecl_network_security_group_rule_v2.rule_1", &securityGroupRule),
					resource.TestCheckResourceAttr("ecl_network_security_group_rule_v2.rule_1", "direction", "egress"),
					resource.TestCheckResourceAttr("ecl_network_security_group_rule_v2.rule_1", "ethertype", "IPv4"),
					resource.TestCheckResourceAttr("ecl_network_security_group_rule_v2.rule_1", "remote_ip_prefix", "0.0.0.0/0"),
				),
			},
		},
	})
}

func TestAccNetworkV2SecurityGroupRule_remoteGroupId(t *testing.T) {
	var securityGroup1 security_groups.SecurityGroup
	var securityGroup2 security_groups.SecurityGroup
	var securityGroupRule security_group_rules.SecurityGroupRule

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2SecurityGroupRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2SecurityGroupRuleRemoteGroupId,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SecurityGroupExists("ecl_network_security_group_v2.secgroup_1", &securityGroup1),
					testAccCheckNetworkV2SecurityGroupExists("ecl_network_security_group_v2.secgroup_2", &securityGroup2),
					testAccCheckNetworkV2SecurityGroupRuleExists("ecl_network_security_group_rule_v2.rule_1", &securityGroupRule),
					resource.TestCheckResourceAttr("ecl_network_security_group_rule_v2.rule_1", "direction", "ingress"),
					resource.TestCheckResourceAttr("ecl_network_security_group_rule_v2.rule_1", "ethertype", "IPv4"),
					resource.TestCheckResourceAttr("ecl_network_security_group_rule_v2.rule_1", "protocol", "tcp"),
					resource.TestCheckResourceAttr("ecl_network_security_group_rule_v2.rule_1", "port_range_min", "80"),
					resource.TestCheckResourceAttr("ecl_network_security_group_rule_v2.rule_1", "port_range_max", "80"),
					resource.TestCheckResourceAttrSet("ecl_network_security_group_rule_v2.rule_1", "remote_group_id"),
				),
			},
		},
	})
}

func testAccCheckNetworkV2SecurityGroupRuleDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	networkClient, err := config.networkV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_network_security_group_rule_v2" {
			continue
		}

		_, err := security_group_rules.Get(networkClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("Security group rule still exists")
		}
	}

	return nil
}

func testAccCheckNetworkV2SecurityGroupRuleExists(n string, securityGroupRule *security_group_rules.SecurityGroupRule) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		networkClient, err := config.networkV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating ECL network client: %s", err)
		}

		found, err := security_group_rules.Get(networkClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("Security group rule not found")
		}

		*securityGroupRule = *found

		return nil
	}
}

const testAccNetworkV2SecurityGroupRuleBasic = `
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
`

const testAccNetworkV2SecurityGroupRuleDefaultValues = `
resource "ecl_network_security_group_v2" "secgroup_1" {
  name        = "secgroup_1"
  description = "security group for rule testing"
}

resource "ecl_network_security_group_rule_v2" "rule_1" {
  security_group_id = ecl_network_security_group_v2.secgroup_1.id
  direction         = "ingress"
  ethertype         = "IPv4"
  remote_ip_prefix  = "0.0.0.0/0"
}
`

const testAccNetworkV2SecurityGroupRuleICMP = `
resource "ecl_network_security_group_v2" "secgroup_1" {
  name        = "secgroup_1"
  description = "security group for rule testing"
}

resource "ecl_network_security_group_rule_v2" "rule_1" {
  security_group_id = ecl_network_security_group_v2.secgroup_1.id
  direction         = "ingress"
  ethertype         = "IPv4"
  protocol          = "icmp"
  remote_ip_prefix  = "0.0.0.0/0"
  description       = "Allow ICMP"
}
`

const testAccNetworkV2SecurityGroupRuleEgress = `
resource "ecl_network_security_group_v2" "secgroup_1" {
  name        = "secgroup_1"
  description = "security group for rule testing"
}

resource "ecl_network_security_group_rule_v2" "rule_1" {
  security_group_id = ecl_network_security_group_v2.secgroup_1.id
  direction         = "egress"
  ethertype         = "IPv4"
  remote_ip_prefix  = "0.0.0.0/0"
  description       = "Allow all outbound"
}
`

const testAccNetworkV2SecurityGroupRuleRemoteGroupId = `
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
  port_range_min    = 80
  port_range_max    = 80
  remote_group_id   = ecl_network_security_group_v2.secgroup_2.id
  description       = "Allow from secgroup_2"
}
`
