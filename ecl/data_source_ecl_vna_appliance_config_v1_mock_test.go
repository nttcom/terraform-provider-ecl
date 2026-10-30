package ecl

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccVNAV1ApplianceConfigDataSource_basic(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "virtual_network_appliances", fmt.Sprintf("/v1.0/virtual_network_appliances/%s/export-config", testMockVNAV1ApplianceConfigVSRXID), testMockVNAV1ApplianceConfigExportSet)
	mc.Register(t, "virtual_network_appliances", fmt.Sprintf("/v1.0/virtual_network_appliances/%s/export-config", testMockVNAV1ApplianceConfigVThunderID), testMockVNAV1ApplianceConfigExportA10JSON)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVNAV1ApplianceConfigDataSourceSet,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_vna_appliance_config_v1.config_1", "id", testMockVNAV1ApplianceConfigVSRXID+"/set"),
					resource.TestCheckResourceAttr("data.ecl_vna_appliance_config_v1.config_1", "virtual_network_appliance_id", testMockVNAV1ApplianceConfigVSRXID),
					resource.TestCheckResourceAttr("data.ecl_vna_appliance_config_v1.config_1", "format", "set"),
					resource.TestCheckResourceAttr("data.ecl_vna_appliance_config_v1.config_1", "data", testMockVNAV1ApplianceConfigSetData),
				),
			},
			{
				Config: testAccVNAV1ApplianceConfigDataSourceA10JSON,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_vna_appliance_config_v1.config_1", "id", testMockVNAV1ApplianceConfigVThunderID+"/a10-json"),
					resource.TestCheckResourceAttr("data.ecl_vna_appliance_config_v1.config_1", "format", "a10-json"),
					resource.TestCheckResourceAttr("data.ecl_vna_appliance_config_v1.config_1", "data", testMockVNAV1ApplianceConfigA10JSONData),
				),
			},
		},
	})
}

func TestMockedAccVNAV1ApplianceConfigDataSource_notActive(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)
	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "virtual_network_appliances", fmt.Sprintf("/v1.0/virtual_network_appliances/%s/export-config", testMockVNAV1ApplianceConfigVSRXID), testMockVNAV1ApplianceConfigExportNotActive)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccVNAV1ApplianceConfigDataSourceSet,
				ExpectError: regexp.MustCompile("Error exporting configuration of Virtual Network Appliance"),
			},
		},
	})
}

func TestMockedAccVNAV1ApplianceConfigDataSource_invalidFormat(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccVNAV1ApplianceConfigDataSourceInvalidFormat,
				ExpectError: regexp.MustCompile(`expected format to be one of \[set text a10-json\]`),
			},
		},
	})
}

const testMockVNAV1ApplianceConfigVSRXID = "45db3e66-31af-45a6-8ad2-d01521726141"
const testMockVNAV1ApplianceConfigVThunderID = "45db3e66-31af-45a6-8ad2-d01521726142"

// "set system host-name vsrx\n"
const testMockVNAV1ApplianceConfigSetData = "c2V0IHN5c3RlbSBob3N0LW5hbWUgdnNyeAo="

// "a10-url:/axapi/v3/hostname\n{\"hostname\":{\"value\":\"vthunder\"}}\n"
const testMockVNAV1ApplianceConfigA10JSONData = "YTEwLXVybDovYXhhcGkvdjMvaG9zdG5hbWUKeyJob3N0bmFtZSI6eyJ2YWx1ZSI6InZ0aHVuZGVyIn19Cg=="

var testAccVNAV1ApplianceConfigDataSourceSet = fmt.Sprintf(`
data "ecl_vna_appliance_config_v1" "config_1" {
	virtual_network_appliance_id = %q
	format = "set"
}
`,
	testMockVNAV1ApplianceConfigVSRXID,
)

var testAccVNAV1ApplianceConfigDataSourceA10JSON = fmt.Sprintf(`
data "ecl_vna_appliance_config_v1" "config_1" {
	virtual_network_appliance_id = %q
	format = "a10-json"
}
`,
	testMockVNAV1ApplianceConfigVThunderID,
)

var testAccVNAV1ApplianceConfigDataSourceInvalidFormat = fmt.Sprintf(`
data "ecl_vna_appliance_config_v1" "config_1" {
	virtual_network_appliance_id = %q
	format = "json"
}
`,
	testMockVNAV1ApplianceConfigVSRXID,
)

var testMockVNAV1ApplianceConfigExportSet = fmt.Sprintf(`
request:
    method: POST
    body: >
        {"format":"set"}
response:
    code: 200
    body: >
        {
          "configuration": {
            "format": "set",
            "data": %q
          }
        }
`,
	testMockVNAV1ApplianceConfigSetData,
)

var testMockVNAV1ApplianceConfigExportA10JSON = fmt.Sprintf(`
request:
    method: POST
    body: >
        {"format":"a10-json"}
response:
    code: 200
    body: >
        {
          "configuration": {
            "format": "a10-json",
            "data": %q
          }
        }
`,
	testMockVNAV1ApplianceConfigA10JSONData,
)

var testMockVNAV1ApplianceConfigExportNotActive = `
request:
    method: POST
    body: >
        {"format":"set"}
response:
    code: 400
    body: >
        {
          "message": "Specified appliance is not in the active state."
        }
`
