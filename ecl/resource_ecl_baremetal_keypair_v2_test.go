package ecl

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"

	"github.com/nttcom/eclcloud/v2/ecl/baremetal/v2/keypairs"
)

func TestAccBaremetalV2Keypair_basic(t *testing.T) {
	var keypair keypairs.KeyPair

	fingerprintRe := regexp.MustCompile(`[a-f0-9:]+`)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBaremetalV2KeypairDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccBaremetalV2KeypairBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBaremetalV2KeypairExists("ecl_baremetal_keypair_v2.kp_1", &keypair),
					resource.TestMatchResourceAttr(
						"ecl_baremetal_keypair_v2.kp_1", "fingerprint", fingerprintRe),
				),
			},
		},
	})
}

func TestAccBaremetalV2Keypair_generatePrivate(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var keypair keypairs.KeyPair

	fingerprintRe := regexp.MustCompile(`[a-f0-9:]+`)
	privateKeyRe := regexp.MustCompile(`.*BEGIN RSA PRIVATE KEY.*`)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBaremetalV2KeypairDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccBaremetalV2KeypairGeneratePrivate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckBaremetalV2KeypairExists("ecl_baremetal_keypair_v2.kp_2", &keypair),
					resource.TestMatchResourceAttr(
						"ecl_baremetal_keypair_v2.kp_2", "fingerprint", fingerprintRe),
					resource.TestMatchResourceAttr(
						"ecl_baremetal_keypair_v2.kp_2", "private_key", privateKeyRe),
				),
			},
		},
	})
}

func testAccCheckBaremetalV2KeypairDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	baremetalClient, err := config.baremetalV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL baremetal client: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_baremetal_keypair_v2" {
			continue
		}

		_, err := keypairs.Get(baremetalClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("Keypair still exists")
		}
	}

	return nil
}

func testAccCheckBaremetalV2KeypairExists(n string, kp *keypairs.KeyPair) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		config := testAccProvider.Meta().(*Config)
		baremetalClient, err := config.baremetalV2Client(OS_REGION_NAME)
		if err != nil {
			return fmt.Errorf("Error creating ECL baremetal client: %s", err)
		}

		found, err := keypairs.Get(baremetalClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.Name != rs.Primary.ID {
			return fmt.Errorf("Keypair not found")
		}

		*kp = *found

		return nil
	}
}

const testAccBaremetalV2KeypairBasic = `
resource "ecl_baremetal_keypair_v2" "kp_1" {
  name = "kp1"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDAjpC1hwiOCCmKEWxJ4qzTTsJbKzndLo1BCz5PcwtUnflmU+gHJtWMZKpuEGVi29h0A/+ydKek1O18k10Ff+4tyFjiHDQAT9+OfgWf7+b1yK+qDip3X1C0UPMbwHlTfSGWLGZquwhvEFx9k3h/M+VtMvwR1lJ9LUyTAImnNjWG7TAIPmui30HvM2UiFEmqkr4ijq45MyX2+fLIePLRIFuu1p4whjHAQYufqyno3BS48icQb4p6iVEZPo4AE2o9oIyQvj2mx4dk5Y8CgSETOZTYDOR3rU2fZTRDRgPJDH9FWvQjF5tA0p3d9CoWWd2s6GKKbfoUIi8R/Db1BSPJwkqB jrp-hp-pc"
}
`

const testAccBaremetalV2KeypairGeneratePrivate = `
resource "ecl_baremetal_keypair_v2" "kp_2" {
  name = "kp2"
}
`
