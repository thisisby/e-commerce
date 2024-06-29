package mobizon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const defaultClientTimeout = 10 * time.Second

type Client struct {
	baseUrl string
	apiKey  string
}

func NewClient(baseUrl, apiKey string) Client {
	return Client{
		baseUrl,
		apiKey,
	}
}

type SendSmsRequest struct {
	Recipient string `json:"recipient"`
	Text      string `json:"text"`
}

func (c *Client) SendSms(phone, code string) error {
	if c.baseUrl == "" || c.apiKey == "" {
		return fmt.Errorf("mobizon: baseUrl and apiKey are required")
	}

	// remove leading '+' Symbol from phone number
	phone = phone[1:]

	apiEndpoint := fmt.Sprintf("%s/Message/SendSmsMessage?apiKey=%s", c.baseUrl, c.apiKey)
	text := fmt.Sprintf("Ваш код подтверждения на интернет магазине ga_market: %s", code)

	requestBody, err := json.Marshal(&SendSmsRequest{
		Recipient: phone,
		Text:      text,
	})
	if err != nil {
		return fmt.Errorf("mobizon: failed to marshal request body: %w", err)
	}

	request, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("mobizon: failed to create request: %w", err)
	}
	request.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{Timeout: defaultClientTimeout}

	response, err := httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("mobizon: failed to send request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("mobizon: unexpected status code: %d", response.StatusCode)
	}

	return nil
}
