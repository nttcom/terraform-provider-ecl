package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/ecl/sss/v1/users"
)

func TestAccSSSV1User_basic(t *testing.T) {
	var user users.User
	var loginID = fmt.Sprintf("ACCPTTEST-%s", acctest.RandString(15))
	var mailAddress = fmt.Sprintf("%s@example.com", loginID)

	loginIDUpdate := fmt.Sprintf("%supdate", loginID)
	mailAddressUpdate := fmt.Sprintf("%supdate@example.com", loginID)

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSSSV1UserDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccSSSV1UserBasic(loginID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSSSV1UserExists("ecl_sss_user_v1.user_1", &user),
					resource.TestCheckResourceAttr(
						"ecl_sss_user_v1.user_1", "login_id", loginID),
					resource.TestCheckResourceAttr(
						"ecl_sss_user_v1.user_1", "mail_address", mailAddress),
				),
			},
			resource.TestStep{
				Config: testAccSSSV1UserUpdate(loginID),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSSSV1UserExists("ecl_sss_user_v1.user_1", &user),
					resource.TestCheckResourceAttrPtr(
						"ecl_sss_user_v1.user_1", "login_id", &loginIDUpdate),
					resource.TestCheckResourceAttrPtr(
						"ecl_sss_user_v1.user_1", "mail_address", &mailAddressUpdate),
				),
			},
		},
	})
}

func testAccCheckSSSV1UserDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	client, err := config.sssV1Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL sss client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_sss_user_v1" {
			continue
		}

		_, err := users.Get(client, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("User still exists")
		}
	}

	return nil
}

func testAccCheckSSSV1UserExists(n string, tenant *users.User) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		client, err := config.sssV1Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating ECL sss client: %s", err)
		}

		found, err := users.Get(client, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.UserID != rs.Primary.ID {
			return fmt.Errorf("User not found")
		}

		*tenant = *found

		return nil
	}
}

func testAccSSSV1UserBasic(loginID string) string {
	return fmt.Sprintf(`
	resource "ecl_sss_user_v1" "user_1" {
	  login_id = "%s"
	  mail_address = "%s@example.com"
	  password = "Passw0rd"
	  notify_password = "false"
	}`, loginID, loginID)
}

func testAccSSSV1UserUpdate(loginID string) string {
	return fmt.Sprintf(`
	resource "ecl_sss_user_v1" "user_1" {
		login_id = "%supdate"
		mail_address = "%supdate@example.com"
		password = "Passw0rdupdate"
		notify_password = "false"
	}`, loginID, loginID)
}
