package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccVNAV1ApplianceConfigDataSource_basic(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckVNA(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccVNAV1ApplianceBasic,
			},
			resource.TestStep{
				Config: testAccVNAV1ApplianceConfigDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVNAV1ApplianceDataSourceID("data.ecl_vna_appliance_config_v1.config_1"),
					resource.TestCheckResourceAttr("data.ecl_vna_appliance_config_v1.config_1", "format", "set"),
					resource.TestCheckResourceAttrSet("data.ecl_vna_appliance_config_v1.config_1", "data"),
				),
			},
		},
	})
}

var testAccVNAV1ApplianceConfigDataSourceBasic = fmt.Sprintf(`
%s

data "ecl_vna_appliance_config_v1" "config_1" {
	virtual_network_appliance_id = "${ecl_vna_appliance_v1.appliance_1.id}"
	format = "set"
}
`,
	testAccVNAV1ApplianceBasic,
)
