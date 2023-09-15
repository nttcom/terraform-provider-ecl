/*
Generated by https://github.com/tamac-io/openapi-to-terraform-rb
*/
package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"

	"github.com/nttcom/terraform-provider-ecl/ecl/testhelper/mock"
)

func TestMockedAccMLBV1CertificateDataSource(t *testing.T) {
	mc := mock.NewMockController()
	defer mc.TerminateMockControllerSafety()

	postKeystone := fmt.Sprintf(fakeKeystonePostTmpl, mc.Endpoint(), OS_REGION_NAME)

	mc.Register(t, "keystone", "/v3/auth/tokens", postKeystone)
	mc.Register(t, "certificates", "/v1.0/certificates", testMockMLBV1CertificatesListNameQuery)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1CertificateDataSourceQueryName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ecl_mlb_certificate_v1.certificate_1", "id"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "name", "certificate"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "ca_cert.0.status", "NOT_UPLOADED"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "ssl_cert.0.status", "NOT_UPLOADED"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "ssl_key.0.status", "NOT_UPLOADED"),
				),
			},
		},
	})
}

var testAccMLBV1CertificateDataSourceQueryName = fmt.Sprintf(`
data "ecl_mlb_certificate_v1" "certificate_1" {
  name = "certificate"
}
`)

var testMockMLBV1CertificatesListNameQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    name:
      - certificate
response:
  code: 200
  body: >
    {
      "certificates": [
        {
          "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
          "name": "certificate",
          "description": "description",
          "tags": {
            "key": "value"
          },
          "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
          "ca_cert": {
            "status": "NOT_UPLOADED"
          },
          "ssl_cert": {
            "status": "NOT_UPLOADED"
          },
          "ssl_key": {
            "status": "NOT_UPLOADED"
          }
        }
      ]
    }
`)