package one_c

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
)

type Client struct {
	BaseUrl  string
	Username string
	Password string
}

func NewClient(baseUrl, username, password string) Client {
	return Client{
		BaseUrl:  baseUrl,
		Username: username,
		Password: password,
	}
}

type Customer struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

func (c *Client) CreateCustomerRequest(customer Customer) error {
	endpoint := c.BaseUrl + "/Royal_Skin_Prog/hs/api/v1/customer"

	body, err := json.Marshal(customer)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		slog.Error("Failed to create request: Customer create", err)
		return err
	}

	req.SetBasicAuth(c.Username, c.Password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Failed to send request: Customer create", err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		slog.Error("Failed to send request: Customer create", err)
		return err
	}

	return nil
}
