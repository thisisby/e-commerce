package usecases

import (
	"ga_marketplace/internal/business/domains"
	"net/http"
)

type productStockUsecase struct {
	productStockRepo domains.ProductStockRepository
}

func NewProductStockUsecase(productStockRepo domains.ProductStockRepository) domains.ProductStockUsecase {
	return &productStockUsecase{
		productStockRepo: productStockRepo,
	}
}

func (p *productStockUsecase) Save(productStock domains.ProductStockDomain) (int, error) {
	err := p.productStockRepo.Save(productStock)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (p *productStockUsecase) Update(productStock domains.ProductStockDomain) (int, error) {
	err := p.productStockRepo.Update(productStock)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (p *productStockUsecase) FindById(id string) (domains.ProductStockDomain, int, error) {
	productStock, err := p.productStockRepo.FindById(id)
	if err != nil {
		return productStock, http.StatusInternalServerError, err
	}

	return productStock, http.StatusOK, nil
}
