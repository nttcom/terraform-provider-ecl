package ecl

import (
	"fmt"
	"testing"

	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/nttcom/eclcloud/ecl/imagestorage/v2/images"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const localFileForResourceTest = "/tmp/tempfile.img"

func TestAccImageStoragesV2Image_basic(t *testing.T) {
	var image images.Image

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			createTemporalImage(localFileForDataSourceTest)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckImageStoragesV2ImageDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccImageStoragesV2ImageBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImageStoragesV2ImageExists("ecl_imagestorages_image_v2.image_1", &image),
					resource.TestCheckResourceAttr(
						"ecl_imagestorages_image_v2.image_1", "container_format", "bare"),
					resource.TestCheckResourceAttr(
						"ecl_imagestorages_image_v2.image_1", "name", "Temp Terraform AccTest"),
					resource.TestCheckResourceAttr(
						"ecl_imagestorages_image_v2.image_1", "disk_format", "qcow2"),
					resource.TestCheckResourceAttr(
						"ecl_imagestorages_image_v2.image_1", "schema", "/v2/schemas/image"),
				),
			},
		},
	})
}

func TestAccImageStoragesV2Image_name(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var image images.Image

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			createTemporalImage(localFileForDataSourceTest)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckImageStoragesV2ImageDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccImageStoragesV2ImageName1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImageStoragesV2ImageExists("ecl_imagestorages_image_v2.image_1", &image),
					resource.TestCheckResourceAttr(
						"ecl_imagestorages_image_v2.image_1", "name", "Temp Terraform AccTest"),
				),
			},
			resource.TestStep{
				Config: testAccImageStoragesV2ImageName2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImageStoragesV2ImageExists("ecl_imagestorages_image_v2.image_1", &image),
					resource.TestCheckResourceAttr(
						"ecl_imagestorages_image_v2.image_1", "name", stringMaxLength),
				),
			},
			resource.TestStep{
				Config: testAccImageStoragesV2ImageName3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImageStoragesV2ImageExists("ecl_imagestorages_image_v2.image_1", &image),
					resource.TestCheckResourceAttr(
						"ecl_imagestorages_image_v2.image_1", "name", ""),
				),
			},
		},
	})
}

func TestAccImageStoragesV2Image_tags(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var image images.Image

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			createTemporalImage(localFileForDataSourceTest)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckImageStoragesV2ImageDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccImageStoragesV2ImageTags1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImageStoragesV2ImageExists("ecl_imagestorages_image_v2.image_1", &image),
					testAccCheckImageStoragesV2ImageHasTag("ecl_imagestorages_image_v2.image_1", "foo"),
					testAccCheckImageStoragesV2ImageHasTag("ecl_imagestorages_image_v2.image_1", "bar"),
					testAccCheckImageStoragesV2ImageTagCount("ecl_imagestorages_image_v2.image_1", 2),
				),
			},
			resource.TestStep{
				Config: testAccImageStoragesV2ImageTags2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImageStoragesV2ImageExists("ecl_imagestorages_image_v2.image_1", &image),
					testAccCheckImageStoragesV2ImageHasTag("ecl_imagestorages_image_v2.image_1", "foo"),
					testAccCheckImageStoragesV2ImageHasTag("ecl_imagestorages_image_v2.image_1", "bar"),
					testAccCheckImageStoragesV2ImageHasTag("ecl_imagestorages_image_v2.image_1", "baz"),
					testAccCheckImageStoragesV2ImageTagCount("ecl_imagestorages_image_v2.image_1", 3),
				),
			},
			resource.TestStep{
				Config: testAccImageStoragesV2ImageTags3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImageStoragesV2ImageExists("ecl_imagestorages_image_v2.image_1", &image),
					testAccCheckImageStoragesV2ImageHasTag("ecl_imagestorages_image_v2.image_1", "foo"),
					testAccCheckImageStoragesV2ImageHasTag("ecl_imagestorages_image_v2.image_1", "baz"),
					testAccCheckImageStoragesV2ImageTagCount("ecl_imagestorages_image_v2.image_1", 2),
				),
			},
			resource.TestStep{
				Config: testAccImageStoragesV2ImageTags4,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImageStoragesV2ImageExists("ecl_imagestorages_image_v2.image_1", &image),
					testAccCheckImageStoragesV2ImageTagCount("ecl_imagestorages_image_v2.image_1", 0),
				),
			},
		},
	})
}

