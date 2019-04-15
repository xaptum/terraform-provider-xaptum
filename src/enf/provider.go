package enf

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

type Config struct {
    ApiToken  string
    BaseURL string
}


//global token variable
var Cred_token_byte []byte
var Cred_token_string string  

var config Config

func Provider() *schema.Provider {
        //authenticate
        jsonData := map[string]string{"username": "xap@admin", "token": "Test1234"}
        jsonValue, _ := json.Marshal(jsonData)
        base_url := "https://dev.xaptum.io"
        response, err := http.Post(base_url + "/api/xcr/v2/xauth", "application/json", bytes.NewBuffer(jsonValue))
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

            //set environment variables 
            Schema: map[string]*schema.Schema{
                "api_token": {
                    Type:        schema.TypeString,
                    Optional:    true,
                    DefaultFunc: schema.EnvDefaultFunc("ENF_API_TOKEN", Cred_token_string),
                    Description: "Token from authenticating with dev.xaptum.io",
                },
                "base_url": {
                    Type:        schema.TypeString,
                    Optional:    true,
                    DefaultFunc: schema.EnvDefaultFunc("ENF_API_URL", base_url),
                    Description: "Base URL for authentication",
                },
            },

            ConfigureFunc: providerConfigure,

                ResourcesMap: map[string]*schema.Resource{
                                "enf_firewall": enfFirewallRule(),
                                "enf_domain": enfDomain(),
                                "enf_network": enfNetwork(),
                                "enf_connection": enfConnection(),
                	            "enf_group": enfGroup(),
                		        "enf_endpoint": enfEndpoint(),
                                "enf_ratelimit": enfRatelimit(),

                	},
        }
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
    config := Config{
        ApiToken:  d.Get("api_key").(string),
        BaseURL: d.Get("base_url").(string),
    }
    return &config, nil
}
