package enf

import (
        "github.com/hashicorp/terraform/helper/schema"
        "fmt"
        "net/http"
        "io/ioutil"
        "encoding/json"
        "bytes"
        "log"

)

func enfFirewallRule() *schema.Resource {
        return &schema.Resource{
                Create: enfFirewallRuleCreate,
                Read:   enfFirewallRuleRead,
                Update: enfFirewallRuleUpdate,
                Delete: enfFirewallRuleDelete,

                Schema: map[string]*schema.Schema{
                        "network": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                },
        }
}

func enfFirewallRuleCreate(d *schema.ResourceData, m interface{}) error {


        domain_url := m.(*EnfClient).DomainURL
        network := d.Get("network").(string)
        url := domain_url + "/api/xfw/v1/" + network + "/rule"




        jsonData := map[string]string{"rule": "400"}
        jsonValue, _ := json.Marshal(jsonData)


        var bearer = "Bearer" + m.(*EnfClient).ApiToken

        // Create a new request
        req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))

         // add authorization header to the request
        req.Header.Add("Authorization", bearer)

        //create client to make the request
        client := m.(*EnfClient).HTTPClient
        response, err := client.Do(req)
        if err != nil {
                log.Printf("[ERROR] The HTTP request failed with error %s\n", err)
                return err        
        } else {
            data, _ := ioutil.ReadAll(response.Body)
            log.Printf("IN FIREWALL.GO CREATE: ", string(data))
        }

        d.SetId(network)//TODO: response should have ID


        return nil
}

func enfFirewallRuleRead(d *schema.ResourceData, m interface{}) error {

        domain_url := m.(*EnfClient).DomainURL
        network := d.Get("network").(string)
        url := domain_url + "/api/xfw/v1/" + network + "/rule"


        var bearer = "Bearer" + m.(*EnfClient).ApiToken

        // Create a new request
        req, err := http.NewRequest("GET", url, nil)

         // add authorization header to the request
        req.Header.Add("Authorization", bearer)

        //create client to make the request
        client := m.(*EnfClient).HTTPClient
        response, err := client.Do(req)
        if err != nil {
                log.Printf("[ERROR] The HTTP request failed with error %s\n", err)
                return err
        } else {
            data, _ := ioutil.ReadAll(response.Body)
            log.Printf("IN FIREWALL.GO READ: ", string(data))

        }

        d.SetId(network)//TODO: response should have ID
        return nil
}

func enfFirewallRuleUpdate(d *schema.ResourceData, m interface{}) error {
        //TODO: do delete and add API calls here


        //return enfFirewallRuleRead(d, m)
}

func enfFirewallRuleDelete(d *schema.ResourceData, m interface{}) error {
        
        domain_url := m.(*EnfClient).DomainURL
        network := d.Get("network").(string)
        id := d.Get("rule_id").(string)
        url := domain_url + "/api/xfw/v1/" + network + "/rule/" + id


        var bearer = "Bearer" + m.(*EnfClient).ApiToken

        // Create a new request
        req, err := http.NewRequest("DELETE", url, nil)

         // add authorization header to the request
        req.Header.Add("Authorization", bearer)

        //create client to make the request
        client := m.(*EnfClient).HTTPClient
        response, err := client.Do(req)
        if err != nil {
                fmt.Printf("The HTTP request failed with error %s\n", err)
        } else {
            data, _ := ioutil.ReadAll(response.Body)
            fmt.Println(string(data))
        }

        d.SetId(network)//TODO: response should have ID

        return nil
}


