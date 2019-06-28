package enf

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func enfConnection() *schema.Resource {
	return &schema.Resource{
		Create: enfConnectionCreate,
		Read:   enfConnectionRead,
		Update: enfConnectionUpdate,
		Delete: enfConnectionDelete,

		Schema: map[string]*schema.Schema{
			"ipv6": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func enfConnectionCreate(d *schema.ResourceData, m interface{}) error {
	ipv6 := d.Get("ipv6").(string)
	d.SetId(ipv6)
	return enfGroupRead(d, m)
}

func enfConnectionRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func enfConnectionUpdate(d *schema.ResourceData, m interface{}) error {
	return enfConnectionRead(d, m)
}

func enfConnectionDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
