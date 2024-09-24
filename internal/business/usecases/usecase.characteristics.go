package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"net/http"
)

type characteristicsUsecase struct {
	characteristicsRepo domains.CharacteristicsRepository
}

func NewCharacteristicsUsecase(characteristicsRepo domains.CharacteristicsRepository) domains.CharacteristicsUsecase {
	return &characteristicsUsecase{
		characteristicsRepo: characteristicsRepo,
	}
}

func (c *characteristicsUsecase) FindAll() ([]domains.CharacteristicsDomain, int, error) {
	characteristics, err := c.characteristicsRepo.FindAll()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(characteristics) == 0 {
		return nil, http.StatusNotFound, errors.New("characteristics not found")
	}

	return characteristics, http.StatusOK, nil
}

func (c *characteristicsUsecase) Save(domain domains.CharacteristicsDomain) (int, error) {
	err := c.characteristicsRepo.Save(domain)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (c *characteristicsUsecase) Update(domain domains.CharacteristicsDomain) (int, error) {
	err := c.characteristicsRepo.Update(domain)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (c *characteristicsUsecase) Delete(id int) (int, error) {
	err := c.characteristicsRepo.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (c *characteristicsUsecase) FindAllBySubcategoryId(subcategoryId int) ([]domains.CharacteristicsDomain, int, error) {
	characteristics, err := c.characteristicsRepo.FindAllBySubcategoryId(subcategoryId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(characteristics) == 0 {
		return nil, http.StatusNotFound, errors.New("characteristics not found")
	}

	return characteristics, http.StatusOK, nil
}

func (c *characteristicsUsecase) FindById(id int) (domains.CharacteristicsDomain, int, error) {
	characteristic, err := c.characteristicsRepo.FindById(id)
	if err != nil {
		return domains.CharacteristicsDomain{}, http.StatusInternalServerError, err
	}

	if characteristic.Id == 0 {
		return domains.CharacteristicsDomain{}, http.StatusNotFound, errors.New("characteristic not found")
	}

	return characteristic, http.StatusOK, nil
}
