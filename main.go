package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform-provider-openstack_legacy/openstack"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: openstack.Provider})
}
