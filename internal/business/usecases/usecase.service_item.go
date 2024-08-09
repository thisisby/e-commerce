package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"net/http"
)

type serviceItemUsecase struct {
	serviceItemRepo domains.ServiceItemRepository
}

func NewServiceItemUsecase(serviceItemRepo domains.ServiceItemRepository) domains.ServiceItemUsecase {
	return &serviceItemUsecase{
		serviceItemRepo: serviceItemRepo,
	}
}

func (s *serviceItemUsecase) FindAll() ([]domains.ServiceItemDomain, int, error) {
	serviceItems, err := s.serviceItemRepo.FindAll()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(serviceItems) == 0 {
		return nil, http.StatusNotFound, errors.New("service items not found")
	}

	return serviceItems, http.StatusOK, nil
}

func (s *serviceItemUsecase) FindById(id int) (domains.ServiceItemDomain, int, error) {
	serviceItem, err := s.serviceItemRepo.FindById(id)
	if err != nil {
		return domains.ServiceItemDomain{}, http.StatusInternalServerError, err
	}

	if serviceItem.Id == 0 {
		return domains.ServiceItemDomain{}, http.StatusNotFound, errors.New("service item not found")
	}

	return serviceItem, http.StatusOK, nil
}

func (s *serviceItemUsecase) FindBySubServiceId(subServiceId int) ([]domains.ServiceItemDomain, int, error) {
	serviceItems, err := s.serviceItemRepo.FindBySubServiceId(subServiceId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(serviceItems) == 0 {
		return nil, http.StatusNotFound, errors.New("service items not found")
	}

	return serviceItems, http.StatusOK, nil
}

func (s *serviceItemUsecase) Update(serviceItem domains.ServiceItemDomain) (int, error) {
	err := s.serviceItemRepo.Update(serviceItem)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (s *serviceItemUsecase) Delete(id int) (int, error) {
	err := s.serviceItemRepo.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (s *serviceItemUsecase) Save(serviceItem domains.ServiceItemDomain) (int, error) {
	err := s.serviceItemRepo.Save(serviceItem)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}
