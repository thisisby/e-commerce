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

type SubcategoriesHandler struct {
	subcategoriesUsecase domains.SubcategoriesUsecase
}

func NewSubcategoriesHandler(subcategoriesUsecase domains.SubcategoriesUsecase) SubcategoriesHandler {
	return SubcategoriesHandler{
		subcategoriesUsecase: subcategoriesUsecase,
	}
}

func (h *SubcategoriesHandler) FindAll(ctx echo.Context) error {
	subcategories, statusCode, err := h.subcategoriesUsecase.FindAll()
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Subcategories fetched successfully", responses.ToArrayOfSubcategoryResponse(subcategories))
}

func (h *SubcategoriesHandler) Save(ctx echo.Context) error {
	var createSubcategoryRequest requests.CreateSubcategoryRequest

	if err := helpers.BindAndValidate(ctx, &createSubcategoryRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	subcategory := createSubcategoryRequest.ToDomain()
	statusCode, err := h.subcategoriesUsecase.Save(subcategory)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Subcategory saved successfully", nil)
}

func (h *SubcategoriesHandler) Update(ctx echo.Context) error {
	var updateSubcategoryRequest requests.UpdateSubcategoryRequest
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	if err := helpers.BindAndValidate(ctx, &updateSubcategoryRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	subcategoryDomain, statusCode, err := h.subcategoriesUsecase.FindById(id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	if updateSubcategoryRequest.Name != nil {
		subcategoryDomain.Name = *updateSubcategoryRequest.Name
	}
	if updateSubcategoryRequest.CategoryId != nil {
		subcategoryDomain.CategoryId = *updateSubcategoryRequest.CategoryId
	}

	statusCode, err = h.subcategoriesUsecase.Update(subcategoryDomain)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Subcategory updated successfully", nil)
}

func (h *SubcategoriesHandler) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	statusCode, err := h.subcategoriesUsecase.Delete(id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Subcategory deleted successfully", nil)
}

func (h *SubcategoriesHandler) FindById(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	subcategory, statusCode, err := h.subcategoriesUsecase.FindById(id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Subcategory fetched successfully", responses.FromSubcategoryDomain(&subcategory))
}

func (h *SubcategoriesHandler) FindByCategoryId(ctx echo.Context) error {
	categoryId, err := strconv.Atoi(ctx.Param("category_id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid Category ID")
	}

	subcategories, statusCode, err := h.subcategoriesUsecase.FindAllByCategoryId(categoryId)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Subcategories fetched successfully", responses.ToArrayOfSubcategoryResponse(subcategories))
}
