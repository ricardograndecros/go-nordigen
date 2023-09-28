package go_nordigen

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Token struct {
	Access         string `json:"access"`
	AccessExpires  int    `json:"access_expires"`
	Refresh        string `json:"refresh"`
	RefreshExpires int    `json:"refresh_expires"`
}

type Secret struct {
	SecretId  string `json:"secret_id"`
	SecretKey string `json:"secret_key"`
}

func (c Client) newToken(secretID, secretKey string) (*Token, error) {
	payload, err := json.Marshal(Secret{
		SecretId:  secretID,
		SecretKey: secretKey,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(BaseUrl+"token/new/", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	t := Token{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil

}

func (c Client) refreshToken(refreshToken string) (*Token, error) {
	payload, err := json.Marshal(map[string]string{
		"refresh": refreshToken,
	})
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(BaseUrl+"token/refresh/", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	t := Token{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil

}
