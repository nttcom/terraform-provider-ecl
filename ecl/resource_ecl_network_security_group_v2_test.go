package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/v3/ecl/network/v2/security_groups"
)

func TestAccNetworkV2SecurityGroup_basic(t *testing.T) {
	var securityGroup security_groups.SecurityGroup

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2SecurityGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2SecurityGroupBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SecurityGroupExists("ecl_network_security_group_v2.secgroup_1", &securityGroup),
					resource.TestCheckResourceAttr("ecl_network_security_group_v2.secgroup_1", "name", "secgroup_1"),
					resource.TestCheckResourceAttr("ecl_network_security_group_v2.secgroup_1", "description", "security group 1"),
				),
			},
			{
				Config: testAccNetworkV2SecurityGroupUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SecurityGroupExists("ecl_network_security_group_v2.secgroup_1", &securityGroup),
					resource.TestCheckResourceAttr("ecl_network_security_group_v2.secgroup_1", "name", "secgroup_1_updated"),
					resource.TestCheckResourceAttr("ecl_network_security_group_v2.secgroup_1", "description", "security group 1 updated"),
				),
			},
		},
	})
}

func TestAccNetworkV2SecurityGroup_noDescription(t *testing.T) {
	var securityGroup security_groups.SecurityGroup

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2SecurityGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2SecurityGroupNoDescription,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SecurityGroupExists("ecl_network_security_group_v2.secgroup_1", &securityGroup),
					resource.TestCheckResourceAttr("ecl_network_security_group_v2.secgroup_1", "name", "secgroup_1"),
					resource.TestCheckResourceAttr("ecl_network_security_group_v2.secgroup_1", "description", ""),
				),
			},
		},
	})
}

func TestAccNetworkV2SecurityGroup_withTags(t *testing.T) {
	var securityGroup security_groups.SecurityGroup

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetworkV2SecurityGroupDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkV2SecurityGroupWithTags,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SecurityGroupExists("ecl_network_security_group_v2.secgroup_1", &securityGroup),
					resource.TestCheckResourceAttr("ecl_network_security_group_v2.secgroup_1", "tags.environment", "test"),
					resource.TestCheckResourceAttr("ecl_network_security_group_v2.secgroup_1", "tags.managed_by", "terraform"),
				),
			},
			{
				Config: testAccNetworkV2SecurityGroupWithTagsUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetworkV2SecurityGroupExists("ecl_network_security_group_v2.secgroup_1", &securityGroup),
					resource.TestCheckResourceAttr("ecl_network_security_group_v2.secgroup_1", "tags.environment", "production"),
					resource.TestCheckResourceAttr("ecl_network_security_group_v2.secgroup_1", "tags.managed_by", "terraform"),
					resource.TestCheckResourceAttr("ecl_network_security_group_v2.secgroup_1", "tags.owner", "admin"),
				),
			},
		},
	})
}

func testAccCheckNetworkV2SecurityGroupDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	networkClient, err := config.networkV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL network client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_network_security_group_v2" {
			continue
		}

		_, err := security_groups.Get(networkClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("Security group still exists")
		}
	}

	return nil
}

func testAccCheckNetworkV2SecurityGroupExists(n string, securityGroup *security_groups.SecurityGroup) resource.TestCheckFunc {
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

		found, err := security_groups.Get(networkClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("Security group not found")
		}

		*securityGroup = *found

		return nil
	}
}

const testAccNetworkV2SecurityGroupBasic = `
resource "ecl_network_security_group_v2" "secgroup_1" {
  name        = "secgroup_1"
  description = "security group 1"
}
`

const testAccNetworkV2SecurityGroupUpdate = `
resource "ecl_network_security_group_v2" "secgroup_1" {
  name        = "secgroup_1_updated"
  description = "security group 1 updated"
}
`

const testAccNetworkV2SecurityGroupNoDescription = `
resource "ecl_network_security_group_v2" "secgroup_1" {
  name = "secgroup_1"
}
`

const testAccNetworkV2SecurityGroupWithTags = `
resource "ecl_network_security_group_v2" "secgroup_1" {
  name        = "secgroup_1"
  description = "security group with tags"
  
  tags = {
    environment = "test"
    managed_by  = "terraform"
  }
}
`

const testAccNetworkV2SecurityGroupWithTagsUpdate = `
resource "ecl_network_security_group_v2" "secgroup_1" {
  name        = "secgroup_1"
  description = "security group with tags"
  
  tags = {
    environment = "production"
    managed_by  = "terraform"
    owner       = "admin"
  }
}
`
