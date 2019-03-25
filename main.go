package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	"github.com/nstogner/terraform-provider-myfs-validator/validator"
	"github.com/nstogner/terraform-provider-myfs/provider"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return validator.New(provider.New())
		},
	})
}
