package go_nordigen

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	url2 "net/url"
	"time"
)

const AgreementsUrl = "agreements/enduser/"

type EndUserAgreement struct {
	Id                 string    `json:"id"`
	CreatedAt          time.Time `json:"created"`
	MaxHistoricalDays  int       `json:"max_historical_days"`
	AccessValidForDays int       `json:"access_valid_for_days"`
	AccessScope        []string  `json:"access_scope"`
	Accepted           bool      `json:"accepted"`
	InstitutionId      string    `json:"institution_id"`
}

func (c Client) CreateUserAgreement(InstitutionId string) (*EndUserAgreement, error) {
	url, _ := url2.Parse(BaseUrl + AgreementsUrl)

	payload := map[string]interface{}{
		"institution_id":        InstitutionId,
		"max_historical_days":   "90",
		"access_valid_for_days": "30",
		"access_scope":          []string{"balances", "details", "transactions"},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token.Access) // Replace ACCESS_TOKEN with your actual access token

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.New("Error while creating the agreement. " + string(rune(resp.StatusCode)))
	}
	agreement := EndUserAgreement{}
	err = json.Unmarshal(body, &agreement)
	if err != nil {
		return nil, err
	}

	return &agreement, nil

}
