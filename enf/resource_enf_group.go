package enf

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func enfGroup() *schema.Resource {
	return &schema.Resource{
		Create: enfGroupCreate,
		Read:   enfGroupRead,
		Update: enfGroupUpdate,
		Delete: enfGroupDelete,

		Schema: map[string]*schema.Schema{
			"network": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func enfGroupCreate(d *schema.ResourceData, m interface{}) error {
	address := d.Get("address").(string)
	d.SetId(address)
	return enfGroupRead(d, m)
}

func enfGroupRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func enfGroupUpdate(d *schema.ResourceData, m interface{}) error {
	return enfGroupRead(d, m)
}

func enfGroupDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
