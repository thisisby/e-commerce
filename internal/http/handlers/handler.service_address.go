package handlers

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/http/datatransfers/requests"
	"ga_marketplace/pkg/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ServiceAddressHandler struct {
	serviceAddressUsecase domains.ServiceAddressUsecase
}

func NewServiceAddressHandler(serviceAddressUsecase domains.ServiceAddressUsecase) *ServiceAddressHandler {
	return &ServiceAddressHandler{
		serviceAddressUsecase: serviceAddressUsecase,
	}
}

func (h *ServiceAddressHandler) FindAll(ctx echo.Context) error {
	serviceAddress, statusCode, err := h.serviceAddressUsecase.FindAll()
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Service Address fetched successfully", serviceAddress)
}

func (h *ServiceAddressHandler) Save(ctx echo.Context) error {
	var createServiceAddressRequest requests.ServiceAddressCreateRequest

	if err := helpers.BindAndValidate(ctx, &createServiceAddressRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	serviceAddress := createServiceAddressRequest.ToDomain()
	statusCode, err := h.serviceAddressUsecase.Save(*serviceAddress)

	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Service Address saved successfully", nil)
}

func (h *ServiceAddressHandler) Update(ctx echo.Context) error {
	var updateServiceAddressRequest requests.ServiceAddressUpdateRequest
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	if err := helpers.BindAndValidate(ctx, &updateServiceAddressRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	serviceAddresDomain, statusCode, err := h.serviceAddressUsecase.FindById(id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	if updateServiceAddressRequest.Address != nil {
		serviceAddresDomain.Address = *updateServiceAddressRequest.Address
	}
	if updateServiceAddressRequest.CityId != nil {
		serviceAddresDomain.CityId = *updateServiceAddressRequest.CityId
	}

	statusCode, err = h.serviceAddressUsecase.Update(serviceAddresDomain)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Service Address updated successfully", nil)
}

func (h *ServiceAddressHandler) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	statusCode, err := h.serviceAddressUsecase.Delete(id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Service Address deleted successfully", nil)
}

func (h *ServiceAddressHandler) FindByCityId(ctx echo.Context) error {
	cityId, err := strconv.Atoi(ctx.Param("city_id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid City ID")
	}

	serviceAddress, statusCode, err := h.serviceAddressUsecase.FindAllByCityId(cityId)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Service Address fetched successfully", serviceAddress)
}
