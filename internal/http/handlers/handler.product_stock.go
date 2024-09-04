package handlers

import (
	"fmt"
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

	productStockDomain := requests.ConvertToProductStockDomain(productStockCreateRequest)

	fmt.Println("productStockDomain", productStockDomain)

	statusCode, err := p.productStockUsecase.Save(productStockDomain)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusCreated, "Product stock created successfully", nil)
}

func (p *ProductStockHandler) Update(ctx echo.Context) error {
	var productStockCreateRequest requests.UpdateProductStockRequest

	if err := helpers.BindAndValidate(ctx, &productStockCreateRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	transactionId := ctx.Param("transaction_id")

	productStock, statusCode, err := p.productStockUsecase.FindById(transactionId)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	if productStockCreateRequest.TransactionId != nil {
		productStock.TransactionId = *productStockCreateRequest.TransactionId
	}
	if productStockCreateRequest.Date != nil {
		productStock.Date = *productStockCreateRequest.Date
	}
	if productStockCreateRequest.CustomerId != nil {
		productStock.CustomerId = *productStockCreateRequest.CustomerId
	}
	if productStockCreateRequest.Active != nil {
		productStock.Active = *productStockCreateRequest.Active
	}

	statusCode, err = p.productStockUsecase.Update(productStock, transactionId)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Product stock updated successfully", nil)

}

func (p *ProductStockHandler) UpdateProductStockItem(ctx echo.Context) error {
	var productStockCreateRequest requests.UpdateProductStockItemRequest

	if err := helpers.BindAndValidate(ctx, &productStockCreateRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	transactionId := ctx.Param("transaction_id")
	productId := ctx.Param("product_id")

	productStockItem, statusCode, err := p.productStockUsecase.FindStockItem(transactionId, productId)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	if productStockCreateRequest.ProductCode != nil {
		productStockItem.ProductCode = *productStockCreateRequest.ProductCode
	}
	if productStockCreateRequest.Quantity != nil {
		productStockItem.Quantity = *productStockCreateRequest.Quantity
	}
	if productStockCreateRequest.Amount != nil {
		productStockItem.Amount = *productStockCreateRequest.Amount
	}
	if productStockCreateRequest.TransactionType != nil {
		productStockItem.TransactionType = *productStockCreateRequest.TransactionType
	}

	statusCode, err = p.productStockUsecase.UpdateProductStockItem(productStockItem, transactionId, productId)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Product stock item updated successfully", nil)
}
