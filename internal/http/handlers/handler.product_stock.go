package handlers

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/http/datatransfers/requests"
	"ga_marketplace/pkg/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProductStockHandler struct {
	productStockUsecase domains.ProductStockUsecase
}

func NewProductStockHandler(productStockUsecase domains.ProductStockUsecase) ProductStockHandler {
	return ProductStockHandler{
		productStockUsecase: productStockUsecase,
	}
}

func (p *ProductStockHandler) Save(ctx echo.Context) error {
	var productStockCreateRequest requests.CreateProductStockRequest

	if err := helpers.BindAndValidate(ctx, &productStockCreateRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	productStockDomain := productStockCreateRequest.ToDomain()

	statusCode, err := p.productStockUsecase.Save(productStockDomain)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusCreated, "Product stock created successfully", nil)
}

func (p *ProductStockHandler) Update(ctx echo.Context) error {
	var productStockUpdateRequest requests.UpdateProductStockRequest

	if err := helpers.BindAndValidate(ctx, &productStockUpdateRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	id := ctx.Param("id")
	productStockDomain, statusCode, err := p.productStockUsecase.FindById(id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	if productStockUpdateRequest.CCode != nil {
		productStockDomain.CCode = *productStockUpdateRequest.CCode
	}
	if productStockUpdateRequest.Date != nil {
		productStockDomain.Date = *productStockUpdateRequest.Date
	}
	if productStockUpdateRequest.TransactionType != nil {
		productStockDomain.TransactionType = *productStockUpdateRequest.TransactionType
	}
	if productStockUpdateRequest.TransactionId != nil {
		productStockDomain.TransactionId = *productStockUpdateRequest.TransactionId
	}
	if productStockUpdateRequest.Quantity != nil {
		productStockDomain.Quantity = *productStockUpdateRequest.Quantity
	}
	if productStockUpdateRequest.TotalSum != nil {
		productStockDomain.TotalSum = *productStockUpdateRequest.TotalSum
	}
	if productStockUpdateRequest.TransactionStatus != nil {
		productStockDomain.TransactionStatus = *productStockUpdateRequest.TransactionStatus
	}

	statusCode, err = p.productStockUsecase.Update(productStockDomain)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Product stock updated successfully", nil)

}
