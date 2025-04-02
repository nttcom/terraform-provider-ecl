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
	mc.Register(t, "certificates", "/v1.0/certificates", testMockMLBV1CertificatesListDescriptionQuery)
	mc.Register(t, "certificates", "/v1.0/certificates", testMockMLBV1CertificatesListTenantIDQuery)

	mc.StartServer(t)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccMLBV1CertificateDataSourceQueryName,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "name", "certificate"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "ca_cert.status", "UPLOADED"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "ca_cert.info", `{"fingerprint":"db:b1:49:84:f6:2e:ec:c9:41:fc:a1:30:26:12:2c:37:4d:bb:7a:bd","issuer":{"C":"JP","CN":"example.com","L":"Chiyoda-ku","O":"NTT Communications Corporation","ST":"Tokyo"},"key_algorithm":"RSA-4096","not_after":"2024-12-10 06:20:54","not_before":"2023-11-09 06:20:55","serial":"e7:61:4a:49:85:aa:7c:f2","subject":{"C":"JP","CN":"example.com","L":"Chiyoda-ku","O":"NTT Communications Corporation","ST":"Tokyo"}}`),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "ssl_cert.status", "UPLOADED"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "ssl_cert.info", `{"fingerprint":"46:06:c5:ed:f0:e6:9f:c5:e3:bd:06:63:54:88:9f:3d:a7:c5:42:b2","issuer":{"C":"JP","CN":"example.com","L":"Chiyoda-ku","O":"NTT Communications Corporation","ST":"Tokyo"},"key_algorithm":"RSA-4096","not_after":"2024-12-10 06:20:54","not_before":"2023-11-09 06:20:55","serial":"d3:11:fe:4d:a3:71:4e:13","subject":{"C":"JP","CN":"example.com","L":"Chiyoda-ku","O":"NTT Communications Corporation","ST":"Tokyo"}}`),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "ssl_key.status", "UPLOADED"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "ssl_key.info", `{"key_algorithm":"RSA-4096","passphrase":true}`),
				),
			},
			{
				Config: testAccMLBV1CertificateDataSourceQueryDescription,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "name", "certificate"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "ca_cert.status", "UPLOADED"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "ca_cert.info", `{"fingerprint":"db:b1:49:84:f6:2e:ec:c9:41:fc:a1:30:26:12:2c:37:4d:bb:7a:bd","issuer":{"C":"JP","CN":"example.com","L":"Chiyoda-ku","O":"NTT Communications Corporation","ST":"Tokyo"},"key_algorithm":"RSA-4096","not_after":"2024-12-10 06:20:54","not_before":"2023-11-09 06:20:55","serial":"e7:61:4a:49:85:aa:7c:f2","subject":{"C":"JP","CN":"example.com","L":"Chiyoda-ku","O":"NTT Communications Corporation","ST":"Tokyo"}}`),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "ssl_cert.status", "UPLOADED"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "ssl_cert.info", `{"fingerprint":"46:06:c5:ed:f0:e6:9f:c5:e3:bd:06:63:54:88:9f:3d:a7:c5:42:b2","issuer":{"C":"JP","CN":"example.com","L":"Chiyoda-ku","O":"NTT Communications Corporation","ST":"Tokyo"},"key_algorithm":"RSA-4096","not_after":"2024-12-10 06:20:54","not_before":"2023-11-09 06:20:55","serial":"d3:11:fe:4d:a3:71:4e:13","subject":{"C":"JP","CN":"example.com","L":"Chiyoda-ku","O":"NTT Communications Corporation","ST":"Tokyo"}}`),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "ssl_key.status", "UPLOADED"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "ssl_key.info", `{"key_algorithm":"RSA-4096","passphrase":true}`),
				),
			},
			{
				Config: testAccMLBV1CertificateDataSourceQueryTenantID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "id", "497f6eca-6276-4993-bfeb-53cbbbba6f08"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "name", "certificate"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "description", "description"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "tags.key", "value"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "tenant_id", "34f5c98ef430457ba81292637d0c6fd0"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "ca_cert.status", "UPLOADED"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "ca_cert.info", `{"fingerprint":"db:b1:49:84:f6:2e:ec:c9:41:fc:a1:30:26:12:2c:37:4d:bb:7a:bd","issuer":{"C":"JP","CN":"example.com","L":"Chiyoda-ku","O":"NTT Communications Corporation","ST":"Tokyo"},"key_algorithm":"RSA-4096","not_after":"2024-12-10 06:20:54","not_before":"2023-11-09 06:20:55","serial":"e7:61:4a:49:85:aa:7c:f2","subject":{"C":"JP","CN":"example.com","L":"Chiyoda-ku","O":"NTT Communications Corporation","ST":"Tokyo"}}`),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "ssl_cert.status", "UPLOADED"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "ssl_cert.info", `{"fingerprint":"46:06:c5:ed:f0:e6:9f:c5:e3:bd:06:63:54:88:9f:3d:a7:c5:42:b2","issuer":{"C":"JP","CN":"example.com","L":"Chiyoda-ku","O":"NTT Communications Corporation","ST":"Tokyo"},"key_algorithm":"RSA-4096","not_after":"2024-12-10 06:20:54","not_before":"2023-11-09 06:20:55","serial":"d3:11:fe:4d:a3:71:4e:13","subject":{"C":"JP","CN":"example.com","L":"Chiyoda-ku","O":"NTT Communications Corporation","ST":"Tokyo"}}`),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "ssl_key.status", "UPLOADED"),
					resource.TestCheckResourceAttr("data.ecl_mlb_certificate_v1.certificate_1", "ssl_key.info", `{"key_algorithm":"RSA-4096","passphrase":true}`),
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
            "status": "UPLOADED",
            "info": {
              "issuer": {
                "C": "JP",
                "ST": "Tokyo",
                "L": "Chiyoda-ku",
                "O": "NTT Communications Corporation",
                "CN": "example.com"
              },
              "subject": {
                "C": "JP",
                "ST": "Tokyo",
                "L": "Chiyoda-ku",
                "O": "NTT Communications Corporation",
                "CN": "example.com"
              },
              "not_before": "2023-11-09 06:20:55",
              "not_after": "2024-12-10 06:20:54",
              "key_algorithm": "RSA-4096",
              "serial": "e7:61:4a:49:85:aa:7c:f2",
              "fingerprint": "db:b1:49:84:f6:2e:ec:c9:41:fc:a1:30:26:12:2c:37:4d:bb:7a:bd"
            }
          },
          "ssl_cert": {
            "status": "UPLOADED",
            "info": {
              "issuer": {
                "C": "JP",
                "ST": "Tokyo",
                "L": "Chiyoda-ku",
                "O": "NTT Communications Corporation",
                "CN": "example.com"
              },
              "subject": {
                "C": "JP",
                "ST": "Tokyo",
                "L": "Chiyoda-ku",
                "O": "NTT Communications Corporation",
                "CN": "example.com"
              },
              "not_before": "2023-11-09 06:20:55",
              "not_after": "2024-12-10 06:20:54",
              "key_algorithm": "RSA-4096",
              "serial": "d3:11:fe:4d:a3:71:4e:13",
              "fingerprint": "46:06:c5:ed:f0:e6:9f:c5:e3:bd:06:63:54:88:9f:3d:a7:c5:42:b2"
            }
          },
          "ssl_key": {
            "status": "UPLOADED",
            "info": {
              "key_algorithm": "RSA-4096",
              "passphrase": true
            }
          }
        }
      ]
    }
