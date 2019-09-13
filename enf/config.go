package enf

import (
	"context"
	"net/http"

	"github.com/xaptum/go-enf/enf"
)

type Config struct {
	Username  string
	Password  string
	DomainURL string
}

type EnfClient struct {
	ApiToken   string
	DomainURL  string
	HTTPClient *http.Client
	DomainID   int64
	Client     *enf.Client
}

func (c *Config) Client() (interface{}, error) {
	client, err := enf.NewClient(c.DomainURL, nil)
	if err != nil {
		return nil, err
	}

	authReq := &enf.AuthRequest{Username: &c.Username, Password: &c.Password}
	auth, _, err := client.Auth.Authenticate(context.Background(), authReq)
	if err != nil {
		return nil, err
	}

	enfClient := &EnfClient{
		ApiToken:   *auth.Token,
		DomainID:   *auth.DomainID,
		DomainURL:  c.DomainURL,
		HTTPClient: &http.Client{},
		Client:     client,
	}

	client.ApiToken = *auth.Token

	return enfClient, nil
}
