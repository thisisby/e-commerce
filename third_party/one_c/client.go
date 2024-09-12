package one_c

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"
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

type ProductStock struct {
	CustomerId      string     `json:"customer_id"`
	TransactionId   string     `json:"transaction_id_mp"`
	Active          bool       `json:"active"`
	TransactionDate time.Time  `json:"transaction_date"`
	Products        []Products `json:"products"`
}

type Products struct {
	Quantity  int     `json:"quantity"`
	Amount    float64 `json:"amount"`
	ProductId string  `json:"product_id"`
}

func (c *Client) CreateProductStockRequest(productStock ProductStock) error {
	endpoint := c.BaseUrl + "/Royal_Skin_Prog/hs/api/v1/sales"
	slog.Info("url: ", c.BaseUrl)
	slog.Info(endpoint)
	body, err := json.Marshal(productStock)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	if err != nil {
		slog.Error("Failed to create request: Product stock create", err)
		return err
	}

	req.SetBasicAuth(c.Username, c.Password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Failed to send request: Product stock create", err)
		return err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Failed to read response body", err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		slog.Error("Failed to send request: Product stock create", err)
		slog.Error("Response Body: %s\n", string(bodyBytes))
		return err
	}

	return nil
}
