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

	statusParam := ctx.QueryParam("status")

	orders, statusCode, err := o.ordersUsecase.FindByUserId(jwtClaims.UserId, statusParam)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Orders found", responses.ToArrayOfOrdersResponse(orders))
}

func (o *OrdersHandler) Update(ctx echo.Context) error {
	var orderUpdateRequest requests.UpdateOrderRequest
	orderId := ctx.Param("id")

	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid order id")
	}
	err = helpers.BindAndValidate(ctx, &orderUpdateRequest)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	orderDomain, statusCode, err := o.ordersUsecase.FindById(orderIdInt)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	if orderUpdateRequest.Status != nil {
		orderDomain.Status = *orderUpdateRequest.Status
	}
	if orderUpdateRequest.Street != nil {
		orderDomain.Street = *orderUpdateRequest.Street
	}
	if orderUpdateRequest.Region != nil {
		orderDomain.Region = *orderUpdateRequest.Region
	}
	if orderUpdateRequest.Apartment != nil {
		orderDomain.Apartment = *orderUpdateRequest.Apartment
	}
	if orderUpdateRequest.CityId != nil {
		orderDomain.CityId = *orderUpdateRequest.CityId
	}

	statusCode, err = o.ordersUsecase.Update(orderDomain)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Order updated successfully", nil)
}