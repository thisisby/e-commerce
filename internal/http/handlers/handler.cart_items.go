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

type CartItemsHandler struct {
	cartUsecase domains.CartUsecase
	redisCache  caches.RedisCache
}

func NewCartsHandler(cartUsecase domains.CartUsecase, redisCache caches.RedisCache) CartItemsHandler {
	return CartItemsHandler{
		cartUsecase: cartUsecase,
		redisCache:  redisCache,
	}
}

func (c *CartItemsHandler) GetAllCartItemsByUserId(ctx echo.Context) error {
	userId := ctx.Param("id")

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid user id")
	}
	carts, statusCode, err := c.cartUsecase.FindAllByUserId(userIdInt)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Carts fetched successfully", responses.ToArrayOfCartItemsResponse(carts))
}

func (c *CartItemsHandler) SaveToMyCartItems(ctx echo.Context) error {
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

func (c *CartItemsHandler) DeleteMyCartItem(ctx echo.Context) error {
	cartId := ctx.Param("id")
	jwtClaims := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims)

	cartIdInt, err := strconv.Atoi(cartId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid cart id")
	}

	statusCode, err := c.cartUsecase.DeleteByIdAndUserId(cartIdInt, jwtClaims.UserId)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Cart deleted successfully", nil)
}

func (c *CartItemsHandler) GetAllMyCartItems(ctx echo.Context) error {
	jwtClaims := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims)

	carts, statusCode, err := c.cartUsecase.FindAllByUserId(jwtClaims.UserId)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	totalAmount, statusCode, err := c.cartUsecase.FindTotalAmountByUserId(jwtClaims.UserId)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	totalAmountResponse := responses.CartItemTotalAmountResponse{
		TotalAmount:   totalAmount.TotalAmount,
		TotalDiscount: totalAmount.TotalDiscount,
	}

	outResponse := map[string]interface{}{
		"total_amount": totalAmountResponse,
		"cart_items":   responses.ToArrayOfCartItemsResponse(carts),
	}

	return NewSuccessResponse(ctx, statusCode, "Carts fetched successfully", outResponse)
}

func (c *CartItemsHandler) UpdateMyCartItem(ctx echo.Context) error {
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
	statusCode, err := c.cartUsecase.UpdateByIdAndUserId(cartIdInt, jwtClaims.UserId, cartDomain)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Cart updated successfully", nil)
}

func (c *CartItemsHandler) DeleteAllMyCartItems(ctx echo.Context) error {
	jwtClaims := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims)

	statusCode, err := c.cartUsecase.DeleteAllByUserId(jwtClaims.UserId)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Carts deleted successfully", nil)
}

func (c *CartItemsHandler) GetAllCartItemsForUser(ctx echo.Context) error {
	userId := ctx.Param("user-id")

	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid user id")
	}
	carts, statusCode, err := c.cartUsecase.FindAllByUserId(userIdInt)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Carts fetched successfully", responses.ToArrayOfCartItemsAdminResponse(carts))
}

func (c *CartItemsHandler) DeleteMyCartItemsByIds(ctx echo.Context) error {
	var cartIdsRequest requests.CartDeleteRequest
	jwtClaims := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims)

	err := helpers.BindAndValidate(ctx, &cartIdsRequest)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	statusCode, err := c.cartUsecase.DeleteByIdsAndUserId(jwtClaims.UserId, cartIdsRequest.Ids)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Carts deleted successfully", nil)
}
