package enf

import (
	"context"
	"errors"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/xaptum/go-enf/enf"
)

func resourceEnfEndpointRateLimit() *schema.Resource {
	return &schema.Resource{
		Create: resourceEnfEndpointRateLimitCreate,
		Read:   resourceEnfEndpointRateLimitRead,
		Update: resourceEnfEndpointRateLimitUpdate,
		Delete: resourceEnfEndpointRateLimitDelete,

		Schema: map[string]*schema.Schema{
			"endpoint": {
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
				Default:  "current",
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceEnfEndpointRateLimitCreate(d *schema.ResourceData, meta interface{}) error {
	// The ENF automatically creates a rate limit for each endpoint, so just update the values.
	return resourceEnfEndpointRateLimitUpdate(d, meta)
}

func resourceEnfEndpointRateLimitRead(d *schema.ResourceData, meta interface{}) error {
	switch d.Get("type") {
	case "current":
		return endpointCurrentRateLimitRead(d, meta)
	case "max":
		return endpointMaxRateLimitRead(d, meta)
	default:
		return fmt.Errorf("Invalid rate limit type %s.The rate limit type must be either current or max", d.Get("type"))
	}
}

func endpointCurrentRateLimitRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*EnfClient).Client
	endpointRateLimit, _, err := client.Endpoint.GetCurrentRateLimits(context.Background(), d.Get("endpoint").(string))
	if err != nil {
		return err
	}

	d.SetId(d.Get("endpoint").(string))
	d.Set("packets_per_second", endpointRateLimit.PacketsPerSecond)
	d.Set("packets_burst_size", endpointRateLimit.PacketsBurstSize)
	d.Set("bytes_per_second", endpointRateLimit.BytesPerSecond)
	d.Set("bytes_burst_size", endpointRateLimit.BytesBurstSize)
	d.Set("inherit", endpointRateLimit.Inherit)
	return nil
}

func endpointMaxRateLimitRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*EnfClient).Client
	endpointRateLimit, _, err := client.Endpoint.GetMaxRateLimits(context.Background(), d.Get("endpoint").(string))
	if err != nil {
		return err
	}

	d.SetId(d.Get("endpoint").(string))
	d.Set("packets_per_second", endpointRateLimit.PacketsPerSecond)
	d.Set("packets_burst_size", endpointRateLimit.PacketsBurstSize)
	d.Set("bytes_per_second", endpointRateLimit.BytesPerSecond)
	d.Set("bytes_burst_size", endpointRateLimit.BytesBurstSize)
	d.Set("inherit", endpointRateLimit.Inherit)
	return nil
}
func resourceEnfEndpointRateLimitUpdate(d *schema.ResourceData, meta interface{}) error {
	if err := validateRateLimitRequest(d); err != nil {
		return err
	}

	if d.Get("type") == "current" {
		return endpointCurrentRateLimitUpdate(d, meta)
	} else if d.Get("type") == "max" {
		return endpointMaxRateLimitUpdate(d, meta)
	} else {
		return errors.New("The rate limit type must be either current or max")
	}
}

func endpointCurrentRateLimitUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*EnfClient).Client
	endpointRateLimitRequest := &enf.EndpointRateLimits{
		PacketsPerSecond: enf.Int(d.Get("packets_per_second").(int)),
		PacketsBurstSize: enf.Int(d.Get("packets_burst_size").(int)),
		BytesPerSecond:   enf.Int(d.Get("bytes_per_second").(int)),
		BytesBurstSize:   enf.Int(d.Get("bytes_burst_size").(int)),
		Inherit:          enf.Bool(d.Get("inherit").(bool)),
	}

	_, _, err := client.Endpoint.SetCurrentRateLimits(context.Background(), endpointRateLimitRequest, d.Get("endpoint").(string))
	if err != nil {
		return err
	}

	return resourceEnfEndpointRateLimitRead(d, meta)
}

func endpointMaxRateLimitUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*EnfClient).Client
	endpointRateLimitRequest := &enf.EndpointRateLimits{
		PacketsPerSecond: enf.Int(d.Get("packets_per_second").(int)),
		PacketsBurstSize: enf.Int(d.Get("packets_burst_size").(int)),
		BytesPerSecond:   enf.Int(d.Get("bytes_per_second").(int)),
		BytesBurstSize:   enf.Int(d.Get("bytes_burst_size").(int)),
		Inherit:          enf.Bool(d.Get("inherit").(bool)),
	}

	_, _, err := client.Endpoint.SetMaxRateLimits(context.Background(), endpointRateLimitRequest, d.Get("endpoint").(string))
	if err != nil {
		return err
	}

	return resourceEnfEndpointRateLimitRead(d, meta)
}

func resourceEnfEndpointRateLimitDelete(d *schema.ResourceData, meta interface{}) error {
	// Since we cannot "delete" a rate limit, we'll instead reset to the default values by updating the inherited flag to 'true'.
	d.Set("inherit", true)

	// We don't want to call the validate function, so we will use the helper
	// update functions rather than the main update function.
	switch d.Get("type") {
	case "current":
		err := endpointCurrentRateLimitUpdate(d, meta)
		if err != nil {
			return err
		}
	case "max":
		err := endpointMaxRateLimitUpdate(d, meta)
		if err != nil {
			return err
		}
	default:
		return errors.New("The rate limit type must be either current or max")
	}
	return nil
}
