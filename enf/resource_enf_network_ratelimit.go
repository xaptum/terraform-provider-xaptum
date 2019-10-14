package enf

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/xaptum/go-enf/enf"
)

func resourceEnfNetworkRateLimit() *schema.Resource {
	return &schema.Resource{
		Create: resourceEnfNetworkRateLimitCreate,
		Read:   resourceEnfNetworkRateLimitRead,
		Update: resourceEnfNetworkRateLimitUpdate,
		Delete: resourceEnfNetworkRateLimitDelete,

		Schema: map[string]*schema.Schema{
			"network": {
				Type:     schema.TypeString,
				Required: true,
			},
			"packets_per_second": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"packets_burst_size": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"bytes_per_second": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"bytes_burst_size": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"inherit": {
				Type:     schema.TypeBool,
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

func resourceEnfNetworkRateLimitCreate(d *schema.ResourceData, meta interface{}) error {
	// The ENF automatically creates a rate limit for each domain, so just update the values.
	return resourceEnfNetworkRateLimitUpdate(d, meta)
}

func resourceEnfNetworkRateLimitRead(d *schema.ResourceData, meta interface{}) error {
	switch d.Get("type") {
	case "default":
		return networkDefaultRateLimitRead(d, meta)
	case "max":
		return networkMaxRateLimitRead(d, meta)
	default:
		return fmt.Errorf("Invalid rate limit type %s. The type must be either default or max", d.Get("type"))
	}
}

func networkDefaultRateLimitRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*EnfClient).Client
	networkRateLimit, _, err := client.Network.GetDefaultEndpointRateLimits(context.Background(), d.Get("network").(string))
	if err != nil {
		return err
	}

	d.SetId(d.Get("network").(string))
	d.Set("packets_per_second", networkRateLimit.PacketsPerSecond)
	d.Set("packets_burst_size", networkRateLimit.PacketsBurstSize)
	d.Set("bytes_per_second", networkRateLimit.BytesPerSecond)
	d.Set("bytes_burst_size", networkRateLimit.BytesBurstSize)
	d.Set("inherit", networkRateLimit.Inherit)

	return nil
}

func networkMaxRateLimitRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*EnfClient).Client
	networkRateLimit, _, err := client.Network.GetMaxDefaultEndpointRateLimits(context.Background(), d.Get("network").(string))
	if err != nil {
		return err
	}

	d.SetId(d.Get("network").(string))
	d.Set("packets_per_second", networkRateLimit.PacketsPerSecond)
	d.Set("packets_burst_size", networkRateLimit.PacketsBurstSize)
	d.Set("bytes_per_second", networkRateLimit.BytesPerSecond)
	d.Set("bytes_burst_size", networkRateLimit.BytesBurstSize)
	d.Set("inherit", networkRateLimit.Inherit)

	return nil
}

func resourceEnfNetworkRateLimitUpdate(d *schema.ResourceData, meta interface{}) error {
	if err := validateRateLimitRequest(d); err != nil {
		return err
	}

	switch d.Get("type") {
	case "default":
		return networkDefaultRateLimitUpdate(d, meta)
	case "max":
		return networkMaxRateLimitUpdate(d, meta)
	default:
		return errors.New("The rate limit type must be either default or max")
	}
}

func networkDefaultRateLimitUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*EnfClient).Client
	networkRateLimitRequest := &enf.NetworkRateLimits{
		PacketsPerSecond: enf.Int(d.Get("packets_per_second").(int)),
		PacketsBurstSize: enf.Int(d.Get("packets_burst_size").(int)),
		BytesPerSecond:   enf.Int(d.Get("bytes_per_second").(int)),
		BytesBurstSize:   enf.Int(d.Get("bytes_burst_size").(int)),
		Inherit:          enf.Bool(d.Get("inherit").(bool)),
	}

	_, _, err := client.Network.SetDefaultEndpointRateLimits(context.Background(), networkRateLimitRequest, d.Get("network").(string))
	if err != nil {
		return err
	}

	return resourceEnfNetworkRateLimitRead(d, meta)
}

func networkMaxRateLimitUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*EnfClient).Client
	networkRateLimitRequest := &enf.NetworkRateLimits{
		PacketsPerSecond: enf.Int(d.Get("packets_per_second").(int)),
		PacketsBurstSize: enf.Int(d.Get("packets_burst_size").(int)),
		BytesPerSecond:   enf.Int(d.Get("bytes_per_second").(int)),
		BytesBurstSize:   enf.Int(d.Get("bytes_burst_size").(int)),
		Inherit:          enf.Bool(d.Get("inherit").(bool)),
	}

	_, _, err := client.Network.SetMaxDefaultEndpointRateLimits(context.Background(), networkRateLimitRequest, d.Get("network").(string))
	if err != nil {
		return err
	}

	return resourceEnfNetworkRateLimitRead(d, meta)
}

func resourceEnfNetworkRateLimitDelete(d *schema.ResourceData, meta interface{}) error {
	// Since we cannot "delete" a rate limit, we'll instead reset to the default values by updating the inherited flag to 'true'.
	d.Set("inherit", true)

	// We don't want to call the validate function, so we will use the helper
	// update functions rather than the main update function.
	switch d.Get("type") {
	case "default":
		err := networkDefaultRateLimitUpdate(d, meta)
		if err != nil {
			return err
		}
	case "max":
		err := networkMaxRateLimitUpdate(d, meta)
		if err != nil {
			return err
		}
	default:
		return errors.New("The rate limit type must be either current or max")
	}
	return nil
}
