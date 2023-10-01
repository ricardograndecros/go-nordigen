package go_nordigen

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
)

const AccountsUrl string = "accounts/"
const TransactionsUrl string = "transactions/"

type Account struct {
	Id            string `json:"id"`
	Created       string `json:"created"`
	LastAccessed  string `json:"last_accessed"`
	Iban          string `json:"iban"`
	InstitutionId string `json:"institution_id"`
	Status        string `json:"status"`
	OwnerName     string `json:"owner_name"`
}

type Transaction struct {
	TransactionID                     string            `json:"internalTransactionId,omitempty"`
	DebtorName                        string            `json:"debtorName,omitempty"`
	DebtorAccount                     DebtorAccount     `json:"debtorAccount,omitempty"`
	TransactionAmount                 TransactionAmount `json:"transactionAmount"`
	BankTransactionCode               string            `json:"bankTransactionCode,omitempty"`
	BookingDate                       string            `json:"bookingDate,omitempty"` // Assuming this is a string representation of a date
	ValueDate                         string            `json:"valueDate"`             // Assuming this is a string representation of a date
	RemittanceInformationUnstructured string            `json:"remittanceInformationUnstructured,omitempty"`
}

type DebtorAccount struct {
	IBAN string `json:"iban"`
}

type TransactionAmount struct {
	Currency string `json:"currency"`
	Amount   string `json:"amount"` // Assuming this is a string representation of the amount
}

type Transactions struct {
	Booked  []Transaction `json:"booked"`
	Pending []Transaction `json:"pending"`
}

type TransactionsResponse struct {
	Transactions Transactions `json:"transactions"`
}

func (c Client) GetAccountInfo(accountId string) (*Account, error) {
	req, err := http.NewRequest("GET", BaseUrl+AccountsUrl+accountId+"/", bytes.NewBuffer([]byte("")))
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
		return nil, errors.New("Error while fetching account info. " + string(rune(resp.StatusCode)))
	}

	account := Account{}
	err = json.Unmarshal(body, &account)
	if err != nil {
		log.Print("Error unmarshal")
		return nil, err
	}

	return &account, nil
}

func (c Client) GetAccountTransactions(accountId, date_from, date_to string) (*Transactions, error) {
	endpoint := BaseUrl + AccountsUrl + accountId + "/" + TransactionsUrl

	req, err := http.NewRequest("GET", endpoint, bytes.NewBuffer([]byte("")))
	if err != nil {
		return nil, err
	}

	q := url.Values{}
	if date_from != "" {
		q.Add("date_from", date_from)
	}
	if date_to != "" {
		q.Add("date_to", date_to)
	}
	req.URL.RawQuery = q.Encode()

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
		return nil, errors.New("Error while fetching account transactions. " + string(rune(resp.StatusCode)))
	}

	var response TransactionsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Print("Error unmarshal")
		return nil, err
	}

	return &response.Transactions, nil
}
