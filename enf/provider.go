package enf

import (
	"errors"

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
			"enf_domain_ratelimit":   resourceEnfDomainRateLimit(),
			"enf_endpoint_ratelimit": resourceEnfEndpointRateLimit(),
			"enf_firewall_rule":      resourceEnfFirewallRule(),
			"enf_network":            resourceEnfNetwork(),
			"enf_network_ratelimit":  resourceEnfNetworkRateLimit(),
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

func validateRateLimitRequest(d *schema.ResourceData) error {
	if d.Get("inherit").(bool) {
		// Assume the inherit flag on d is true, and make sure no other rate limit values are set.
		_, packetsPerSecondIsSet := d.GetOkExists("packets_per_second")
		_, packetsBurstSizeIsSet := d.GetOkExists("packets_burst_size")
		_, bytesPerSecondIsSet := d.GetOkExists("bytes_per_second")
		_, bytesBurstSizeIsSet := d.GetOkExists("bytes_burst_size")
		if packetsPerSecondIsSet || packetsBurstSizeIsSet || bytesPerSecondIsSet || bytesBurstSizeIsSet {
			return errors.New("If inherit is true then no other rate limit values can be set")
		}
	}

	return nil
}
