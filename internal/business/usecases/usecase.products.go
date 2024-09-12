package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/constants"
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

func (p *productsUsecase) FindAllForMe(id int, filter domains.ProductFilter) ([]domains.ProductDomain, int, error) {
	products, err := p.productsRepo.FindAllForMe(id, filter)
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

func (p *productsUsecase) FindAllForMeByBrandId(id int, brandId int) ([]domains.ProductDomain, int, error) {
	products, err := p.productsRepo.FindAllForMeByBrandId(id, brandId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(products) == 0 {
		return nil, http.StatusNotFound, errors.New("products not found")
	}

	return products, http.StatusOK, nil
}

func (p *productsUsecase) SaveFrom1c(product *domains.ProductDomainV2) (int, error) {
	product.BrandId = 1
	product.SubcategoryId = 1
	product.Ingredients = " "
	product.Image = " "
	err := p.productsRepo.SaveFrom1c(product)
	if err != nil {
		if errors.Is(err, constants.ErrRowExists) {
			return http.StatusConflict, err
		}
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (p *productsUsecase) UpdateFrom1c(code string, product *domains.ProductDomain) (int, error) {
	err := p.productsRepo.UpdateFrom1c(code, product)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (p *productsUsecase) FindByCode(code string) (*domains.ProductDomain, int, error) {
	product, err := p.productsRepo.FindByCode(code)
	if err != nil {
		if errors.Is(err, constants.ErrRowNotFound) {
			return nil, http.StatusNotFound, errors.New("product not found")
		}
		return nil, http.StatusInternalServerError, err
	}

	if product == nil {
		return nil, http.StatusNotFound, errors.New("product not found")
	}

	return product, http.StatusOK, nil
}

func (p *productsUsecase) FindByIdForUser(id int, userId int) (*domains.ProductDomain, int, error) {
	product, err := p.productsRepo.FindByIdForUser(id, userId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if product == nil {
		return nil, http.StatusNotFound, errors.New("product not found")
	}

	return product, http.StatusOK, nil
}

func (p *productsUsecase) FindAll(filter domains.ProductFilter) ([]domains.ProductDomain, int, error) {
	products, total, err := p.productsRepo.FindAll(filter)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(products) == 0 {
		return nil, http.StatusNotFound, errors.New("products not found")
	}

	return products, total, nil
}
