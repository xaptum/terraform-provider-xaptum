package enf

import (
        "github.com/hashicorp/terraform/helper/schema"
        "log"
)



func Provider() *schema.Provider {

        return &schema.Provider{

            Schema: map[string]*schema.Schema{
                "username": {
                    Type:        schema.TypeString,
                    Optional:    true,
                    DefaultFunc: schema.EnvDefaultFunc("ENF_USERNAME", nil),
                    Description: "Username for authenticating with dev.xaptum.io",
                },
                "password": {
                    Type:        schema.TypeString,
                    Optional:    true,
                    DefaultFunc: schema.EnvDefaultFunc("ENF_PASSWORD", nil),
                    Description: "Password for authenticating with dev.xaptum.io",
                },
                "domain_url": {
                    Type:        schema.TypeString,
                    Optional:    true,
                    DefaultFunc: schema.EnvDefaultFunc("ENF_DOMAIN_URL", nil),
                    Description: "Base URL for API calls",
                },
            },

                ConfigureFunc: providerConfigure,

                ResourcesMap: map[string]*schema.Resource{
                    },
        }
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

    config := Config{
        Username: d.Get("username").(string),
        Password: d.Get("password").(string),
        DomainURL: d.Get("domain_url").(string),

    }
    return config.Client()
}
