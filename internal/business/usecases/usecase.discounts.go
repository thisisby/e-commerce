package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/constants"
	"net/http"
)

type discountsUsecase struct {
	discountRepo domains.DiscountsRepository
}

func NewDiscountsUsecase(discountRepo domains.DiscountsRepository) domains.DiscountsUsecase {
	return &discountsUsecase{
		discountRepo: discountRepo,
	}
}

func (d *discountsUsecase) Save(discount *domains.DiscountsDomain) (statusCode int, err error) {
	discountExists, err := d.discountRepo.FindByProductId(discount.ProductId)
	if err != nil {
		if !errors.Is(err, constants.ErrRowNotFound) {
			return http.StatusInternalServerError, err
		}
	}

	if discountExists != nil {
		return http.StatusBadRequest, constants.ErrDiscountExists
	}

	err = d.discountRepo.Save(discount)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (d *discountsUsecase) DeleteByProductId(id int) (statusCode int, err error) {
	err = d.discountRepo.DeleteByProductId(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
