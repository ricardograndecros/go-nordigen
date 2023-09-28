package go_nordigen

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const UrlInstitutions string = "institutions/"

type Institution struct {
	Id                   string   `json:"id"`
	Name                 string   `json:"name"`
	Bic                  string   `json:"bic"`
	TransactionTotalDays string   `json:"transaction_total_days"`
	Countries            []string `json:"countries"`
	Logo                 string   `json:"logo"`
}

func (c Client) GetSupportedInstitutions(countryCode string) ([]Institution, error) {
	params := url.Values{
		"country": {countryCode},
	}

	req, err := http.NewRequest("GET", BaseUrl+UrlInstitutions+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}

	// Set authentication headers or credentials here
	req.Header.Add("Authorization", "Bearer "+c.token.Access)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var institutions []Institution
	err = json.Unmarshal(body, &institutions)
	if err != nil {
		return nil, err
	}

	return institutions, nil

}

func (c Client) GetInstitutionById(id string) (*Institution, error) {

	req, err := http.NewRequest("GET", BaseUrl+UrlInstitutions+id, nil)
	if err != nil {
		return nil, err
	}

	// Set authentication headers or credentials here
	req.Header.Add("Authorization", "Bearer "+c.token.Access)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var institution Institution
	err = json.Unmarshal(body, &institution)
	if err != nil {
		return nil, err
	}

	return &institution, nil

}
