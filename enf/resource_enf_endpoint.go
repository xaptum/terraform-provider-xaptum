package enf

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func enfEndpoint() *schema.Resource {
	return &schema.Resource{
		Create: enfEndpointCreate,
		Read:   enfEndpointRead,
		Update: enfEndpointUpdate,
		Delete: enfEndpointDelete,

		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func enfEndpointCreate(d *schema.ResourceData, m interface{}) error {
	address := d.Get("address").(string)
	d.SetId(address)
	return enfEndpointRead(d, m)
}

func enfEndpointRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func enfEndpointUpdate(d *schema.ResourceData, m interface{}) error {
	return enfEndpointRead(d, m)
}

func enfEndpointDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
