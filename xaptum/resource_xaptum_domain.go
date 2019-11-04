package xaptum

import (
	"context"
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/xaptum/go-enf/enf"
)

func resourceXaptumDomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceXaptumDomainCreate,
		Read:   resourceXaptumDomainRead,
		Update: resourceXaptumDomainUpdate,
		Delete: resourceXaptumDomainDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "CUSTOMER_SOURCE",
			},
			"network": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"admin_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceXaptumDomainCreate(d *schema.ResourceData, meta interface{}) error {
	// Need to validate that the type is either (1) not set or (2) equal to CUSTOMER_SOURCE
	_, typeIsSet := d.GetOkExists("type")
	if typeIsSet && d.Get("type").(string) != "CUSTOMER_SOURCE" {
		return errors.New("type must be either (1) not set or (2) equal to CUSTOMER_SOURCE")
	}

	// Also need to check if the status field is set, since this is not allowed.
	_, statusIsSet := d.GetOkExists("status")
	if statusIsSet {
		return errors.New("status will be computed initially, so it cannot be set during resource creation")
	}

	req := &enf.DomainRequest{
		Name:       enf.String(d.Get("name").(string)),
		Type:       enf.String(d.Get("type").(string)),
		AdminName:  enf.String(d.Get("admin_name").(string)),
		AdminEmail: enf.String(d.Get("admin_email").(string)),
	}

	client := meta.(*EnfClient).Client
	domain, _, err := client.Domains.CreateDomain(context.Background(), req)
	if err != nil {
		return err
	}

	// Use the domain's ::/48 address as the identifier.
	d.SetId(*domain.Network)

	return resourceXaptumDomainRead(d, meta)
}

func resourceXaptumDomainRead(d *schema.ResourceData, meta interface{}) error {
	address := d.Id()

	client := meta.(*EnfClient).Client
	domain, _, err := client.Domains.GetDomain(context.Background(), address)
	if err != nil {
		return err
	}

	// name, admin_name, admin_email, and type should all be set
	d.Set("network", domain.Network)
	d.Set("status", domain.Status)

	return nil
}

func resourceXaptumDomainUpdate(d *schema.ResourceData, meta interface{}) error {
	// Update should only happen for activating/deactivating a domain.
	err := validateUpdateDomainRequest(d)
	if err != nil {
		return err
	}

	client := meta.(*EnfClient).Client
	address := d.Get("network").(string)

	switch d.Get("status") {
	case "READY":
		_, _, err := client.Domains.DeactivateDomain(context.Background(), address)
		if err != nil {
			return err
		}
	case "ACTIVE":
		_, _, err := client.Domains.ActivateDomain(context.Background(), address)
		if err != nil {
			return err
		}
	default:
		return errors.New("status must be either READY or ACTIVE")
	}

	return resourceXaptumDomainRead(d, meta)
}

func resourceXaptumDomainDelete(d *schema.ResourceData, meta interface{}) error {
	return errors.New("The domain resource does not support deletion")
}
