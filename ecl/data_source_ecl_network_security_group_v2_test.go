package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccNetworkV2SecurityGroupDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2SecurityGroupDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SecurityGroupDataSourceID("data.ecl_network_security_group_v2.secgroup_1"),
					resource.TestCheckResourceAttr("data.ecl_network_security_group_v2.secgroup_1", "name", "secgroup_1"),
					resource.TestCheckResourceAttr("data.ecl_network_security_group_v2.secgroup_1", "description", "security group 1"),
				),
			},
		},
	})
}

func TestAccNetworkV2SecurityGroupDataSource_byName(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2SecurityGroupDataSourceByName,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SecurityGroupDataSourceID("data.ecl_network_security_group_v2.secgroup_1"),
					resource.TestCheckResourceAttr("data.ecl_network_security_group_v2.secgroup_1", "name", "secgroup_1"),
				),
			},
		},
	})
}

func TestAccNetworkV2SecurityGroupDataSource_withTags(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2SecurityGroupDataSourceWithTags,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SecurityGroupDataSourceID("data.ecl_network_security_group_v2.secgroup_1"),
					resource.TestCheckResourceAttr("data.ecl_network_security_group_v2.secgroup_1", "name", "secgroup_1"),
					resource.TestCheckResourceAttr("data.ecl_network_security_group_v2.secgroup_1", "tags.environment", "test"),
				),
			},
		},
	})
}

func testAccCheckNetworkV2SecurityGroupDataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find security group data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Security group data source ID not set")
		}

		return nil
	}
}

const testAccNetworkV2SecurityGroupDataSourceBasic = `
resource "ecl_network_security_group_v2" "secgroup_1" {
  name        = "secgroup_1"
  description = "security group 1"
}

data "ecl_network_security_group_v2" "secgroup_1" {
  security_group_id = ecl_network_security_group_v2.secgroup_1.id
}
`

const testAccNetworkV2SecurityGroupDataSourceByName = `
resource "ecl_network_security_group_v2" "secgroup_1" {
  name        = "secgroup_1"
  description = "security group 1"
}

data "ecl_network_security_group_v2" "secgroup_1" {
  name = ecl_network_security_group_v2.secgroup_1.name
}
`

const testAccNetworkV2SecurityGroupDataSourceWithTags = `
resource "ecl_network_security_group_v2" "secgroup_1" {
  name        = "secgroup_1"
  description = "security group with tags"
  
  tags = {
    environment = "test"
  }
}

data "ecl_network_security_group_v2" "secgroup_1" {
  security_group_id = ecl_network_security_group_v2.secgroup_1.id
}
`
