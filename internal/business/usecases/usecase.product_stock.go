package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/constants"
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

func (p *productStockUsecase) Update(productStock domains.ProductStockDomain, transactionId string) (int, error) {
	err := p.productStockRepo.Update(productStock, transactionId)
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

func (p *productStockUsecase) FindStockItem(transactionId string, productId string) (domains.ProductStockItemDomain, int, error) {
	productStockItem, err := p.productStockRepo.FindStockItem(transactionId, productId)
	if err != nil {
		return productStockItem, http.StatusInternalServerError, err
	}

	return productStockItem, http.StatusOK, nil
}

func (p *productStockUsecase) UpdateProductStockItem(item domains.ProductStockItemDomain, id string, id2 string) (int, error) {
	err := p.productStockRepo.UpdateProductStockItem(item, id, id2)
	if err != nil {
		if errors.Is(err, constants.ErrForeignKeyViolation) {
			return http.StatusNotFound, errors.New("product code not found")
		}
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
