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

type CitiesHandler struct {
	citiesUsecase domains.CitiesUsecase
}

func NewCitiesHandler(citiesUsecase domains.CitiesUsecase) CitiesHandler {
	return CitiesHandler{
		citiesUsecase: citiesUsecase,
	}
}

func (h *CitiesHandler) FindAll(ctx echo.Context) error {
	cities, statusCode, err := h.citiesUsecase.FindAll()
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Cities fetched successfully", responses.ToArrayOfCityResponse(cities))
}

func (h *CitiesHandler) FindById(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	city, statusCode, err := h.citiesUsecase.FindById(id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "City fetched successfully", responses.FromCityDomain(&city))
}

func (h *CitiesHandler) Save(ctx echo.Context) error {
	var createCityRequest requests.CreateCityRequest

	if err := helpers.BindAndValidate(ctx, &createCityRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	city := createCityRequest.ToDomain()
	statusCode, err := h.citiesUsecase.Save(city)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "City saved successfully", nil)
}

func (h *CitiesHandler) Update(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	var updateCityRequest requests.UpdateCityRequest
	if err := helpers.BindAndValidate(ctx, &updateCityRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	city := updateCityRequest.ToDomain()
	city.Id = id
	statusCode, err := h.citiesUsecase.Update(city)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "City updated successfully", nil)
}

func (h *CitiesHandler) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	statusCode, err := h.citiesUsecase.Delete(id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "City deleted successfully", nil)
}
