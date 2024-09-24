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

type CharacteristicsHandler struct {
	CharacteristicsUsecase domains.CharacteristicsUsecase
}

func NewCharacteristicsHandler(c domains.CharacteristicsUsecase) *CharacteristicsHandler {
	return &CharacteristicsHandler{
		CharacteristicsUsecase: c,
	}
}

func (h *CharacteristicsHandler) FindAll(ctx echo.Context) error {
	characteristics, status, err := h.CharacteristicsUsecase.FindAll()
	if err != nil {
		return NewErrorResponse(ctx, status, err.Error())
	}

	return NewSuccessResponse(ctx, status, "Characteristics fetched successfully", responses.ToArrayOfCharacteristicsResponse(characteristics))
}

func (h *CharacteristicsHandler) Save(ctx echo.Context) error {
	var createCharacteristicsRequest requests.CreateCharacteristicRequest

	if err := helpers.BindAndValidate(ctx, &createCharacteristicsRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	characteristics := createCharacteristicsRequest.ToDomain()
	status, err := h.CharacteristicsUsecase.Save(characteristics)
	if err != nil {
		return NewErrorResponse(ctx, status, err.Error())
	}

	return NewSuccessResponse(ctx, status, "Characteristics saved successfully", nil)
}

func (h *CharacteristicsHandler) Update(ctx echo.Context) error {
	var updateCharacteristicsRequest requests.UpdateCharacteristicRequest
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	if err := helpers.BindAndValidate(ctx, &updateCharacteristicsRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	characteristic, statusCode, err := h.CharacteristicsUsecase.FindById(id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	if updateCharacteristicsRequest.Name != nil {
		characteristic.Name = *updateCharacteristicsRequest.Name
	}
	if updateCharacteristicsRequest.SubcategoryId != nil {
		characteristic.SubcategoryId = *updateCharacteristicsRequest.SubcategoryId
	}

	status, err := h.CharacteristicsUsecase.Update(characteristic)
	if err != nil {
		return NewErrorResponse(ctx, status, err.Error())
	}

	return NewSuccessResponse(ctx, status, "Characteristics updated successfully", nil)
}

func (h *CharacteristicsHandler) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	status, err := h.CharacteristicsUsecase.Delete(id)
	if err != nil {
		return NewErrorResponse(ctx, status, err.Error())
	}

	return NewSuccessResponse(ctx, status, "Characteristics deleted successfully", nil)
}

func (h *CharacteristicsHandler) FindAllBySubcategoryId(ctx echo.Context) error {
	subcategoryId, err := strconv.Atoi(ctx.Param("subcategoryId"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	characteristics, status, err := h.CharacteristicsUsecase.FindAllBySubcategoryId(subcategoryId)
	if err != nil {
		return NewErrorResponse(ctx, status, err.Error())
	}

	return NewSuccessResponse(ctx, status, "Characteristics fetched successfully", responses.ToArrayOfCharacteristicsResponse(characteristics))
}
