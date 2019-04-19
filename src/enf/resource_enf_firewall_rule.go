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

// type firewallRules struct {
//     firewallRule firewallRule
// }

type firewallRule struct {
    Id   string `json:"id"`
    Priority   int `json:"priority"`
    Protocol    string    `json:"protocol"`
    Direction string `json:"direction"`
    SourceIP string `json:"source_ip"`
    SourcePort int `json:"source_port"`
    DestIP string `json:"dest_ip"`
    DestPort int `json:"dest_port"`
    Ack string `json:"ack"`
    Action string `json:"action"`
    IPFamily string `json:"ip_family"`
    Network string `json:"network"`
}



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


        var bearer = "Bearer " + m.(*EnfClient).ApiToken

        // Create a new request
        req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))

         // add authorization header to the request
        req.Header.Add("Authorization", bearer)
        req.Header.Add("content-type", "application/json")
        req.Header.Add("accept", "application/json")

        log.Printf("[DEBUG] enfFirewallRuleCreate() request is: ", req)


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
        log.Printf("[DEBUG] domain_url from m interface is: ", domain_url)

        network := d.Get("network").(string)
        url := domain_url + "/api/xfw/v1/" + network + "/rule"


        var bearer = "Bearer " + m.(*EnfClient).ApiToken

        // Create a new request
        req, err := http.NewRequest("GET", url, nil)

         // add authorization header to the request
        req.Header.Add("Authorization", bearer)
        req.Header.Add("content-type", "application/json")
        req.Header.Add("accept", "application/json")

        log.Printf("[DEBUG] enfFirewallRuleRead() request is: ", req)

        //create client to make the request
        client := m.(*EnfClient).HTTPClient
        response, err := client.Do(req)
        log.Printf("[DEBUG] enfFirewallRuleRead() error is: ", err)
        if err != nil {
                log.Printf("[ERROR] The HTTP request failed with error %s\n", err)
                return err
        } else {
            data, _ := ioutil.ReadAll(response.Body)
            log.Printf("[DEBUG] enfFirewallRuleRead() response code is: ", response.StatusCode)
            log.Printf("[DEBUG] enfFirewallRuleRead() response is: ", (response))
            log.Printf("[DEBUG] enfFirewallRuleRead() response.Body: ", response.Body)
            log.Printf("[DEBUG] enfFirewallRuleRead() string(data): ", string(data))


            //var fw_rules firewallRules
            var fw_rule []firewallRule
            json.Unmarshal([]byte(data), &fw_rule)
            //fw_rule = fw_rules

            log.Printf("[DEBUG] enfFirewallRuleRead() fw_rule[0] is: ", fw_rule[0])
            log.Printf("[DEBUG] enfFirewallRuleRead() fw_rule.Id is: ", fw_rule[0].Id)




        }

        d.SetId(network)//TODO: response should have ID
        return nil
}

func enfFirewallRuleUpdate(d *schema.ResourceData, m interface{}) error {
        //TODO: do delete and add API calls here

        return nil
        //return enfFirewallRuleRead(d, m)
}

func enfFirewallRuleDelete(d *schema.ResourceData, m interface{}) error {
        
        domain_url := m.(*EnfClient).DomainURL
        network := d.Get("network").(string)
        id := d.Get("rule_id").(string)
        url := domain_url + "/api/xfw/v1/" + network + "/rule/" + id


        var bearer = "Bearer " + m.(*EnfClient).ApiToken

        // Create a new request
        req, err := http.NewRequest("DELETE", url, nil)

         // add authorization header to the request
        req.Header.Add("Authorization", bearer)
        req.Header.Add("content-type", "application/json")
        req.Header.Add("accept", "application/json")

        log.Printf("[DEBUG] enfFirewallRuleDelete() request is: ", req)

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


