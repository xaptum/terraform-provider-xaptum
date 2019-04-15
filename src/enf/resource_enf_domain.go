package enf

import (
        "github.com/hashicorp/terraform/helper/schema"
)

func enfDomain() *schema.Resource {
        return &schema.Resource{
                Create: enfDomainCreate,
                Read:   enfDomainRead,
                Update: enfDomainUpdate,
                Delete: enfDomainDelete,

                Schema: map[string]*schema.Schema{
                        "domain_id": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                        "domain_network": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                        "network_cidr": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                },
        }
}

func enfDomainCreate(d *schema.ResourceData, m interface{}) error {
        domain_id := d.Get("domain_id").(string)
        d.SetId(domain_id)
        return enfDomainRead(d, m)
}

func enfDomainRead(d *schema.ResourceData, m interface{}) error {
        return nil
}

func enfDomainUpdate(d *schema.ResourceData, m interface{}) error {
        return enfDomainRead(d, m)
}

func enfDomainDelete(d *schema.ResourceData, m interface{}) error {
        return nil
}


