provider "ecl" {
  force_sss_endpoint = "https://sss-jp5-ecl.api.ntt.com/api/v1.0/"
}

resource "ecl_sss_tenant_v2" "tenant_1" {
  tenant_name   = "example_tenant"
  description   = "new example tenant by terraform"
  tenant_region = "jp5"
  provisioner "local-exec" {
    command = "sleep 120 "
  }
}

module "network" {
  source    = "./modules/"
  tenant_id = ecl_sss_tenant_v2.tenant_1.tenant_id
}
