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

type firewallRuleCreate struct {
    IP_family string `json:"ip_family"`
    Direction string `json:"direction"`
    Action string `json:"action"`
    Priority int `json:"priority"`
    Protocol string `json:"protocol"`
    SourceIP string `json:"source_ip"`
    SourcePort int `json:"source_port"`
    DestIP string `json:"dest_ip"`
    DestPort int `json:"dest_port"`
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
                        "priority": &schema.Schema{
                                Type:     schema.TypeInt,
                                Required: true,
                        },
                        "protocol": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                        "direction": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                        "source_ip": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Default: "*",
                        },
                        "source_port": &schema.Schema{
                                Type:     schema.TypeInt,
                                Optional: true,
                                Default: 0,
                        },
                        "dest_ip": &schema.Schema{
                                Type:     schema.TypeString,
                                Optional: true,
                                Default: "*",
                        },
                        "dest_port": &schema.Schema{
                                Type:     schema.TypeInt,
                                Optional: true,
                                Default: 0,
                        },
                        "action": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                        "ip_family": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },

                },
        }
}


func enfFirewallRuleCreate(d *schema.ResourceData, m interface{}) error {

    if d.Get("id").(string) != "" {

        domain_url := m.(*EnfClient).DomainURL
        network := d.Get("network").(string)
        url := domain_url + "/api/xfw/v1/" + network + "/rule"


        var newRule firewallRuleCreate
        newRule.IP_family = d.Get("ip_family").(string)
        newRule.Direction = d.Get("direction").(string)
        newRule.Action = d.Get("action").(string)
        newRule.Priority = d.Get("priority").(int)
        newRule.Protocol = d.Get("protocol").(string)

        jsonData := "ip_family: " + newRule.IP_family + "direction: " + newRule.Direction + "action: " + newRule.Action + "priority: " + string(newRule.Priority) + "protocol: " + newRule.Protocol  
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

            //TODO: assign an ID to the newly created rule

        }

        d.SetId(network)//TODO: response should have ID

    }

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

            log.Printf("[DEBUG] enfFirewallRuleRead() len(fw_rule) is: ", len(fw_rule))
            for i := 0; i < len(fw_rule); i++ {
              d.SetId(fw_rule[i].Id)
            }

        }
        //d.SetId(network)//TODO: response should have ID
        return nil
}

func enfFirewallRuleUpdate(d *schema.ResourceData, m interface{}) error {
        //TODO: do delete and add API calls here
        enfFirewallRuleDelete(d, m)
        return enfFirewallRuleCreate(d, m)

        //return nil

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

        //d.SetId(network)//TODO: response should have ID

        return nil
}


