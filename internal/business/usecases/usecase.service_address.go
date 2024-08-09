package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"net/http"
)

type serviceAddressUsecase struct {
	serviceAddressRepo domains.ServiceAddressRepository
}

func NewServiceAddressUsecase(serviceAddressRepo domains.ServiceAddressRepository) domains.ServiceAddressUsecase {
	return &serviceAddressUsecase{
		serviceAddressRepo: serviceAddressRepo,
	}
}

func (s *serviceAddressUsecase) FindAll() ([]domains.ServiceAddressDomain, int, error) {
	serviceAddress, err := s.serviceAddressRepo.FindAll()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(serviceAddress) == 0 {
		return nil, http.StatusNotFound, errors.New("service address not found")
	}

	return serviceAddress, http.StatusOK, nil
}

func (s *serviceAddressUsecase) FindById(id int) (domains.ServiceAddressDomain, int, error) {
	serviceAddress, err := s.serviceAddressRepo.FindById(id)
	if err != nil {
		return domains.ServiceAddressDomain{}, http.StatusInternalServerError, err
	}

	if serviceAddress.Id == 0 {
		return domains.ServiceAddressDomain{}, http.StatusNotFound, errors.New("service address not found")
	}

	return serviceAddress, http.StatusOK, nil
}

func (s *serviceAddressUsecase) Save(serviceAddress domains.ServiceAddressDomain) (int, error) {
	err := s.serviceAddressRepo.Save(serviceAddress)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (s *serviceAddressUsecase) Update(serviceAddress domains.ServiceAddressDomain) (int, error) {
	err := s.serviceAddressRepo.Update(serviceAddress)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (s *serviceAddressUsecase) Delete(id int) (int, error) {
	err := s.serviceAddressRepo.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (s *serviceAddressUsecase) FindAllByCityId(cityId int) ([]domains.ServiceAddressDomain, int, error) {
	serviceAddress, err := s.serviceAddressRepo.FindAllByCityId(cityId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(serviceAddress) == 0 {
		return nil, http.StatusNotFound, errors.New("service address not found")
	}

	return serviceAddress, http.StatusOK, nil
}
