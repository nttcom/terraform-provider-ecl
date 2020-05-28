package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/terraform-providers/terraform-provider-ecl/ecl"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ecl.Provider})
}
