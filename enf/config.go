package enf

import (
	"context"
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

	Client *Client
}

func (c *Config) Client() (interface{}, error) {
	client, err := NewClient(c.DomainURL, nil)
	if err != nil {
		return nil, err
	}

	authReq := &AuthRequest{Username: &c.Username, Password: &c.Password}
	auth, _, err := client.Auth.Authenticate(context.Background(), authReq)
	if err != nil {
		return nil, err
	}

	enfClient := &EnfClient{
		ApiToken:   *auth.Token,
		DomainURL:  c.DomainURL,
		HTTPClient: &http.Client{},

		Client: client,
	}

	return enfClient, nil
}
