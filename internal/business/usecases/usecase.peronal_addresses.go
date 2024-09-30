package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"net/http"
)

type personalAddressesUsecase struct {
	personalAddressesRepo domains.PersonalAddressesRepository
}

func NewPersonalAddressesUsecase(personalAddressesRepo domains.PersonalAddressesRepository) domains.PersonalAddressesUsecase {
	return &personalAddressesUsecase{
		personalAddressesRepo: personalAddressesRepo,
	}
}

func (p *personalAddressesUsecase) FindAll() ([]domains.PersonalAddressesDomain, int, error) {
	personalAddresses, err := p.personalAddressesRepo.FindAll()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(personalAddresses) == 0 {
		return nil, http.StatusNotFound, errors.New("personal addresses not found")
	}

	return personalAddresses, http.StatusOK, nil
}

func (p *personalAddressesUsecase) FindByUserId(userId int) ([]domains.PersonalAddressesDomain, int, error) {
	personalAddresses, err := p.personalAddressesRepo.FindByUserId(userId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(personalAddresses) == 0 {
		return nil, http.StatusNotFound, errors.New("personal addresses not found")
	}

	return personalAddresses, http.StatusOK, nil
}

func (p *personalAddressesUsecase) Save(domain domains.PersonalAddressesDomain) (int, error) {
	err := p.personalAddressesRepo.Save(domain)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (p *personalAddressesUsecase) Update(domain domains.PersonalAddressesDomain, id int) (int, error) {
	err := p.personalAddressesRepo.Update(domain, id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (p *personalAddressesUsecase) Delete(id int) (int, error) {
	err := p.personalAddressesRepo.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
