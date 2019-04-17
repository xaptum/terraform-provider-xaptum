package enf

//fixed missing package with: go get github.com/hashicorp/terraform
import (
        "github.com/hashicorp/terraform/helper/schema"
        "net/http"
        "encoding/json"
        "fmt"
        "bytes"
        "io/ioutil"
        "os"
)
type Response struct {
    Data []Data
    Page Pages
}

type Data struct {
    Username   string `json:"username"`
    Token   string `json:"token"`
    UserID    int    `json:"user_id"`
    Type string `json:"type"`
    DomainID int `json:"domain_id"`
    DomainNetwork string `json:"domain_network"`
}

type Pages struct {
    Curr int `json:"curr"`
    Next int `json: "next"`
    Prev int `json: "prev"` 
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

                var resp Response
                json.Unmarshal([]byte(data_body), &resp)

                fmt.Println("Returned data_body is: ", string(data_body))
                fmt.Println("Returned data is: ", (resp))
                fmt.Println("Creds: ", resp.Data)
                fmt.Println("Token:", resp.Data[0].Token)
                fmt.Println("Pages is: ", resp.Page)

                //var cred_token string
                //cred_token = cred.Token
                //Cred_token_string = cred_token
                //Cred_token_byte = []byte(cred_token)

                os.Setenv("ENF_API_TOKEN", resp.Data[0].Token)
                fmt.Println("ENF_API_TOKEN:", os.Getenv("ENF_API_TOKEN"))

        }




        return &schema.Provider{

            //set environment variables 
            Schema: map[string]*schema.Schema{
                "api_token": {
                    Type:        schema.TypeString,
                    Optional:    true,
                    DefaultFunc: schema.EnvDefaultFunc("ENF_API_TOKEN", nil),
                    Description: "Token from authenticating with dev.xaptum.io",
                },
                "base_url": {
                    Type:        schema.TypeString,
                    Optional:    true,
                    DefaultFunc: schema.EnvDefaultFunc("ENF_API_URL", nil),
                    Description: "Base URL for authentication",
                },
            },

            //ConfigureFunc: providerConfigure,

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

// func providerConfigure(d *schema.ResourceData) (interface{}, error) {
//     config := Config{
//         ApiToken:  d.Get("api_key").(string),
//         BaseURL: d.Get("base_url").(string),
//     }
//     return &config, nil
// }
