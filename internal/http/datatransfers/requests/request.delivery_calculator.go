package requests

import (
	"fmt"
	"strings"
)

type DeliveryCalculatorLocation struct {
	Street    string `json:"street"`
	Region    string `json:"region"`
	Apartment string `json:"apartment"`
	StreetNum string `json:"street_num"`
	City      string `json:"city"`
}

func (l DeliveryCalculatorLocation) ToAddressString() string {
	var parts []string

	if l.Region != "" {
		parts = append(parts, l.Region)
	}
	if l.Street != "" {
		parts = append(parts, l.Street)
	}
	if l.StreetNum != "" {
		parts = append(parts, l.StreetNum)
	}
	if l.Apartment != "" {
		parts = append(parts, fmt.Sprintf("Apartment %s", l.Apartment))
	}

	return strings.Join(parts, ", ")
}

type DeliveryCalculatorRequest struct {
	FromLocation DeliveryCalculatorLocation `json:"from_location"`
	ToLocation   DeliveryCalculatorLocation `json:"to_location"`
	Weight       int                        `json:"weight"`
}
