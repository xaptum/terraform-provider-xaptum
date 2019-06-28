package enf

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func enfNetwork() *schema.Resource {
	return &schema.Resource{
		Create: enfNetworkCreate,
		Read:   enfNetworkRead,
		Update: enfNetworkUpdate,
		Delete: enfNetworkDelete,

		Schema: map[string]*schema.Schema{
			"subnet": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func enfNetworkCreate(d *schema.ResourceData, m interface{}) error {
	subnet := d.Get("subnet").(string)
	d.SetId(subnet)
	return enfNetworkRead(d, m)
}

func enfNetworkRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func enfNetworkUpdate(d *schema.ResourceData, m interface{}) error {
	return enfNetworkRead(d, m)
}

func enfNetworkDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
