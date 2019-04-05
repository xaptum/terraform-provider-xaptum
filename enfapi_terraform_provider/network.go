package main

import (
        "github.com/hashicorp/terraform/helper/schema"
)

func network() *schema.Resource {
        return &schema.Resource{
                Create: networkCreate,
                Read:   networkRead,
                Update: networkUpdate,
                Delete: networkDelete,

                Schema: map[string]*schema.Schema{
                        "address": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                },
        }
}

func networkCreate(d *schema.ResourceData, m interface{}) error {
        address := d.Get("address").(string)
        d.SetId(address)
        return enfGroupRead(d, m)
}

func networkRead(d *schema.ResourceData, m interface{}) error {
        return nil
}

func networkUpdate(d *schema.ResourceData, m interface{}) error {
        return resourceServerRead(d, m)
}

func networkDelete(d *schema.ResourceData, m interface{}) error {
        return nil
}