func TestAccImageStoragesV2Image_licenseSwitch(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var image images.Image

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
			createTemporalImage(localFileForDataSourceTest)
		},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckImageStoragesV2ImageDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccImageStoragesV2ImageLicenseSwitch,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImageStoragesV2ImageExists("ecl_imagestorages_image_v2.image_1", &image),
					resource.TestCheckResourceAttr(
						"ecl_imagestorages_image_v2.image_1", "name", "Temp Terraform AccTest"),
					resource.TestCheckResourceAttr(
						"ecl_imagestorages_image_v2.image_1", "container_format", "bare"),
					resource.TestCheckResourceAttr(
						"ecl_imagestorages_image_v2.image_1", "disk_format", "qcow2"),
					resource.TestCheckResourceAttr(
						"ecl_imagestorages_image_v2.image_1", "license_switch", "WindowsServer_2012R2_Standard_64bit_ComLicense"),
					resource.TestCheckResourceAttr(
						"ecl_imagestorages_image_v2.image_1", "schema", "/v2/schemas/image"),
				),
			},
		},
	})
}

func TestAccImageStoragesV2Image_properties(t *testing.T) {
	if testing.Short() {
		t.Skip("skip this test in short mode")
	}

	var image1 images.Image
	var image2 images.Image
	var image3 images.Image
	var image4 images.Image
	var image5 images.Image

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckImageStoragesV2ImageDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccImageStoragesV2ImageBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImageStoragesV2ImageExists("ecl_imagestorages_image_v2.image_1", &image1),
				),
			},
			resource.TestStep{
				Config: testAccImageStoragesV2ImageProperties1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImageStoragesV2ImageExists("ecl_imagestorages_image_v2.image_1", &image2),
					resource.TestCheckResourceAttr(
						"ecl_imagestorages_image_v2.image_1", "properties.foo", "bar"),
					resource.TestCheckResourceAttr(
						"ecl_imagestorages_image_v2.image_1", "properties.bar", "foo"),
				),
			},
			resource.TestStep{
				Config: testAccImageStoragesV2ImageProperties2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImageStoragesV2ImageExists("ecl_imagestorages_image_v2.image_1", &image3),
					resource.TestCheckResourceAttr(
						"ecl_imagestorages_image_v2.image_1", "properties.foo", "bar"),
				),
			},
			resource.TestStep{
				Config: testAccImageStoragesV2ImageProperties3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImageStoragesV2ImageExists("ecl_imagestorages_image_v2.image_1", &image4),
					resource.TestCheckResourceAttr(
						"ecl_imagestorages_image_v2.image_1", "properties.foo", "baz"),
				),
			},
			resource.TestStep{
				Config: testAccImageStoragesV2ImageProperties4,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckImageStoragesV2ImageExists("ecl_imagestorages_image_v2.image_1", &image5),
					resource.TestCheckResourceAttr(
						"ecl_imagestorages_image_v2.image_1", "properties.foo", "baz"),
					resource.TestCheckResourceAttr(
						"ecl_imagestorages_image_v2.image_1", "properties.bar", "foo"),
				),
			},
		},
	})
}

