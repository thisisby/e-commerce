package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"net/http"
)

type countriesUsecase struct {
	countriesRepo domains.CountriesRepository
}

func NewCountriesUsecase(countriesRepo domains.CountriesRepository) domains.CountriesUsecase {
	return &countriesUsecase{
		countriesRepo: countriesRepo,
	}
}

func (c *countriesUsecase) FindAll() ([]domains.CountryDomain, int, error) {
	countries, err := c.countriesRepo.FindAll()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(countries) == 0 {
		return nil, http.StatusNotFound, errors.New("countries not found")
	}

	return countries, http.StatusOK, nil
}

func (c *countriesUsecase) FindById(id int) (domains.CountryDomain, int, error) {
	country, err := c.countriesRepo.FindById(id)
	if err != nil {
		return domains.CountryDomain{}, http.StatusInternalServerError, err
	}

	if country.Id == 0 {
		return domains.CountryDomain{}, http.StatusNotFound, errors.New("country not found")
	}

	return country, http.StatusOK, nil
}

func (c *countriesUsecase) Save(country domains.CountryDomain) (int, error) {
	err := c.countriesRepo.Save(country)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (c *countriesUsecase) Update(country domains.CountryDomain) (int, error) {
	err := c.countriesRepo.Update(country)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (c *countriesUsecase) Delete(id int) (int, error) {
	err := c.countriesRepo.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
