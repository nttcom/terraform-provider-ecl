package ecl

import (
	"fmt"
	"testing"

	"github.com/nttcom/eclcloud"

	"github.com/nttcom/eclcloud/ecl/rca/v1/users"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccRCAV1User_basic(t *testing.T) {
	var user users.User

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckRCAV1UserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccRCAV1UserBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRCAV1UserExists("ecl_rca_user_v1.user_1", &user),
					resource.TestCheckResourceAttr("ecl_rca_user_v1.user_1", "password", "dummy_passw@rd"),
					resource.TestCheckResourceAttrSet("ecl_rca_user_v1.user_1", "name"),
					resource.TestCheckResourceAttrSet("ecl_rca_user_v1.user_1", "vpn_endpoints.0.endpoint"),
					resource.TestCheckResourceAttrSet("ecl_rca_user_v1.user_1", "vpn_endpoints.0.type"),
				),
			},
			{
				Config: testAccRCAV1UserUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ecl_rca_user_v1.user_1", "password", "dummy_passw@rd_updated"),
					resource.TestCheckResourceAttrSet("ecl_rca_user_v1.user_1", "name"),
					resource.TestCheckResourceAttrSet("ecl_rca_user_v1.user_1", "vpn_endpoints.0.endpoint"),
					resource.TestCheckResourceAttrSet("ecl_rca_user_v1.user_1", "vpn_endpoints.0.type"),
				),
			},
		},
	})
}

func testAccCheckRCAV1UserDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.rcaV1Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("error creating ECL RCA client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_rca_user_v1" {
			continue
		}

		if _, err := users.Get(client, rs.Primary.ID).Extract(); err != nil {
			if _, ok := err.(eclcloud.ErrDefault404); ok {
				continue
			}

			return fmt.Errorf("error getting ECL RCA user: %s", err)
		}

		return fmt.Errorf("user still exists")
	}

	return nil
}

func testAccCheckRCAV1UserExists(n string, user *users.User) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		client, err := config.rcaV1Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("error creating ECL RCA client: %s", err)
		}

		found, err := users.Get(client, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.Name != rs.Primary.ID {
			return fmt.Errorf("user not found")
		}

		*user = *found

		return nil
	}
}

const testAccRCAV1UserBasic = `
resource "ecl_rca_user_v1" "user_1" {
    password = "dummy_passw@rd"
}
`

const testAccRCAV1UserUpdate = `
resource "ecl_rca_user_v1" "user_1" {
    password = "dummy_passw@rd_updated"
}
`
