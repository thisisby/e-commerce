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

type CountriesHandler struct {
	countriesUsecase domains.CountriesUsecase
}

func NewCountriesHandler(countriesUsecase domains.CountriesUsecase) CountriesHandler {
	return CountriesHandler{
		countriesUsecase: countriesUsecase,
	}
}

func (c *CountriesHandler) FindAll(ctx echo.Context) error {
	countries, statusCode, err := c.countriesUsecase.FindAll()
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Countries fetched successfully", responses.ToArrayOfCountryResponse(countries))
}

func (c *CountriesHandler) FindById(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	country, statusCode, err := c.countriesUsecase.FindById(id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Country fetched successfully", responses.FromCountryDomain(&country))
}

func (c *CountriesHandler) Save(ctx echo.Context) error {
	var createCountryRequest requests.CreateCountryRequest

	if err := helpers.BindAndValidate(ctx, &createCountryRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	country := createCountryRequest.ToDomain()
	statusCode, err := c.countriesUsecase.Save(country)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Country saved successfully", nil)
}

func (c *CountriesHandler) Update(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	var updateCountryRequest requests.UpdateCountryRequest

	if err := helpers.BindAndValidate(ctx, &updateCountryRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	country := updateCountryRequest.ToDomain()

	country.Id = id
	statusCode, err := c.countriesUsecase.Update(country)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Country updated successfully", nil)
}

func (c *CountriesHandler) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	statusCode, err := c.countriesUsecase.Delete(id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Country deleted successfully", nil)
}
