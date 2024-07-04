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

type WishHandler struct {
	wishUsecase domains.WishUsecase
}

func NewWishHandler(wishUsecase domains.WishUsecase) WishHandler {
	return WishHandler{
		wishUsecase: wishUsecase,
	}
}

func (w *WishHandler) GetMyWishes(ctx echo.Context) error {
	jwtClaims := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims)

	wishes, statusCode, err := w.wishUsecase.FindByUserId(jwtClaims.UserId)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Wishes fetched successfully", responses.ToArrayOfWishResponse(wishes))
}

func (w *WishHandler) SaveToMyWishes(ctx echo.Context) error {
	var wishCreateRequest requests.WishCreateRequest
	jwtClaims := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims)

	err := helpers.BindAndValidate(ctx, &wishCreateRequest)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	wishDomain := wishCreateRequest.ToDomain()
	wishDomain.UserId = jwtClaims.UserId
	statusCode, err := w.wishUsecase.Save(wishDomain)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Wish saved successfully", nil)
}

func (w *WishHandler) DeleteMyWish(ctx echo.Context) error {
	id := ctx.Param("id")
	userId := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims).UserId

	wishIdInt, err := strconv.Atoi(id)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid wish id")
	}

	statusCode, err := w.wishUsecase.DeleteByIdAndUserId(wishIdInt, userId)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Wish deleted successfully", nil)
}