`)

var testAccMLBV1CertificateDataSourceQueryDescription = fmt.Sprintf(`
data "ecl_mlb_certificate_v1" "certificate_1" {
  description = "description"
}
`)

var testMockMLBV1CertificatesListDescriptionQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    description:
      - description
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
            "status": "UPLOADED",
            "info": {
              "issuer": {
                "C": "JP",
                "ST": "Tokyo",
                "L": "Chiyoda-ku",
                "O": "NTT Communications Corporation",
                "CN": "example.com"
              },
              "subject": {
                "C": "JP",
                "ST": "Tokyo",
                "L": "Chiyoda-ku",
                "O": "NTT Communications Corporation",
                "CN": "example.com"
              },
              "not_before": "2023-11-09 06:20:55",
              "not_after": "2024-12-10 06:20:54",
              "key_algorithm": "RSA-4096",
              "serial": "e7:61:4a:49:85:aa:7c:f2",
              "fingerprint": "db:b1:49:84:f6:2e:ec:c9:41:fc:a1:30:26:12:2c:37:4d:bb:7a:bd"
            }
          },
          "ssl_cert": {
            "status": "UPLOADED",
            "info": {
              "issuer": {
                "C": "JP",
                "ST": "Tokyo",
                "L": "Chiyoda-ku",
                "O": "NTT Communications Corporation",
                "CN": "example.com"
              },
              "subject": {
                "C": "JP",
                "ST": "Tokyo",
                "L": "Chiyoda-ku",
                "O": "NTT Communications Corporation",
                "CN": "example.com"
              },
              "not_before": "2023-11-09 06:20:55",
              "not_after": "2024-12-10 06:20:54",
              "key_algorithm": "RSA-4096",
              "serial": "d3:11:fe:4d:a3:71:4e:13",
              "fingerprint": "46:06:c5:ed:f0:e6:9f:c5:e3:bd:06:63:54:88:9f:3d:a7:c5:42:b2"
            }
          },
          "ssl_key": {
            "status": "UPLOADED",
            "info": {
              "key_algorithm": "RSA-4096",
              "passphrase": true
            }
          }
        }
      ]
    }
`)

