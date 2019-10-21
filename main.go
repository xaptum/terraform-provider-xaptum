package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/xaptum/terraform-provider-xaptum/xaptum"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: xaptum.Provider})
}
