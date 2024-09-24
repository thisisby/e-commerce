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

type AttributesHandler struct {
	AttributesUsecase domains.AttributesUsecase
}

func NewAttributesHandler(au domains.AttributesUsecase) *AttributesHandler {
	return &AttributesHandler{
		AttributesUsecase: au,
	}
}

func (h *AttributesHandler) FindAll(ctx echo.Context) error {
	attributes, statusCode, err := h.AttributesUsecase.FindAll()
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Attributes fetched successfully", responses.ToArrayOfAttributesResponse(attributes))
}

func (h *AttributesHandler) Save(ctx echo.Context) error {
	var createAttributesRequest requests.CreateAttributeRequest

	if err := helpers.BindAndValidate(ctx, &createAttributesRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	attributes := createAttributesRequest.ToDomain()
	statusCode, err := h.AttributesUsecase.Save(attributes)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Attributes saved successfully", nil)
}

func (h *AttributesHandler) Update(ctx echo.Context) error {
	var updateAttributesRequest requests.UpdateAttributeRequest
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	if err := helpers.BindAndValidate(ctx, &updateAttributesRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	attributes, statusCode, err := h.AttributesUsecase.FindById(id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	if updateAttributesRequest.Name != nil {
		attributes.Name = *updateAttributesRequest.Name
	}
	if updateAttributesRequest.CharacteristicId != nil {
		attributes.CharacteristicsId = *updateAttributesRequest.CharacteristicId
	}

	statusCode, err = h.AttributesUsecase.Update(attributes)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Attributes updated successfully", nil)
}

func (h *AttributesHandler) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	statusCode, err := h.AttributesUsecase.Delete(id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Attributes deleted successfully", nil)
}

func (h *AttributesHandler) FindAllByCharacteristicsId(ctx echo.Context) error {
	characteristicsId, err := strconv.Atoi(ctx.Param("characteristicId"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	attributes, statusCode, err := h.AttributesUsecase.FindAllByCharacteristicsId(characteristicsId)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Attributes fetched successfully", responses.ToArrayOfAttributesResponse(attributes))
}
