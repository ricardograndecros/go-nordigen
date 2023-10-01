package go_nordigen

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

const RequisitionUrl = "requisitions/"

type Requisition struct {
	Id                string    `json:"id"`
	InstitutionID     string    `json:"institution_id"`
	Created           time.Time `json:"created"`
	Redirect          string    `json:"redirect"`
	Status            string    `json:"status"`
	Agreement         string    `json:"agreement"`
	Reference         string    `json:"reference"`
	Accounts          []string  `json:"accounts"`
	Language          string    `json:"user_language"`
	Link              string    `json:"link"`
	Ssn               string    `json:"ssn"`
	AccountSelection  bool      `json:"account_selection"`
	RedirectImmediate bool      `json:"redirect_immediate"`
}

func (c Client) NewRequisition(r Requisition) (*Requisition, error) {
	payload, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", BaseUrl+RequisitionUrl, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token.Access)

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

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		return nil, errors.New("Error while creating the requisition. " + string(rune(resp.StatusCode)))
	}

	requisition := Requisition{}
	err = json.Unmarshal(body, &requisition)
	if err != nil {
		log.Print("Error unmarshal")
		return nil, err
	}

	return &requisition, nil
}

func (c Client) GetRequisitionsById(id string) (*Requisition, error) {
	req, err := http.NewRequest("GET", BaseUrl+RequisitionUrl+id+"/", bytes.NewBuffer([]byte("")))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.token.Access)

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

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Error while looking for the requisitions. " + string(rune(resp.StatusCode)))
	}

	requisition := Requisition{}
	err = json.Unmarshal(body, &requisition)
	if err != nil {
		log.Print("Error unmarshal")
		return nil, err
	}

	return &requisition, nil
}
