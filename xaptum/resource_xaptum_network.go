package xaptum

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/xaptum/go-enf/enf"
)

func resourceXaptumNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceEnfNetworkCreate,
		Read:   resourceEnfNetworkRead,
		Update: resourceEnfNetworkUpdate,
		Delete: resourceEnfNetworkDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceEnfNetworkCreate(d *schema.ResourceData, meta interface{}) error {
	domainID := meta.(*EnfClient).DomainID

	params := &enf.NetworkRequest{
		Name:        enf.String(d.Get("name").(string)),
		Description: enf.String(d.Get("description").(string)),
	}

	client := meta.(*EnfClient).Client
	network, _, err := client.Network.CreateNetwork(context.Background(), strconv.FormatInt(domainID, 10), params)
	if err != nil {
		return err
	}

	d.SetId(*network.Network)

	return resourceEnfNetworkRead(d, meta)
}

func resourceEnfNetworkRead(d *schema.ResourceData, meta interface{}) error {
	address := d.Id() // this is the network address

	client := meta.(*EnfClient).Client
	network, resp, err := client.Network.GetNetwork(context.Background(), address)
	if err != nil {
		if resp.StatusCode == http.StatusNotFound {
			d.SetId("")
			return nil
		}
		return err
	}

	d.Set("name", network.Name)
	d.Set("description", network.Description)
	d.Set("status", network.Status)

	return nil
}

func resourceEnfNetworkUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*EnfClient).Client

	fields := &enf.NetworkRequest{
		Name:        enf.String(d.Get("name").(string)),
		Description: enf.String(d.Get("description").(string)),
	}
	_, _, err := client.Network.UpdateNetwork(context.Background(), d.Id(), fields)
	if err != nil {
		return fmt.Errorf("Error updating network: %s", err)
	}

	return resourceEnfNetworkRead(d, meta)
}

func resourceEnfNetworkDelete(d *schema.ResourceData, meta interface{}) error {
	return errors.New("The network resource does not support deletion")
}
