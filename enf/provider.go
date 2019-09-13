package enf

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider
func Provider() terraform.ResourceProvider {

	// The actual provider
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ENF_USERNAME", nil),
				Description: "Username for authenticating with dev.xaptum.io",
			},

			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ENF_PASSWORD", nil),
				Description: "Password for authenticating with dev.xaptum.io",
			},

			"domain_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ENF_DOMAIN_URL", nil),
				Description: "Base URL for API calls",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"enf_firewall_rule": resourceEnfFirewallRule(),
			"enf_network":       resourceEnfNetwork(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Username:  d.Get("username").(string),
		Password:  d.Get("password").(string),
		DomainURL: d.Get("domain_url").(string),
	}

	return config.Client()
}
