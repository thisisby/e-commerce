package one_c

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	endpoint := c.BaseUrl + "/Royal_Skin/hs/api/v1/customer"

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
	endpoint := c.BaseUrl + "/Royal_Skin/hs/api/v1/sales"
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
		return fmt.Errorf("product stock create failed: status code %d, response body: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}

func (c *Client) CheckProductStockRequest(productId string, quantity int) (bool, error) {
	endpoint := fmt.Sprintf("%s/Royal_Skin/hs/api/v1/balance?product_id_1c=%s", c.BaseUrl, productId)
	slog.Info("Checking product stock at: ", endpoint)

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		slog.Error("Failed to create request: Product stock check", err)
		return false, err
	}

	req.SetBasicAuth(c.Username, c.Password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Failed to send request: Product stock check", err)
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		slog.Error("Failed to fetch product stock", fmt.Errorf("status: %d, body: %s", resp.StatusCode, string(bodyBytes)))
		return false, fmt.Errorf("failed to fetch product stock: status code %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Failed to read response body", err)
		return false, err
	}

	var products []struct {
		ProductID string `json:"product_id_1c"`
		Quantity  int    `json:"quantity"`
	}

	if err := json.Unmarshal(bodyBytes, &products); err != nil {
		slog.Error("Failed to parse response body", err)
		return false, err
	}

	if len(products) == 0 {
		slog.Error("No product found with the given product ID")
		return false, fmt.Errorf("product with ID %s not found", productId)
	}

	availableQuantity := products[0].Quantity
	slog.Info(fmt.Sprintf("Available quantity: %d, Requested quantity: %d", availableQuantity, quantity))

	if availableQuantity < quantity {
		slog.Error("Insufficient stock")
		return false, fmt.Errorf("insufficient stock: available %d, required %d", availableQuantity, quantity)
	}

	slog.Info("Sufficient stock available")
	return true, nil
}
