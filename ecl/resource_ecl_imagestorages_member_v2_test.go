package ecl

import (
	"fmt"
	"testing"

	"github.com/nttcom/eclcloud/ecl/imagestorage/v2/members"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccImageStoragesV2Member_basic(t *testing.T) {
	var member members.Member

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheckImageMember(t)
			createTemporalImage(localFileForDataSourceTest)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckImageStoragesV2ImageDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccImageStoragesV2MemberBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImageStoragesV2MemberExists("ecl_imagestorages_member_v2.member_1", &member),
					resource.TestCheckResourceAttr(
						"ecl_imagestorages_member_v2.member_1", "member_id", OS_ACCEPTER_TENANT_ID),
					resource.TestCheckResourceAttr(
						"ecl_imagestorages_member_v2.member_1", "status", "pending"),
				),
			},
		},
	})
}

func testAccCheckImageStoragesV2MemberDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	imageClient, err := config.imageV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL Image: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_imagestorages_member_v2" {
			continue
		}

		imageId, memberId, err := imageStoragesMemberV2ParseID(rs.Primary.ID)
		if err != nil {
			return err
		}

		_, err = members.Get(imageClient, imageId, memberId).Extract()
		if err == nil {
			return fmt.Errorf("Image still exists")
		}
	}

	return nil
}

func testAccCheckImageStoragesV2MemberExists(n string, member *members.Member) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		imageClient, err := config.imageV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating ECL Image: %s", err)
		}

		imageId, memberId, err := imageStoragesMemberV2ParseID(rs.Primary.ID)
		if err != nil {
			return err
		}

		found, err := members.Get(imageClient, imageId, memberId).Extract()
		if err != nil {
			return err
		}

		if found.MemberID != memberId {
			return fmt.Errorf("Member not found")
		}

		*member = *found

		return nil
	}
}

var testAccImageStoragesV2MemberBasic = fmt.Sprintf(`
  resource "ecl_imagestorages_image_v2" "image_1" {
      name   = "Temp Terraform AccTest"
      local_file_path = "%s"
      container_format = "bare"
      disk_format = "qcow2"

      timeouts {
        create = "10m"
      }
  }
 
  resource "ecl_imagestorages_member_v2" "member_1" {
	  image_id = "${ecl_imagestorages_image_v2.image_1.id}"
	  member_id = "%s"
  }
  `, localFileForResourceTest,
	OS_ACCEPTER_TENANT_ID)
