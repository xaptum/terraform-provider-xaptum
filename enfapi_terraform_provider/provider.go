package main

//fixed missing package with: go get github.com/hashicorp/terraform
import (
        "github.com/hashicorp/terraform/helper/schema"
        "net/http"
        "encoding/json"
        "fmt"
        "bytes"
        "io/ioutil"
)
type Data struct {
    Credentials Credentials
}

type Credentials struct {
    Username   string `json:"username"`
    Token   string `json:"token"`
    UserID    string    `json:"user_id"`
    Type string `json:"type"`
    DomainID string `json:"domain_id"`
    DomainNetwork string `json:"domain_network"`
}



//global token variable
var Cred_token_byte []byte
var Cred_token_string string  

func Provider() *schema.Provider {
        //authenticate
        jsonData := map[string]string{"username": "xap@admin", "token": "Test1234"}
        jsonValue, _ := json.Marshal(jsonData)
        response, err := http.Post("https://dev.xaptum.io/api/xcr/v2/xauth", "application/json", bytes.NewBuffer(jsonValue))
        if err != nil {
              fmt.Printf("The HTTP request failed with error %s\n", err)
        } else {
                data_body, _ := ioutil.ReadAll(response.Body)

                var data Data
                var cred Credentials
                json.Unmarshal([]byte(data_body), &data)

                cred = data.Credentials
                var cred_token string
                cred_token = cred.Token
                Cred_token_string = cred_token
                Cred_token_byte = []byte(cred_token)
        }





        return &schema.Provider{
                ResourcesMap: map[string]*schema.Resource{
                                "enf_firewall": enfFirewall(),
                                "enf_domain": enfDomain(),
                                "enf_network": enfNetwork(),
                                "enf_connection": enfConnection(),
                		"enf_group": enfGroup(),
                		"enf_endpoint": enfEndpoint(),
                                "enf_ratelimit": enfRatelimit(),

                	},
        }
}