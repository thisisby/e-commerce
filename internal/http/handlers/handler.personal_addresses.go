package handlers

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/constants"
	"ga_marketplace/internal/http/datatransfers/requests"
	"ga_marketplace/internal/http/datatransfers/responses"
	"ga_marketplace/pkg/helpers"
	"ga_marketplace/pkg/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type PersonalAddressesHandler struct {
	PersonalAddressesUsecase domains.PersonalAddressesUsecase
}

func NewPersonalAddressesHandler(personalAddressesUsecase domains.PersonalAddressesUsecase) *PersonalAddressesHandler {
	return &PersonalAddressesHandler{
		PersonalAddressesUsecase: personalAddressesUsecase,
	}
}

func (h *PersonalAddressesHandler) FindAll(ctx echo.Context) error {
	personalAddresses, statusCode, err := h.PersonalAddressesUsecase.FindAll()
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Personal addresses fetched successfully", responses.ToArrayOfPersonalAddressesResponse(personalAddresses))
}

func (h *PersonalAddressesHandler) FindByUserId(ctx echo.Context) error {
	jwtClaims := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims)
	userId := jwtClaims.UserId

	personalAddresses, statusCode, err := h.PersonalAddressesUsecase.FindByUserId(userId)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Personal addresses fetched successfully", responses.ToArrayOfPersonalAddressesResponse(personalAddresses))
}

func (h *PersonalAddressesHandler) Save(ctx echo.Context) error {
	var createPersonalAddressRequest requests.CreatePersonalAddressRequest
	jwtClaims := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims)

	if err := helpers.BindAndValidate(ctx, &createPersonalAddressRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	personalAddress := createPersonalAddressRequest.ToDomain()
	personalAddress.UserId = jwtClaims.UserId
	statusCode, err := h.PersonalAddressesUsecase.Save(personalAddress)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Personal address saved successfully", nil)
}

func (h *PersonalAddressesHandler) Update(ctx echo.Context) error {
	var updatePersonalAddressRequest requests.UpdatePersonalAddressRequest
	jwtClaims := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims)
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	if err := helpers.BindAndValidate(ctx, &updatePersonalAddressRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	personalAddress := updatePersonalAddressRequest.ToDomain()
	personalAddress.UserId = jwtClaims.UserId
	statusCode, err := h.PersonalAddressesUsecase.Update(personalAddress, id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Personal address updated successfully", nil)
}

func (h *PersonalAddressesHandler) Delete(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid ID")
	}

	statusCode, err := h.PersonalAddressesUsecase.Delete(id)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Personal address deleted successfully", nil)
}
