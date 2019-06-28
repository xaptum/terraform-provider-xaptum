package enf

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Data []Data
	Page Pages
}

type Data struct {
	Username      string `json:"username"`
	Token         string `json:"token"`
	UserID        int    `json:"user_id"`
	Type          string `json:"type"`
	DomainID      int    `json:"domain_id"`
	DomainNetwork string `json:"domain_network"`
}

type Pages struct {
	Curr int `json:"curr"`
	Next int `json: "next"`
	Prev int `json: "prev"`
}

type Config struct {
	Username  string
	Password  string
	DomainURL string
}

type EnfClient struct {
	ApiToken   string
	DomainURL  string
	HTTPClient *http.Client
}

func (c *Config) Client() (interface{}, error) {

	jsonData := map[string]string{"username": c.Username, "token": c.Password}
	jsonValue, _ := json.Marshal(jsonData)

	domain_url := "https://dev.xaptum.io"
	url := domain_url + "/api/xcr/v2/xauth"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))

	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept", "application/json")
	http_client := http.Client{}
	response, err := http_client.Do(req)

	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
		return nil, err
	} else {
		data_body, _ := ioutil.ReadAll(response.Body)

		var resp Response
		json.Unmarshal([]byte(data_body), &resp)

		client := &EnfClient{
			ApiToken:   resp.Data[0].Token,
			DomainURL:  domain_url,
			HTTPClient: &http_client,
		}

		return client, nil
	}

}