var testAccMLBV1CertificateDataSourceQueryTenantID = fmt.Sprintf(`
data "ecl_mlb_certificate_v1" "certificate_1" {
  tenant_id = "34f5c98ef430457ba81292637d0c6fd0"
}
`)

var testMockMLBV1CertificatesListTenantIDQuery = fmt.Sprintf(`
request:
  method: GET
  query:
    tenant_id:
      - 34f5c98ef430457ba81292637d0c6fd0
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
            "status": "UPLOADED",
            "info": {
              "issuer": {
                "C": "JP",
                "ST": "Tokyo",
                "L": "Chiyoda-ku",
                "O": "NTT Communications Corporation",
                "CN": "example.com"
              },
              "subject": {
                "C": "JP",
                "ST": "Tokyo",
                "L": "Chiyoda-ku",
                "O": "NTT Communications Corporation",
                "CN": "example.com"
              },
              "not_before": "2023-11-09 06:20:55",
              "not_after": "2024-12-10 06:20:54",
              "key_algorithm": "RSA-4096",
              "serial": "e7:61:4a:49:85:aa:7c:f2",
              "fingerprint": "db:b1:49:84:f6:2e:ec:c9:41:fc:a1:30:26:12:2c:37:4d:bb:7a:bd"
            }
          },
          "ssl_cert": {
            "status": "UPLOADED",
            "info": {
              "issuer": {
                "C": "JP",
                "ST": "Tokyo",
                "L": "Chiyoda-ku",
                "O": "NTT Communications Corporation",
                "CN": "example.com"
              },
              "subject": {
                "C": "JP",
                "ST": "Tokyo",
                "L": "Chiyoda-ku",
                "O": "NTT Communications Corporation",
                "CN": "example.com"
              },
              "not_before": "2023-11-09 06:20:55",
              "not_after": "2024-12-10 06:20:54",
              "key_algorithm": "RSA-4096",
              "serial": "d3:11:fe:4d:a3:71:4e:13",
              "fingerprint": "46:06:c5:ed:f0:e6:9f:c5:e3:bd:06:63:54:88:9f:3d:a7:c5:42:b2"
            }
          },
          "ssl_key": {
            "status": "UPLOADED",
            "info": {
              "key_algorithm": "RSA-4096",
              "passphrase": true
            }
          }
        }
      ]
    }
`)
