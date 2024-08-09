package requests

import "ga_marketplace/internal/business/domains"

type ServiceAddressCreateRequest struct {
	CityId  int    `json:"city_id" validate:"required"`
	Address string `json:"address" validate:"required"`
}

type ServiceAddressUpdateRequest struct {
	CityId  *int    `json:"city_id"`
	Address *string `json:"address"`
}

func (s *ServiceAddressCreateRequest) ToDomain() *domains.ServiceAddressDomain {
	return &domains.ServiceAddressDomain{
		CityId:  s.CityId,
		Address: s.Address,
	}
}
