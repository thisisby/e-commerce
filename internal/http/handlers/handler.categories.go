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

type CategoriesHandler struct {
	categoriesUsecase domains.CategoriesUsecase
}

func NewCategoriesHandler(categoriesUsecase domains.CategoriesUsecase) CategoriesHandler {
	return CategoriesHandler{
		categoriesUsecase: categoriesUsecase,
	}
}

func (h *CategoriesHandler) FindAll(ctx echo.Context) error {
	categories, statusCode, err := h.categoriesUsecase.FindAll()
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Categories fetched successfully", responses.ToArrayOfCategoryResponse(categories))
}

func (h *CategoriesHandler) Save(ctx echo.Context) error {
	var createCategoryRequest requests.CreateCategoryRequest

	if err := helpers.BindAndValidate(ctx, &createCategoryRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	category := createCategoryRequest.ToDomain()
	statusCode, err := h.categoriesUsecase.Save(category)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Category saved successfully", nil)
}

func (h *CategoriesHandler) Update(ctx echo.Context) error {
	var updateCategoryRequest requests.UpdateCategoryRequest
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	if err := helpers.BindAndValidate(ctx, &updateCategoryRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	category := updateCategoryRequest.ToDomain()
	category.Id = id
	statusCode, err := h.categoriesUsecase.Update(category)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Category updated successfully", nil)
}

func (h *CategoriesHandler) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	statusCode, err := h.categoriesUsecase.Delete(id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Category deleted successfully", nil)
}
