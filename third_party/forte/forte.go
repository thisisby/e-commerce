package forte

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ga_marketplace/internal/http/datatransfers/requests"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	baseUrl  string
	username string
	password string
}

func NewClient(baseUrl, username, password string) *Client {
	return &Client{
		baseUrl:  baseUrl,
		username: username,
		password: password,
	}
}

func (c *Client) CreatePayment(paymentInfo requests.CreatePaymentRequest) (any, int, error) {
	url := fmt.Sprintf("%s/transactions/payments", c.baseUrl)

	fmt.Printf("Payment Request: %v\n", paymentInfo)

	payload := map[string]any{
		"request": paymentInfo,
	}

	fmt.Printf("Payment Request: %v\n", payload)

	requestBody, err := json.Marshal(payload)
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("failed to marshal payment request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(c.username, c.password)

	client := &http.Client{Timeout: 20 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 404, fmt.Errorf("failed to read response body: %w", err)
	}

	fmt.Printf("Payment Response: %s\n", string(body))

	if resp.StatusCode != http.StatusOK {
		return nil, http.StatusBadRequest, fmt.Errorf("received non-200 response: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	var result map[string]any
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, http.StatusBadRequest, fmt.Errorf("failed to decode response: %w", err)
	}

	return result, http.StatusOK, nil
}
