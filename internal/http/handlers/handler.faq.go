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

type FaqHandler struct {
	FaqUsecase domains.FaqUsecase
}

func NewFaqHandler(faqUsecase domains.FaqUsecase) *FaqHandler {
	return &FaqHandler{
		FaqUsecase: faqUsecase,
	}
}

func (h *FaqHandler) FindAll(ctx echo.Context) error {
	filials, statusCode, err := h.FaqUsecase.FindAll()
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Faq addresses fetched successfully", responses.ToArrayOfFaqResponse(filials))
}

func (h *FaqHandler) Save(ctx echo.Context) error {
	var createFaqRequest requests.CreateFaqRequest

	if err := helpers.BindAndValidate(ctx, &createFaqRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	faq := createFaqRequest.ToDomain()
	statusCode, err := h.FaqUsecase.Save(faq)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Faq address saved successfully", nil)
}

func (h *FaqHandler) Update(ctx echo.Context) error {
	var updateFaqRequest requests.CreateFaqRequest

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	if err := helpers.BindAndValidate(ctx, &updateFaqRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	faq := updateFaqRequest.ToDomain()
	statusCode, err := h.FaqUsecase.Update(faq, id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "faq address updated successfully", nil)
}

func (h *FaqHandler) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	statusCode, err := h.FaqUsecase.Delete(id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Faq address deleted successfully", nil)
}
