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

type ContactHandler struct {
	contactUsecase domains.ContactUsecase
}

func NewContactHandler(contactUsecase domains.ContactUsecase) *ContactHandler {
	return &ContactHandler{
		contactUsecase: contactUsecase,
	}
}

func (c *ContactHandler) FindAll(ctx echo.Context) error {
	contacts, statusCode, err := c.contactUsecase.FindAll()
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Contacts fetched successfully", responses.ToArrayOfContactResponse(contacts))
}

func (c *ContactHandler) Save(ctx echo.Context) error {
	var createContactRequest requests.CreateContactRequest

	if err := helpers.BindAndValidate(ctx, &createContactRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	contact := createContactRequest.ToDomain()

	statusCode, err := c.contactUsecase.Save(contact)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Contact saved successfully", nil)
}

func (c *ContactHandler) Update(ctx echo.Context) error {
	var updateContactRequest requests.UpdateContactRequest
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	if err := helpers.BindAndValidate(ctx, &updateContactRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	contactDomain, statusCode, err := c.contactUsecase.FindById(id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	if updateContactRequest.Title != nil {
		contactDomain.Title = *updateContactRequest.Title
	}
	if updateContactRequest.Value != nil {
		contactDomain.Value = *updateContactRequest.Value
	}

	statusCode, err = c.contactUsecase.Update(contactDomain)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Contact updated successfully", nil)

}

func (c *ContactHandler) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	statusCode, err := c.contactUsecase.Delete(id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Contact deleted successfully", nil)
}
