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

type ServicesHandler struct {
	serviceUsecase domains.ServicesUsecase
}

func NewServicesHandler(serviceUsecase domains.ServicesUsecase) *ServicesHandler {
	return &ServicesHandler{
		serviceUsecase: serviceUsecase,
	}
}

func (h *ServicesHandler) FindAll(ctx echo.Context) error {
	services, statusCode, err := h.serviceUsecase.FindAll()
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Services fetched successfully", responses.ToArrayOfServiceResponse(services))
}

func (h *ServicesHandler) Save(ctx echo.Context) error {
	var createServiceRequest requests.CreateServiceRequest

	if err := helpers.BindAndValidate(ctx, &createServiceRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	service := createServiceRequest.ToDomain()
	statusCode, err := h.serviceUsecase.Save(service)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Service saved successfully", nil)
}

func (h *ServicesHandler) Update(ctx echo.Context) error {
	var updateServiceRequest requests.UpdateServiceRequest
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	if err := helpers.BindAndValidate(ctx, &updateServiceRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	service := updateServiceRequest.ToDomain()
	service.Id = id
	statusCode, err := h.serviceUsecase.Update(service)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Service updated successfully", nil)
}

func (h *ServicesHandler) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	statusCode, err := h.serviceUsecase.Delete(id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Service deleted successfully", nil)
}
