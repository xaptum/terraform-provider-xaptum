package main

import (
        "github.com/hashicorp/terraform/helper/schema"
)

func enfRatelimit() *schema.Resource {
        return &schema.Resource{
                Create: enfRatelimitCreate,
                Read:   enfRatelimitRead,
                Update: enfRatelimitUpdate,
                Delete: enfRatelimitDelete,

                Schema: map[string]*schema.Schema{
                        "limit": &schema.Schema{
                                Type:     schema.TypeInt,
                                Required: true,
                        },
                        "ipv6": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },  
                        "network": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                        "api_uri": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        }, 
                        "filter": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },   
                        "domain_network": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },                                
                },
        }
}

func enfRatelimitCreate(d *schema.ResourceData, m interface{}) error {
        subnet := d.Get("subnet").(string)
        d.SetId(subnet)
        return enfRatelimitRead(d, m)
}

func enfRatelimitRead(d *schema.ResourceData, m interface{}) error {
        return nil
}

func enfRatelimitUpdate(d *schema.ResourceData, m interface{}) error {
        return enfRatelimitRead(d, m)
}

func enfRatelimitDelete(d *schema.ResourceData, m interface{}) error {
        return nil
}


