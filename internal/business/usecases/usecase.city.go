package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"net/http"
)

type citiesUsecase struct {
	cityRepo domains.CitiesRepository
}

func NewCitiesUsecase(cityRepo domains.CitiesRepository) domains.CitiesUsecase {
	return &citiesUsecase{
		cityRepo: cityRepo,
	}
}

func (c *citiesUsecase) FindAll() ([]domains.CityDomain, int, error) {
	cities, err := c.cityRepo.FindAll()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(cities) == 0 {
		return nil, http.StatusNotFound, errors.New("cities not found")
	}

	return cities, http.StatusOK, nil
}

func (c *citiesUsecase) FindById(id int) (domains.CityDomain, int, error) {
	city, err := c.cityRepo.FindById(id)
	if err != nil {
		return domains.CityDomain{}, http.StatusInternalServerError, err
	}

	return city, http.StatusOK, nil
}

func (c *citiesUsecase) Save(city domains.CityDomain) (int, error) {
	err := c.cityRepo.Save(city)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (c *citiesUsecase) Update(city domains.CityDomain) (int, error) {
	err := c.cityRepo.Update(city)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (c *citiesUsecase) Delete(id int) (int, error) {
	err := c.cityRepo.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
