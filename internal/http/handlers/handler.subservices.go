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

type SubServicesHandler struct {
	SubServicesUsecase domains.SubServicesUsecase
}

func NewSubServicesHandler(subServicesUsecase domains.SubServicesUsecase) *SubServicesHandler {
	return &SubServicesHandler{
		SubServicesUsecase: subServicesUsecase,
	}
}

func (h *SubServicesHandler) FindAll(ctx echo.Context) error {
	subServices, statusCode, err := h.SubServicesUsecase.FindAll()
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Subservices fetched successfully", responses.ToArrayOfSubServiceResponse(subServices))
}

func (h *SubServicesHandler) Save(ctx echo.Context) error {
	var createSubServiceRequest requests.CreateSubserviceRequest

	if err := helpers.BindAndValidate(ctx, &createSubServiceRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	subService := createSubServiceRequest.ToDomain()
	statusCode, err := h.SubServicesUsecase.Save(subService)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Subservice saved successfully", nil)
}

func (h *SubServicesHandler) Update(ctx echo.Context) error {
	var updateSubServiceRequest requests.UpdateSubserviceRequest

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	if err := helpers.BindAndValidate(ctx, &updateSubServiceRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	subServiceDomain, statusCode, err := h.SubServicesUsecase.FindById(id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	if updateSubServiceRequest.Name != nil {
		subServiceDomain.Name = *updateSubServiceRequest.Name
	}
	if updateSubServiceRequest.ServiceId != nil {
		subServiceDomain.ServiceId = *updateSubServiceRequest.ServiceId
	}

	statusCode, err = h.SubServicesUsecase.Update(subServiceDomain)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Subservice updated successfully", nil)
}

func (h *SubServicesHandler) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	statusCode, err := h.SubServicesUsecase.Delete(id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Subservice deleted successfully", nil)
}

func (h *SubServicesHandler) FindAllByServiceId(ctx echo.Context) error {
	serviceId, err := strconv.Atoi(ctx.Param("service_id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid Service ID")
	}

	subServices, statusCode, err := h.SubServicesUsecase.FindAllByServiceId(serviceId)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Subservices fetched successfully", responses.ToArrayOfSubServiceResponse(subServices))
}
