package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"net/http"
)

type subcategoriesUsecase struct {
	subcategoriesRepo domains.SubcategoriesRepository
}

func NewSubcategoriesUsecase(subcategoriesRepo domains.SubcategoriesRepository) domains.SubcategoriesUsecase {
	return &subcategoriesUsecase{
		subcategoriesRepo: subcategoriesRepo,
	}
}

func (s *subcategoriesUsecase) FindAll() ([]domains.SubcategoriesDomain, int, error) {
	subcategories, err := s.subcategoriesRepo.FindAll()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(subcategories) == 0 {
		return nil, http.StatusNotFound, errors.New("subcategories not found")
	}

	return subcategories, http.StatusOK, nil
}

func (s *subcategoriesUsecase) FindAllByCategoryId(categoryId int) ([]domains.SubcategoriesDomain, int, error) {
	subcategories, err := s.subcategoriesRepo.FindAllByCategoryId(categoryId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(subcategories) == 0 {
		return nil, http.StatusNotFound, errors.New("subcategories not found")
	}

	return subcategories, http.StatusOK, nil
}

func (s *subcategoriesUsecase) Save(subcategoriesDomain domains.SubcategoriesDomain) (int, error) {
	err := s.subcategoriesRepo.Save(subcategoriesDomain)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (s *subcategoriesUsecase) Update(subcategoriesDomain domains.SubcategoriesDomain) (int, error) {
	err := s.subcategoriesRepo.Update(subcategoriesDomain)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (s *subcategoriesUsecase) Delete(id int) (int, error) {
	err := s.subcategoriesRepo.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (s *subcategoriesUsecase) FindById(id int) (domains.SubcategoriesDomain, int, error) {
	subcategory, err := s.subcategoriesRepo.FindById(id)
	if err != nil {
		return domains.SubcategoriesDomain{}, http.StatusInternalServerError, err
	}

	if subcategory.Id == 0 {
		return domains.SubcategoriesDomain{}, http.StatusNotFound, errors.New("subcategory not found")
	}

	return subcategory, http.StatusOK, nil
}
