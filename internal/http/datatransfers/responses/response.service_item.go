package responses

import "ga_marketplace/internal/business/domains"

type ServiceItemResponse struct {
	Id           int     `json:"id"`
	Title        string  `json:"title"`
	Duration     int     `json:"duration"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	SubserviceId int     `json:"subservice_id"`
}

func FromServiceItemDomain(domain domains.ServiceItemDomain) *ServiceItemResponse {
	if domain.Id == 0 {
		return nil
	}
	return &ServiceItemResponse{
		Id:           domain.Id,
		Title:        domain.Title,
		Duration:     domain.Duration,
		Description:  domain.Description,
		Price:        domain.Price,
		SubserviceId: domain.SubServiceId,
	}
}

func ToArrayOfServiceItem(domain []domains.ServiceItemDomain) []ServiceItemResponse {
	var result []ServiceItemResponse
	for _, item := range domain {
		result = append(result, *FromServiceItemDomain(item))
	}
	return result
}
