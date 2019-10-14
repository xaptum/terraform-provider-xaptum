package enf

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/xaptum/go-enf/enf"
)

func resourceEnfDomainRateLimit() *schema.Resource {
	return &schema.Resource{
		Create: resourceEnfDomainRateLimitCreate,
		Read:   resourceEnfDomainRateLimitRead,
		Update: resourceEnfDomainRateLimitUpdate,
		Delete: resourceEnfDomainRateLimitDelete,

		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"packets_per_second": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"packets_burst_size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"bytes_per_second": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"bytes_burst_size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceEnfDomainRateLimitCreate(d *schema.ResourceData, meta interface{}) error {
	// The ENF automatically creates a rate limit for each domain, so just update the values.
	return resourceEnfDomainRateLimitUpdate(d, meta)
}

func resourceEnfDomainRateLimitRead(d *schema.ResourceData, meta interface{}) error {
	switch d.Get("type") {
	case "default":
		return domainDefaultRateLimitRead(d, meta)
	case "max":
		return domainMaxRateLimitRead(d, meta)
	default:
		return fmt.Errorf("Invalid rate limit type %s. The type must be either default or max", d.Get("type"))
	}
}

func domainDefaultRateLimitRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*EnfClient).Client
	domainRateLimit, resp, err := client.Domains.GetDefaultEndpointRateLimits(context.Background(), d.Get("domain").(string))
	if err != nil {
		if resp.StatusCode == http.StatusNotFound {
			d.SetId("")
			return nil
		}
		return err
	}

	d.SetId(d.Get("domain").(string))
	d.Set("packets_per_second", domainRateLimit.PacketsPerSecond)
	d.Set("packets_burst_size", domainRateLimit.PacketsBurstSize)
	d.Set("bytes_per_second", domainRateLimit.BytesPerSecond)
	d.Set("bytes_burst_size", domainRateLimit.BytesBurstSize)

	return nil
}

func domainMaxRateLimitRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*EnfClient).Client
	domainRateLimit, _, err := client.Domains.GetMaxDefaultEndpointRateLimits(context.Background(), d.Get("domain").(string))
	if err != nil {
		return err
	}

	d.SetId(d.Get("domain").(string))
	d.Set("packets_per_second", domainRateLimit.PacketsPerSecond)
	d.Set("packets_burst_size", domainRateLimit.PacketsBurstSize)
	d.Set("bytes_per_second", domainRateLimit.BytesPerSecond)
	d.Set("bytes_burst_size", domainRateLimit.BytesBurstSize)

	return nil
}

func resourceEnfDomainRateLimitUpdate(d *schema.ResourceData, meta interface{}) error {
	switch d.Get("type") {
	case "default":
		return domainDefaultRateLimitUpdate(d, meta)
	case "max":
		return domainMaxRateLimitUpdate(d, meta)
	default:
		return fmt.Errorf("Invalid rate limit type %s. The type must be either default or max", d.Get("type"))
	}
}

func domainDefaultRateLimitUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*EnfClient).Client
	domainRateLimitRequest := &enf.DomainRateLimits{
		PacketsPerSecond: enf.Int(d.Get("packets_per_second").(int)),
		PacketsBurstSize: enf.Int(d.Get("packets_burst_size").(int)),
		BytesPerSecond:   enf.Int(d.Get("bytes_per_second").(int)),
		BytesBurstSize:   enf.Int(d.Get("bytes_burst_size").(int)),
	}
	_, _, err := client.Domains.SetDefaultEndpointRateLimits(context.Background(), domainRateLimitRequest, d.Get("domain").(string))
	if err != nil {
		return err
	}

	return resourceEnfDomainRateLimitRead(d, meta)
}

func domainMaxRateLimitUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*EnfClient).Client
	domainRateLimitRequest := &enf.DomainRateLimits{
		PacketsPerSecond: enf.Int(d.Get("packets_per_second").(int)),
		PacketsBurstSize: enf.Int(d.Get("packets_burst_size").(int)),
		BytesPerSecond:   enf.Int(d.Get("bytes_per_second").(int)),
		BytesBurstSize:   enf.Int(d.Get("bytes_burst_size").(int)),
	}
	_, _, err := client.Domains.SetMaxDefaultEndpointRateLimits(context.Background(), domainRateLimitRequest, d.Get("domain").(string))
	if err != nil {
		return err
	}

	return resourceEnfDomainRateLimitRead(d, meta)
}

func resourceEnfDomainRateLimitDelete(d *schema.ResourceData, meta interface{}) error {
	if d.Get("type") == "max" {
		return errors.New("The domain max rate limit does not support deletion")
	}

	// If the type of the rate limit being deleted is default, set the default values to the max rate limit values to simulate deletion.
	domainMaxRateLimit, _, err := meta.(*EnfClient).Client.Domains.GetMaxDefaultEndpointRateLimits(context.Background(), d.Get("domain").(string))
	if err != nil {
		return err
	}

	domainDefaultRateLimitRequest := &enf.DomainRateLimits{
		PacketsPerSecond: domainMaxRateLimit.PacketsPerSecond,
		PacketsBurstSize: domainMaxRateLimit.PacketsBurstSize,
		BytesPerSecond:   domainMaxRateLimit.BytesPerSecond,
		BytesBurstSize:   domainMaxRateLimit.BytesBurstSize,
	}

	_, _, err = meta.(*EnfClient).Client.Domains.SetDefaultEndpointRateLimits(context.Background(), domainDefaultRateLimitRequest, d.Get("domain").(string))
	if err != nil {
		return err
	}

	return nil
}
