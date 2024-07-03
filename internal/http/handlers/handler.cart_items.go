package handlers

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/constants"
	"ga_marketplace/internal/datasources/caches"
	"ga_marketplace/internal/http/datatransfers/requests"
	"ga_marketplace/internal/http/datatransfers/responses"
	"ga_marketplace/pkg/helpers"
	"ga_marketplace/pkg/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type CartsHandler struct {
	cartUsecase domains.CartUsecase
	redisCache  caches.RedisCache
}

func NewCartsHandler(cartUsecase domains.CartUsecase, redisCache caches.RedisCache) CartsHandler {
	return CartsHandler{
		cartUsecase: cartUsecase,
		redisCache:  redisCache,
	}
}

func (c *CartsHandler) GetCartItemsByUserId(ctx echo.Context) error {
	userId := ctx.Param("id")

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid user id")
	}
	carts, statusCode, err := c.cartUsecase.FindByUserId(userIdInt)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Carts fetched successfully", responses.ToArrayOfCartItemsResponse(carts))
}

func (c *CartsHandler) SaveCartItem(ctx echo.Context) error {
	var cartCreateRequest requests.CartCreateRequest
	jwtClaims := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims)

	err := helpers.BindAndValidate(ctx, &cartCreateRequest)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	cartDomain := cartCreateRequest.ToDomain()
	cartDomain.UserId = jwtClaims.UserId
	statusCode, err := c.cartUsecase.Save(cartDomain)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Cart saved successfully", nil)
}

func (c *CartsHandler) DeleteCartItem(ctx echo.Context) error {
	cartId := ctx.Param("id")
	jwtClaims := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims)

	cartIdInt, err := strconv.Atoi(cartId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid cart id")
	}

	statusCode, err := c.cartUsecase.Delete(cartIdInt, jwtClaims.UserId)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Cart deleted successfully", nil)
}

func (c *CartsHandler) FindAll(ctx echo.Context) error {
	jwtClaims := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims)

	carts, statusCode, err := c.cartUsecase.FindAll(jwtClaims.UserId, jwtClaims.IsAdmin)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Carts fetched successfully", responses.ToArrayOfCartItemsResponse(carts))
}

func (c *CartsHandler) UpdateCartItem(ctx echo.Context) error {
	var cartUpdateRequest requests.CartUpdateRequest
	jwtClaims := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims)

	err := helpers.BindAndValidate(ctx, &cartUpdateRequest)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	cartId := ctx.Param("id")
	cartIdInt, err := strconv.Atoi(cartId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid cart id")
	}

	cartDomain := cartUpdateRequest.ToDomain()
	statusCode, err := c.cartUsecase.Update(cartIdInt, jwtClaims.UserId, cartDomain)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Cart updated successfully", nil)
}
