package enf

import (
        "github.com/hashicorp/terraform/helper/schema"
        "fmt"
        "net/http"
        "io/ioutil"
        "encoding/json"
        "bytes"

)

func enfFirewallRule() *schema.Resource {
        return &schema.Resource{
                Create: enfFirewallRuleCreate,
                Read:   enfFirewallRuleRead,
                Update: enfFirewallRuleUpdate,
                Delete: enfFirewallRuleDelete,

                Schema: map[string]*schema.Schema{
                        "host": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },

                        "network": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                        "rule_id": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                },
        }
}

func enfFirewallRuleCreate(d *schema.ResourceData, m interface{}) error {
        host := d.Get("host").(string)
        network := d.Get("network").(string)
        url := "https://" + host + "/api/xfw/v1/" + network + "/rule"
        jsonData := map[string]string{"rule": "400"}
        jsonValue, _ := json.Marshal(jsonData)


        var bearer = "Bearer " + Cred_token_string

        // Create a new request
        req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))

         // add authorization header to the request
        req.Header.Add("Authorization", bearer)

        //create client to make the request
        client := &http.Client{}
        response, err := client.Do(req)
        if err != nil {
                fmt.Printf("The HTTP request failed with error %s\n", err)
        } else {
            data, _ := ioutil.ReadAll(response.Body)
            fmt.Println(string(data))
        }

        d.SetId(network)


        return nil
}

func enfFirewallRuleRead(d *schema.ResourceData, m interface{}) error {

        fmt.Println(Cred_token_byte)

        host := d.Get("host").(string)
        network := d.Get("network").(string)
        url := "https://" + host + "/api/xfw/v1/" + network + "/rule"

        response, err := http.Get(url)
        if err != nil {
                fmt.Printf("The HTTP request failed with error %s\n", err)
        } else {
            data, _ := ioutil.ReadAll(response.Body)
            fmt.Println(string(data))
        }

        d.SetId(network) //sets Id, if blank then is destroyed 
        return nil
}

func enfFirewallRuleUpdate(d *schema.ResourceData, m interface{}) error {
        return enfFirewallRuleRead(d, m)
}

func enfFirewallRuleDelete(d *schema.ResourceData, m interface{}) error {
        host := d.Get("host").(string)
        network := d.Get("network").(string)
        id := d.Get("rule_id").(string)
        url := "https://" + host + "/api/xfw/v1/" + network + "/rule/" + id


        var bearer = "Bearer " + Cred_token_string

        // Create a new request
        req, err := http.NewRequest("DELETE", url, nil)

         // add authorization header to the request
        req.Header.Add("Authorization", bearer)

        //create client to make the request
        client := &http.Client{}
        response, err := client.Do(req)
        if err != nil {
                fmt.Printf("The HTTP request failed with error %s\n", err)
        } else {
            data, _ := ioutil.ReadAll(response.Body)
            fmt.Println(string(data))
        }

        d.SetId(network)

        return nil
}


