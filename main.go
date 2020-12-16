package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/nttcom/terraform-provider-ecl/ecl"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ecl.Provider})
}
