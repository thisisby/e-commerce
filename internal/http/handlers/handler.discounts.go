package handlers

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/http/datatransfers/requests"
	"ga_marketplace/pkg/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type DiscountHandler struct {
	discountUsecase domains.DiscountsUsecase
}

func NewDiscountHandler(discountUsecase domains.DiscountsUsecase) DiscountHandler {
	return DiscountHandler{
		discountUsecase: discountUsecase,
	}
}

func (d *DiscountHandler) Save(ctx echo.Context) error {
	var discountCreateRequest requests.DiscountCreateRequest

	err := helpers.BindAndValidate(ctx, &discountCreateRequest)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	discountDomain := discountCreateRequest.ToDomain()
	statusCode, err := d.discountUsecase.Save(discountDomain)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Discount saved successfully", nil)
}

func (d *DiscountHandler) DeleteByProductId(ctx echo.Context) error {
	id := ctx.Param("product_id")

	productIdInt, err := strconv.Atoi(id)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid discount id")
	}

	statusCode, err := d.discountUsecase.DeleteByProductId(productIdInt)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Discount deleted successfully", nil)
}
