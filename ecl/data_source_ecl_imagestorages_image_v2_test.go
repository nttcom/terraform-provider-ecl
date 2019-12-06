package ecl

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const localFileForDataSourceTest = "/tmp/tempfile.img"

func TestAccImageStoragesV2ImageDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			createTemporalImage(localFileForDataSourceTest)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccImageStoragesV2ImageDataSourceCirros,
			},
			resource.TestStep{
				Config: testAccImageStoragesV2ImageDataSourceBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImagesV2DataSourceID("data.ecl_imagestorages_image_v2.image_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_imagestorages_image_v2.image_1", "name", "Temp-tf_1"),
					resource.TestCheckResourceAttr(
						"data.ecl_imagestorages_image_v2.image_1", "container_format", "bare"),
					resource.TestCheckResourceAttr(
						"data.ecl_imagestorages_image_v2.image_1", "disk_format", "qcow2"),
					resource.TestCheckResourceAttr(
						"data.ecl_imagestorages_image_v2.image_1", "min_disk_gb", "0"),
					resource.TestCheckResourceAttr(
						"data.ecl_imagestorages_image_v2.image_1", "min_ram_mb", "0"),
					resource.TestCheckResourceAttr(
						"data.ecl_imagestorages_image_v2.image_1", "protected", "false"),
					resource.TestCheckResourceAttr(
						"data.ecl_imagestorages_image_v2.image_1", "visibility", "private"),
				),
			},
		},
	})
}

func TestAccImagesV2ImageDataSourceTestQueries(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccImageStoragesV2ImageDataSourceCirros,
			},
			resource.TestStep{
				Config: testAccImageStoragesV2ImageDataSourceQueryTag,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImagesV2DataSourceID("data.ecl_imagestorages_image_v2.image_1"),
				),
			},
			resource.TestStep{
				Config: testAccImageStoragesV2ImageDataSourceQuerySizeMin,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImagesV2DataSourceID("data.ecl_imagestorages_image_v2.image_1"),
				),
			},
			resource.TestStep{
				Config: testAccImageStoragesV2ImageDataSourceQuerySizeMax,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImagesV2DataSourceID("data.ecl_imagestorages_image_v2.image_1"),
				),
			},
			resource.TestStep{
				Config: testAccImageStoragesV2ImageDataSourceProperty,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImagesV2DataSourceID("data.ecl_imagestorages_image_v2.image_1"),
				),
			},
			resource.TestStep{
				Config: testAccImageStoragesV2ImageDataSourceCirros,
			},
		},
	})
}

func testAccCheckImagesV2DataSourceID(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find image data source: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Image data source ID not set")
		}

		return nil
	}
}

// Standard CirrOS image
var testAccImageStoragesV2ImageDataSourceCirros = fmt.Sprintf(`
resource "ecl_imagestorages_image_v2" "image_1" {
  name = "Temp-tf_1"
  container_format = "bare"
  disk_format = "qcow2"
  local_file_path = "%s"
  tags = ["cirros-tf_1"]
  properties = {
    foo = "bar"
    bar = "foo"
  }
}

`, localFileForDataSourceTest)

var testAccImageStoragesV2ImageDataSourceBasic = fmt.Sprintf(`
%s

data "ecl_imagestorages_image_v2" "image_1" {
	most_recent = true
	name = "${ecl_imagestorages_image_v2.image_1.name}"
}
`, testAccImageStoragesV2ImageDataSourceCirros)

var testAccImageStoragesV2ImageDataSourceQueryTag = fmt.Sprintf(`
%s

data "ecl_imagestorages_image_v2" "image_1" {
	most_recent = true
	visibility = "private"
	tag = "cirros-tf_1"
}
`, testAccImageStoragesV2ImageDataSourceCirros)

var testAccImageStoragesV2ImageDataSourceQuerySizeMin = fmt.Sprintf(`
%s

data "ecl_imagestorages_image_v2" "image_1" {
	most_recent = true
	visibility = "private"
	size_min = "13000000"
}
`, testAccImageStoragesV2ImageDataSourceCirros)

var testAccImageStoragesV2ImageDataSourceQuerySizeMax = fmt.Sprintf(`
%s

data "ecl_imagestorages_image_v2" "image_1" {
	most_recent = true
	visibility = "private"
	size_max = "23000000"
}
`, testAccImageStoragesV2ImageDataSourceCirros)

var testAccImageStoragesV2ImageDataSourceProperty = fmt.Sprintf(`
%s

data "ecl_imagestorages_image_v2" "image_1" {
  properties = {
    foo = "bar"
    bar = "foo"
  }
}
`, testAccImageStoragesV2ImageDataSourceCirros)
