package main

import (
        "github.com/hashicorp/terraform/plugin"
        "github.com/xaptum/terraform-provider-xaptum/enf"
)

func main() {
        plugin.Serve(&plugin.ServeOpts{
                ProviderFunc: enf.Provider})
}