func TestValidateLicenseSwitch(t *testing.T) {
	validValues := []string{
		"Red_Hat_Enterprise_Linux_6_64bit_BYOL",
		"WindowsServer_2012R2_Standard_64bit_ComLicense",
		"WindowsServer_2012_Standard_64bit_ComLicense",
		"WindowsServer_2008R2_Enterprise_64bit_ComLicense",
		"WindowsServer_2008R2_Standard_64bit_ComLicense",
		"WindowsServer_2008_Enterprise_64bit_ComLicense",
		"WindowsServer_2008_Standard_64bit_ComLicense",
	}
	for _, v := range validValues {
		_, err := validateLicenseSwitch(v, "license_switch")
		if err != nil {
			t.Fatalf("Test failed while checking value = %s as %#v", v, err)
		}
	}

	invalidValues := []string{
		"Ped_Hat_Enterprise_Linux_6_64bit_BYOL",
		"Red_Hat_Enterprise_Linux_6_64bit_BYOS",
		"Red_Hat_Enterprise_Linux_X_64bit_BYOL",
		"VindowsServer_2012R2_Standard_64bit_ComLicense",
		"WindowsServer_2012R2_Enterprise_64bit_ComLicensu",
		"WindowsServer_2012R2_Stundard_64bit_ComLicense",
	}
	for _, iv := range invalidValues {
		_, err := validateLicenseSwitch(iv, "license_switch")
		if err == nil {
			t.Fatalf("Test failed while checking invalid value = %s as %#v", iv, err)
		}
	}

	return
}

func testAccCheckImageStoragesV2ImageDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(*Config)
	imageClient, err := config.imageV2Client(OS_REGION_NAME)
	if err != nil {
		return fmt.Errorf("Error creating ECL Image: %s", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ecl_imagestorages_image_v2" {
			continue
		}

		_, err := images.Get(imageClient, rs.Primary.ID).Extract()
		if err == nil {
			return fmt.Errorf("Image still exists")
		}
	}

	return nil
}

func testAccCheckImageStoragesV2ImageExists(n string, image *images.Image) resource.TestCheckFunc {
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

		found, err := images.Get(imageClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("Image not found")
		}

		*image = *found

		return nil
	}
}

func testAccCheckImageStoragesV2ImageHasTag(n, tag string) resource.TestCheckFunc {
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

		found, err := images.Get(imageClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("Image not found")
		}

		for _, v := range found.Tags {
			if tag == v {
				return nil
			}
		}

		return fmt.Errorf("Tag not found: %s", tag)
	}
}

func testAccCheckImageStoragesV2ImageTagCount(n string, expected int) resource.TestCheckFunc {
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

		found, err := images.Get(imageClient, rs.Primary.ID).Extract()
		if err != nil {
			return err
		}

		if found.ID != rs.Primary.ID {
			return fmt.Errorf("Image not found")
		}

		if len(found.Tags) != expected {
			return fmt.Errorf("Expecting %d tags, found %d", expected, len(found.Tags))
		}

		return nil
	}
}

func createTemporalImage(tempImageSavePath string) {
	_, err := os.Stat(tempImageSavePath)
	if err == nil {
		return
	}
	content := []byte("DummyContent\n")
	fmt.Println("Creating file for AccTest...")
	ioutil.WriteFile(tempImageSavePath, content, os.ModePerm)
	if err != nil {
		fmt.Println("Done")
		return
	}
}

func downloadImage(tempImageSavePath, imageSourceURL string) {
	_, err := os.Stat(tempImageSavePath)
	if err == nil {
		return
	}

	img, _ := os.Create(tempImageSavePath)
	defer img.Close()

	fmt.Println("Downloading file for AccTest...")
	resp, err := http.Get(imageSourceURL)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	io.Copy(img, resp.Body)

	return
}

var testAccImageStoragesV2ImageBasic = fmt.Sprintf(`
  resource "ecl_imagestorages_image_v2" "image_1" {
      name = "Temp Terraform AccTest"
      local_file_path = "%s"
      container_format = "bare"
      disk_format = "qcow2"

      timeouts {
        create = "10m"
      }
  }`, localFileForResourceTest)

var testAccImageStoragesV2ImageName1 = fmt.Sprintf(`
  resource "ecl_imagestorages_image_v2" "image_1" {
      name = "Temp Terraform AccTest"
      local_file_path = "%s"
      container_format = "bare"
      disk_format = "qcow2"
  }`, localFileForResourceTest)

var testAccImageStoragesV2ImageName2 = fmt.Sprintf(`
  resource "ecl_imagestorages_image_v2" "image_1" {
      name = "%s"
      local_file_path = "%s"
      container_format = "bare"
      disk_format = "qcow2"
  }`, stringMaxLength,
	localFileForResourceTest)

