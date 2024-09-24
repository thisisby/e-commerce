package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"net/http"
)

type attributesUsecase struct {
	attributesRepo domains.AttributesRepository
}

func NewAttributesUsecase(attributesRepo domains.AttributesRepository) domains.AttributesUsecase {
	return &attributesUsecase{
		attributesRepo: attributesRepo,
	}
}

func (a *attributesUsecase) FindAll() ([]domains.AttributesDomain, int, error) {
	attributes, err := a.attributesRepo.FindAll()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(attributes) == 0 {
		return nil, http.StatusNotFound, errors.New("attributes not found")
	}

	return attributes, http.StatusOK, nil
}

func (a *attributesUsecase) Save(domain domains.AttributesDomain) (int, error) {
	err := a.attributesRepo.Save(domain)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (a *attributesUsecase) Update(domain domains.AttributesDomain) (int, error) {
	err := a.attributesRepo.Update(domain)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (a *attributesUsecase) Delete(id int) (int, error) {
	err := a.attributesRepo.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (a *attributesUsecase) FindAllByCharacteristicsId(characteristicsId int) ([]domains.AttributesDomain, int, error) {
	attributes, err := a.attributesRepo.FindAllByCharacteristicsId(characteristicsId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(attributes) == 0 {
		return nil, http.StatusNotFound, errors.New("attributes not found")
	}

	return attributes, http.StatusOK, nil
}

func (a *attributesUsecase) FindById(id int) (domains.AttributesDomain, int, error) {
	attributes, err := a.attributesRepo.FindById(id)
	if err != nil {
		return domains.AttributesDomain{}, http.StatusInternalServerError, err
	}

	return attributes, http.StatusOK, nil
}
