package usecases

import (
	"ga_marketplace/internal/business/domains"
	"net/http"
)

type categoriesUsecase struct {
	categoriesRepo domains.CategoriesRepository
}

func NewCategoriesUsecase(categoriesRepo domains.CategoriesRepository) domains.CategoriesUsecase {
	return &categoriesUsecase{
		categoriesRepo: categoriesRepo,
	}
}

func (c *categoriesUsecase) FindAll() ([]domains.CategoriesDomain, int, error) {
	categories, err := c.categoriesRepo.FindAll()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return categories, http.StatusOK, nil
}

func (c *categoriesUsecase) Save(domain domains.CategoriesDomain) (int, error) {
	err := c.categoriesRepo.Save(domain)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (c *categoriesUsecase) Update(domain domains.CategoriesDomain) (int, error) {
	err := c.categoriesRepo.Update(domain)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (c *categoriesUsecase) Delete(id int) (int, error) {
	err := c.categoriesRepo.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
