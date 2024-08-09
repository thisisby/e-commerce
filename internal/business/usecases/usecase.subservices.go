package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"net/http"
)

type subServicesUsecase struct {
	subServicesRepo domains.SubServicesRepository
}

func NewSubServicesUsecase(subServicesRepo domains.SubServicesRepository) domains.SubServicesUsecase {
	return &subServicesUsecase{
		subServicesRepo: subServicesRepo,
	}
}

func (s *subServicesUsecase) FindAll() ([]domains.SubServicesDomain, int, error) {
	subservices, err := s.subServicesRepo.FindAll()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(subservices) == 0 {
		return nil, http.StatusNotFound, errors.New("subservices not found")
	}

	return subservices, http.StatusOK, nil
}

func (s *subServicesUsecase) Save(domain domains.SubServicesDomain) (int, error) {
	err := s.subServicesRepo.Save(domain)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (s *subServicesUsecase) Update(domain domains.SubServicesDomain) (int, error) {
	err := s.subServicesRepo.Update(domain)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (s *subServicesUsecase) Delete(id int) (int, error) {
	err := s.subServicesRepo.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (s *subServicesUsecase) FindAllByServiceId(serviceId int) ([]domains.SubServicesDomain, int, error) {
	subServices, err := s.subServicesRepo.FindAllByServiceId(serviceId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(subServices) == 0 {
		return nil, http.StatusNotFound, errors.New("subservices not found")
	}

	return subServices, http.StatusOK, nil
}

func (s *subServicesUsecase) FindById(id int) (domains.SubServicesDomain, int, error) {
	subService, err := s.subServicesRepo.FindById(id)
	if err != nil {
		return domains.SubServicesDomain{}, http.StatusInternalServerError, err
	}

	return subService, http.StatusOK, nil
}
