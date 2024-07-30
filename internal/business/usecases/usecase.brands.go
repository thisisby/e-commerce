package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"net/http"
)

type brandsUsecase struct {
	brandsRepo domains.BrandsRepository
}

func NewBrandsUsecase(brandsRepo domains.BrandsRepository) domains.BrandsUsecase {
	return &brandsUsecase{
		brandsRepo: brandsRepo,
	}
}

func (b *brandsUsecase) FindAll() ([]domains.BrandsDomain, int, error) {
	brands, err := b.brandsRepo.FindAll()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(brands) == 0 {
		return nil, http.StatusNotFound, errors.New("brands not found")
	}

	return brands, http.StatusOK, nil
}

func (b *brandsUsecase) Save(domain domains.BrandsDomain) (int, error) {
	err := b.brandsRepo.Save(domain)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (b *brandsUsecase) Update(domain domains.BrandsDomain) (int, error) {
	err := b.brandsRepo.Update(domain)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (b *brandsUsecase) Delete(id int) (int, error) {
	err := b.brandsRepo.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
