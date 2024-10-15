package cdek

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
	baseUrl      string
	grantType    string
	clientId     string
	clientSecret string
}

func NewClient(baseUrl, grantType, clientId, clientSecret string) *Client {
	return &Client{
		baseUrl:      baseUrl,
		grantType:    grantType,
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}

func (c *Client) authorize() (string, error) {
	url := fmt.Sprintf("%s/v2/oauth/token?grant_type=%s&client_id=%s&client_secret=%s",
		c.baseUrl, c.grantType, c.clientId, c.clientSecret)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(nil))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("authorization failed: %s", body)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	token, ok := result["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("access_token not found in response")
	}

	return token, nil
}

func (c *Client) GetCityCode(city string) (int, error) {
	token, err := c.authorize()
	if err != nil {
		return 0, fmt.Errorf("authorization failed: %w", err)
	}

	url := fmt.Sprintf("%s/v2/location/suggest/cities?name=%s", c.baseUrl, city)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return 0, fmt.Errorf("failed to get city code: %s", body)
	}

	var cities []struct {
		CityUUID string `json:"city_uuid"`
		Code     int    `json:"code"`
		FullName string `json:"full_name"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&cities); err != nil {
		return 0, fmt.Errorf("failed to decode response: %w", err)
	}

	if len(cities) == 0 {
		return 0, fmt.Errorf("no city found with name: %s", city)
	}

	return cities[0].Code, nil
}

type Package struct {
	Weight int `json:"weight"`
}

type Location struct {
	Code    int    `json:"code"`
	Address string `json:"address"`
	City    string `json:"city"`
}

type DeliveryRequest struct {
	Type         int       `json:"type"`
	Date         string    `json:"date"`
	Currency     int       `json:"currency"`
	Lang         string    `json:"lang"`
	FromLocation Location  `json:"from_location"`
	ToLocation   Location  `json:"to_location"`
	Packages     []Package `json:"packages"`
}

type Tariff struct {
	TariffCode        int     `json:"tariff_code"`
	TariffName        string  `json:"tariff_name"`
	TariffDescription string  `json:"tariff_description"`
	DeliveryMode      int     `json:"delivery_mode"`
	DeliverySum       float64 `json:"delivery_sum"`
	PeriodMin         int     `json:"period_min"`
	PeriodMax         int     `json:"period_max"`
	CalendarMin       int     `json:"calendar_min"`
	CalendarMax       int     `json:"calendar_max"`
}

type CalculatorResponse struct {
	TariffCodes []Tariff `json:"tariff_codes"`
}

func (c *Client) CalculateDelivery(fromCode, toCode int, deliveryCalculatorRequest requests.DeliveryCalculatorRequest) (Tariff, error) {
	token, err := c.authorize()
	if err != nil {
		return Tariff{}, fmt.Errorf("authorization failed: %w", err)
	}

	location, err := time.LoadLocation("Asia/Almaty")
	if err != nil {
		return Tariff{}, fmt.Errorf("failed to load timezone: %w", err)
	}
	currentTime := time.Now().In(location)
	formattedDate := currentTime.Format("2006-01-02T15:04:05-0700")

	requestBody := DeliveryRequest{
		Type:     1,
		Date:     formattedDate,
		Currency: 2,
		Lang:     "rus",
		FromLocation: Location{
			Code:    fromCode,
			Address: deliveryCalculatorRequest.FromLocation.ToAddressString(),
			City:    "Астана",
		},
		ToLocation: Location{
			Code:    toCode,
			Address: deliveryCalculatorRequest.ToLocation.ToAddressString(),
			City:    "Астана",
		},
		Packages: []Package{
			{Weight: deliveryCalculatorRequest.Weight},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return Tariff{}, fmt.Errorf("failed to marshal request body: %w", err)
	}

	url := fmt.Sprintf("%s/v2/calculator/tarifflist", c.baseUrl)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return Tariff{}, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Tariff{}, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return Tariff{}, fmt.Errorf("failed to calculate delivery: %s", body)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Tariff{}, fmt.Errorf("failed to read response: %w", err)
	}

	var calcResponse CalculatorResponse
	if err := json.Unmarshal(body, &calcResponse); err != nil {
		return Tariff{}, fmt.Errorf("failed to parse response: %w", err)
	}

	for _, tariff := range calcResponse.TariffCodes {
		if tariff.TariffCode == 139 {
			return tariff, nil
		}
	}

	return Tariff{}, fmt.Errorf("tariff with code 139 not found")
}
