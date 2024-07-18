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
)

type OrdersHandler struct {
	ordersUsecase    domains.OrdersUsecase
	cartItemsUsecase domains.CartUsecase
}

func NewOrdersHandler(ordersUsecase domains.OrdersUsecase, cartItemsUsecase domains.CartUsecase) OrdersHandler {
	return OrdersHandler{
		ordersUsecase:    ordersUsecase,
		cartItemsUsecase: cartItemsUsecase,
	}
}

func (o *OrdersHandler) Save(ctx echo.Context) error {
	var orderCreateRequest requests.CreateOrderRequest
	jwtClaims := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims)

	err := helpers.BindAndValidate(ctx, &orderCreateRequest)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	cartItems, statusCode, err := o.cartItemsUsecase.FindAllByUserId(jwtClaims.UserId)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}
	totalAmount, statusCode, err := o.cartItemsUsecase.FindTotalAmountByUserId(jwtClaims.UserId)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	orderDomain := orderCreateRequest.ToDomain()
	orderDomain.UserId = jwtClaims.UserId
	statusCode, err = o.ordersUsecase.Save(orderDomain, cartItems, *totalAmount)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Order saved successfully", nil)
}

func (o *OrdersHandler) FindMyOrders(ctx echo.Context) error {
	jwtClaims := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims)

	orders, statusCode, err := o.ordersUsecase.FindByUserId(jwtClaims.UserId)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Orders found", responses.ToArrayOfOrdersResponse(orders))
}
