package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"net/http"
)

type servicesUsecase struct {
	servicesRepo domains.ServicesRepository
}

func NewServicesUsecase(servicesRepo domains.ServicesRepository) domains.ServicesUsecase {
	return &servicesUsecase{
		servicesRepo: servicesRepo,
	}
}

func (s *servicesUsecase) FindAll() ([]domains.ServicesDomain, int, error) {
	services, err := s.servicesRepo.FindAll()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(services) == 0 {
		return nil, http.StatusNotFound, errors.New("services not found")
	}

	return services, http.StatusOK, nil
}

func (s *servicesUsecase) Save(domain domains.ServicesDomain) (int, error) {
	err := s.servicesRepo.Save(domain)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (s *servicesUsecase) Update(domain domains.ServicesDomain) (int, error) {
	err := s.servicesRepo.Update(domain)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (s *servicesUsecase) Delete(id int) (int, error) {
	err := s.servicesRepo.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
