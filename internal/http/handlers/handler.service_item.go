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

type ServiceItemHandler struct {
	serviceItemUsecase domains.ServiceItemUsecase
}

func NewServiceItemHandler(serviceItemUsecase domains.ServiceItemUsecase) *ServiceItemHandler {
	return &ServiceItemHandler{
		serviceItemUsecase: serviceItemUsecase,
	}
}

func (h *ServiceItemHandler) FindAll(ctx echo.Context) error {
	serviceItems, statusCode, err := h.serviceItemUsecase.FindAll()
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Service items fetched successfully", responses.ToArrayOfServiceItem(serviceItems))
}

func (h *ServiceItemHandler) Save(ctx echo.Context) error {
	var createServiceItemRequest requests.ServiceItemCreateRequest

	if err := helpers.BindAndValidate(ctx, &createServiceItemRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	serviceItem := createServiceItemRequest.ToDomain()
	statusCode, err := h.serviceItemUsecase.Save(serviceItem)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Service item saved successfully", nil)
}

func (h *ServiceItemHandler) Update(ctx echo.Context) error {
	var updateServiceItemRequest requests.ServiceItemUpdateRequest
	serviceId := ctx.Param("id")

	serviceIdInt, err := strconv.Atoi(serviceId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid service item id")
	}

	if err := helpers.BindAndValidate(ctx, &updateServiceItemRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	serviceDomain, statusCode, err := h.serviceItemUsecase.FindById(serviceIdInt)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	if updateServiceItemRequest.Title != nil {
		serviceDomain.Title = *updateServiceItemRequest.Title
	}
	if updateServiceItemRequest.Duration != nil {
		serviceDomain.Duration = *updateServiceItemRequest.Duration
	}
	if updateServiceItemRequest.Description != nil {
		serviceDomain.Description = *updateServiceItemRequest.Description
	}
	if updateServiceItemRequest.Price != nil {
		serviceDomain.Price = *updateServiceItemRequest.Price
	}
	if updateServiceItemRequest.Subservice_id != nil {
		serviceDomain.SubServiceId = *updateServiceItemRequest.Subservice_id
	}

	statusCode, err = h.serviceItemUsecase.Update(serviceDomain)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Service item updated successfully", nil)
}

func (h *ServiceItemHandler) Delete(ctx echo.Context) error {
	serviceId := ctx.Param("id")

	serviceIdInt, err := strconv.Atoi(serviceId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid service item id")
	}

	statusCode, err := h.serviceItemUsecase.Delete(serviceIdInt)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Service item deleted successfully", nil)
}

func (h *ServiceItemHandler) FindBySubserviceId(ctx echo.Context) error {
	subserviceId := ctx.Param("subservice_id")

	subserviceIdInt, err := strconv.Atoi(subserviceId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid subservice id")
	}

	serviceItems, statusCode, err := h.serviceItemUsecase.FindBySubServiceId(subserviceIdInt)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Service items fetched successfully", responses.ToArrayOfServiceItem(serviceItems))
}
