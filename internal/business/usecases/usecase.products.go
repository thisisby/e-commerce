package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/pkg/helpers"
	"net/http"
)

type productsUsecase struct {
	productsRepo domains.ProductsRepository
}

func NewProductsUsecase(productsRepo domains.ProductsRepository) domains.ProductsUsecase {
	return &productsUsecase{
		productsRepo: productsRepo,
	}
}

func (p *productsUsecase) Save(product *domains.ProductDomain) (int, error) {
	err := p.productsRepo.Save(product)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (p *productsUsecase) FindAllForMe(id int) ([]domains.ProductDomain, int, error) {
	products, err := p.productsRepo.FindAllForMe(id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(products) == 0 {
		return nil, http.StatusNotFound, errors.New("products not found")
	}

	return products, http.StatusOK, nil
}

func (p *productsUsecase) UpdateById(inDom domains.ProductDomain) (int, error) {
	inDom.UpdatedAt = helpers.GetCurrentTime()
	err := p.productsRepo.UpdateById(inDom)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (p *productsUsecase) FindById(id int) (*domains.ProductDomain, int, error) {
	product, err := p.productsRepo.FindById(id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if product == nil {
		return nil, http.StatusNotFound, errors.New("product not found")
	}

	return product, http.StatusOK, nil
}

func (p *productsUsecase) FindAllForMeBySubcategoryId(id int, subcategoryId int) ([]domains.ProductDomain, int, error) {
	products, err := p.productsRepo.FindAllForMeBySubcategoryId(id, subcategoryId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(products) == 0 {
		return nil, http.StatusNotFound, errors.New("products not found")
	}

	return products, http.StatusOK, nil
}
