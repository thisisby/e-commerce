package usecases

import (
	"ga_marketplace/internal/business/domains"
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

	return products, http.StatusOK, nil
}
