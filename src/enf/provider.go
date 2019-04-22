package enf

//fixed missing package with: go get github.com/hashicorp/terraform
import (
        "github.com/hashicorp/terraform/helper/schema"
        "log"
)



func Provider() *schema.Provider {


        return &schema.Provider{

            //set environment variables 
            Schema: map[string]*schema.Schema{
                "username": {
                    Type:        schema.TypeString,
                    Optional:    true,
                    DefaultFunc: schema.EnvDefaultFunc("ENF_USERNAME", nil),
                    Description: "Token from authenticating with dev.xaptum.io",
                },
                "password": {
                    Type:        schema.TypeString,
                    Optional:    true,
                    DefaultFunc: schema.EnvDefaultFunc("ENF_PASSWORD", nil),
                    Description: "Base URL for authentication",
                },
                "domain_url": {
                    Type:        schema.TypeString,
                    Optional:    true,
                    DefaultFunc: schema.EnvDefaultFunc("ENF_DOMAIN_URL", nil),
                    Description: "Base URL for authentication",
                },
            },

                ConfigureFunc: providerConfigure,//this will return config.Client(), which can be used in resource.go files

                ResourcesMap: map[string]*schema.Resource{
                                "enf_firewall": enfFirewallRule(),
                	},
        }
}

//this returned client can be used in the m interface{} argument for the resource.go files
func providerConfigure(d *schema.ResourceData) (interface{}, error) {

    config := Config{
        Username: d.Get("username").(string),
        Password: d.Get("password").(string),
        DomainURL: d.Get("domain_url").(string),

    }
    return config.Client()
}
