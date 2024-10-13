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

type FilialAddressesHandler struct {
	FilialAddressesUsecase domains.FilialAddressesDomainUsecase
}

func NewFilialAddressesHandler(filialUsecase domains.FilialAddressesDomainUsecase) *FilialAddressesHandler {
	return &FilialAddressesHandler{
		FilialAddressesUsecase: filialUsecase,
	}
}

func (h *FilialAddressesHandler) FindAll(ctx echo.Context) error {
	filials, statusCode, err := h.FilialAddressesUsecase.FindAll()
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Filial addresses fetched successfully", responses.ToArrayOfFilialAddressesResponse(filials))
}

func (h *FilialAddressesHandler) Save(ctx echo.Context) error {
	var createFilialAddressRequest requests.CreateFilialAddressRequest

	if err := helpers.BindAndValidate(ctx, &createFilialAddressRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	filialAddress := createFilialAddressRequest.ToDomain()
	statusCode, err := h.FilialAddressesUsecase.Save(filialAddress)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Filial address saved successfully", nil)
}

func (h *FilialAddressesHandler) Update(ctx echo.Context) error {
	var updateFilialAddressRequest requests.UpdateFilialAddressRequest

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	if err := helpers.BindAndValidate(ctx, &updateFilialAddressRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	personalAddress := updateFilialAddressRequest.ToDomain()
	statusCode, err := h.FilialAddressesUsecase.Update(personalAddress, id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Filial address updated successfully", nil)
}

func (h *FilialAddressesHandler) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	statusCode, err := h.FilialAddressesUsecase.Delete(id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Filial address deleted successfully", nil)
}
