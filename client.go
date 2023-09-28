package go_nordigen

import (
	"net/http"
	"time"
)

const BaseUrl = "https://bankaccountdata.gocardless.com/api/v2/"

type Client struct {
	client       *http.Client
	secretID     string
	clientSecret string
	token        *Token
	expiration   time.Time
}

func NewClient(secretID, secretKey string) (*Client, error) {

	c := Client{client: &http.Client{},
		secretID:     secretID,
		clientSecret: secretKey}

	token, err := c.newToken(c.secretID, c.clientSecret)
	if err != nil {
		return nil, err
	}
	c.token = token
	c.expiration = time.Now().Add(time.Duration(token.AccessExpires) * time.Second)

	return &c, nil
}
