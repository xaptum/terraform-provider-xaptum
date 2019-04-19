package enf

import (
	"log"
    "net/http"
    "encoding/json"
    "bytes"
    "io/ioutil"
)


//Eventually, put these all in a separate file
//can also put ApiToken, BaseURL, HTTP client assignment functions in there too
//and then have Client() func below call them
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

////username,password,domain_url
type Config struct {
	Username string
	Password string
	DomainURL string 
}

type EnfClient struct {
    ApiToken  string
    DomainURL string
    HTTPClient *http.Client
}

//get token, base URL, and Client, when given username and password in Provider.go
func (c *Config) Client() (interface{}, error) {


		jsonData := map[string]string{"username": c.Username, "token": c.Password}
        jsonValue, _ := json.Marshal(jsonData)

	    domain_url := "https://dev.xaptum.io"
	    url := domain_url + "/api/xcr/v2/xauth"
	    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))

        req.Header.Add("content-type", "application/json")
		req.Header.Add("accept", "application/json")
        //response, err := http.Post(domain_url + "/api/xcr/v2/xauth", "application/json", bytes.NewBuffer(jsonValue))
		http_client := http.Client{}
        response, err := http_client.Do(req)


        if err != nil {
              log.Printf("The HTTP request failed with error %s\n", err) //TODO print with log.Printf with [DEBUG] etc
              return nil, err 
        } else {
                data_body, _ := ioutil.ReadAll(response.Body)

                var resp Response
                json.Unmarshal([]byte(data_body), &resp)

                log.Printf("Client response StatusCode is: ", response.StatusCode)
                log.Printf("Returned data_body is: ", string(data_body))
                log.Printf("Returned data is: ", (resp))
                log.Printf("Creds: ", resp.Data)
                log.Printf("Token:", resp.Data[0].Token)
                log.Printf("Pages is: ", resp.Page)

                //os.Setenv("ENF_API_TOKEN", resp.Data[0].Token)
                //fmt.Println("ENF_API_TOKEN:", os.Getenv("ENF_API_TOKEN"))


				


                client := &EnfClient{
                	ApiToken: resp.Data[0].Token,
                	DomainURL: domain_url,
                	HTTPClient: &http_client,
                }
				//client.ApiToken = resp.Data[0].Token
                //client.DomainURL = base_url
                //client.HTTPClient = &http_client
       			return client, nil
        }


}