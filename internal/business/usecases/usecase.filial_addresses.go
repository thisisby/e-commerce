package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"net/http"
)

type filialAddressesUsecase struct {
	filialAddressesRepo domains.FilialAddressesRepository
}

func NewFilialAddressesUsecase(filialRepo domains.FilialAddressesRepository) domains.FilialAddressesDomainUsecase {
	return &filialAddressesUsecase{
		filialAddressesRepo: filialRepo,
	}
}

func (f *filialAddressesUsecase) FindAll() ([]domains.FilialAddressesDomain, int, error) {
	filialAddresses, err := f.filialAddressesRepo.FindAll()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(filialAddresses) == 0 {
		return nil, http.StatusNotFound, errors.New("filials not found")
	}

	return filialAddresses, http.StatusOK, nil
}

func (f *filialAddressesUsecase) FindByUserId(userId int) ([]domains.FilialAddressesDomain, int, error) {
	//TODO implement me
	return nil, http.StatusInternalServerError, errors.New("not used func")
}

func (f *filialAddressesUsecase) Save(domain domains.FilialAddressesDomain) (int, error) {
	err := f.filialAddressesRepo.Save(domain)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (f *filialAddressesUsecase) Update(domain domains.FilialAddressesDomain, id int) (int, error) {
	err := f.filialAddressesRepo.Update(domain, id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (f *filialAddressesUsecase) Delete(id int) (int, error) {
	err := f.filialAddressesRepo.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