var testAccImageStoragesV2ImageName3 = fmt.Sprintf(`
  resource "ecl_imagestorages_image_v2" "image_1" {
      name = ""
      local_file_path = "%s"
      container_format = "bare"
      disk_format = "qcow2"
  }`, localFileForResourceTest)

var testAccImageStoragesV2ImageTags1 = fmt.Sprintf(`
  resource "ecl_imagestorages_image_v2" "image_1" {
      name = "Temp Terraform AccTest"
      local_file_path = "%s"
      container_format = "bare"
      disk_format = "qcow2"
      tags = ["foo","bar"]
  }`, localFileForResourceTest)

var testAccImageStoragesV2ImageTags2 = fmt.Sprintf(`
  resource "ecl_imagestorages_image_v2" "image_1" {
      name = "Temp Terraform AccTest"
      local_file_path = "%s"
      container_format = "bare"
      disk_format = "qcow2"
      tags = ["foo","bar","baz"]
  }`, localFileForResourceTest)

var testAccImageStoragesV2ImageTags3 = fmt.Sprintf(`
  resource "ecl_imagestorages_image_v2" "image_1" {
      name = "Temp Terraform AccTest"
      local_file_path = "%s"
      container_format = "bare"
      disk_format = "qcow2"
      tags = ["foo","baz"]
  }`, localFileForResourceTest)

var testAccImageStoragesV2ImageTags4 = fmt.Sprintf(`
  resource "ecl_imagestorages_image_v2" "image_1" {
      name = "Temp Terraform AccTest"
      local_file_path = "%s"
      container_format = "bare"
      disk_format = "qcow2"
      tags = []
  }`, localFileForResourceTest)

var testAccImageStoragesV2ImageVisibility1 = fmt.Sprintf(`
  resource "ecl_imagestorages_image_v2" "image_1" {
      name = "Temp Terraform AccTest"
      local_file_path = "%s"
      container_format = "bare"
      disk_format = "qcow2"
      visibility = "private"
  }`, localFileForResourceTest)

var testAccImageStoragesV2ImageVisibility2 = fmt.Sprintf(`
  resource "ecl_imagestorages_image_v2" "image_1" {
      name = "Temp Terraform AccTest"
      local_file_path = "%s"
      container_format = "bare"
      disk_format = "qcow2"
      visibility = "public"
  }`, localFileForResourceTest)

var testAccImageStoragesV2ImageProperties1 = fmt.Sprintf(`
  resource "ecl_imagestorages_image_v2" "image_1" {
      name = "Temp Terraform AccTest"
      local_file_path = "%s"
      container_format = "bare"
      disk_format = "qcow2"

      properties = {
        foo = "bar"
        bar = "foo"
      }
  }`, localFileForResourceTest)

var testAccImageStoragesV2ImageProperties2 = fmt.Sprintf(`
  resource "ecl_imagestorages_image_v2" "image_1" {
      name = "Temp Terraform AccTest"
      local_file_path = "%s"
      container_format = "bare"
      disk_format = "qcow2"

      properties = {
        foo = "bar"
      }
  }`, localFileForResourceTest)

var testAccImageStoragesV2ImageProperties3 = fmt.Sprintf(`
  resource "ecl_imagestorages_image_v2" "image_1" {
      name = "Temp Terraform AccTest"
      local_file_path = "%s"
      container_format = "bare"
      disk_format = "qcow2"

      properties = {
        foo = "baz"
      }
  }`, localFileForResourceTest)

var testAccImageStoragesV2ImageProperties4 = fmt.Sprintf(`
  resource "ecl_imagestorages_image_v2" "image_1" {
      name = "Temp Terraform AccTest"
      local_file_path = "%s"
      container_format = "bare"
      disk_format = "qcow2"

      properties = {
        foo = "baz"
        bar = "foo"
      }
  }`, localFileForResourceTest)

var testAccImageStoragesV2ImageLicenseSwitch = fmt.Sprintf(`
  resource "ecl_imagestorages_image_v2" "image_1" {
    name = "Temp Terraform AccTest"
    local_file_path = "%s"
    container_format = "bare"
    disk_format = "qcow2"
    license_switch = "WindowsServer_2012R2_Standard_64bit_ComLicense"
  }`, localFileForResourceTest)
