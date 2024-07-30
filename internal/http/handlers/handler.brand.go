package handlers

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/http/datatransfers/requests"
	"ga_marketplace/internal/http/datatransfers/responses"
	"ga_marketplace/pkg/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type BrandHandler struct {
	brandUsecase domains.BrandsUsecase
}

func NewBrandHandler(brandUsecase domains.BrandsUsecase) BrandHandler {
	return BrandHandler{
		brandUsecase: brandUsecase,
	}
}

func (h *BrandHandler) FindAll(ctx echo.Context) error {
	brands, statusCode, err := h.brandUsecase.FindAll()
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Brands fetched successfully", responses.ToArrayOfBrandResponse(brands))
}

func (h *BrandHandler) Save(ctx echo.Context) error {
	var createBrandRequest requests.CreateBrandRequest

	if err := helpers.BindAndValidate(ctx, &createBrandRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	brand := createBrandRequest.ToDomain()
	statusCode, err := h.brandUsecase.Save(brand)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Brand saved successfully", nil)
}

func (h *BrandHandler) Update(ctx echo.Context) error {
	var updateBrandRequest requests.UpdateBrandRequest
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	if err := helpers.BindAndValidate(ctx, &updateBrandRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	brand := updateBrandRequest.ToDomain()
	brand.Id = id
	statusCode, err := h.brandUsecase.Update(brand)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Brand updated successfully", nil)
}

func (h *BrandHandler) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	statusCode, err := h.brandUsecase.Delete(id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Brand deleted successfully", nil)
}
