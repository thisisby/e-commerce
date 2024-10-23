package handlers

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/constants"
	"ga_marketplace/internal/http/datatransfers/requests"
	"ga_marketplace/internal/http/datatransfers/responses"
	"ga_marketplace/pkg/helpers"
	"ga_marketplace/pkg/jwt"
	"ga_marketplace/third_party/forte"
	"ga_marketplace/third_party/one_c"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type OrdersHandler struct {
	ordersUsecase    domains.OrdersUsecase
	cartItemsUsecase domains.CartUsecase
	forteClient      *forte.Client
	oneC             *one_c.Client
}

func NewOrdersHandler(
	ordersUsecase domains.OrdersUsecase,
	cartItemsUsecase domains.CartUsecase,
	forteClient *forte.Client,
	oneC *one_c.Client) OrdersHandler {
	return OrdersHandler{
		ordersUsecase:    ordersUsecase,
		cartItemsUsecase: cartItemsUsecase,
		forteClient:      forteClient,
		oneC:             oneC,
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

	for _, cartItem := range cartItems {
		ok, err := o.oneC.CheckProductStockRequest(cartItem.Product.CCode, cartItem.Quantity)
		if err != nil {
			return NewErrorResponse(ctx, http.StatusBadRequest, "Failed to check product stock: "+err.Error())
		}

		if !ok {
			return NewErrorResponse(ctx, http.StatusBadRequest, "Product out of stock")
		}
	}

	totalAmount, statusCode, err := o.cartItemsUsecase.FindTotalAmountByUserId(jwtClaims.UserId)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	paymentRequest := requests.CreatePaymentRequest{
		Amount:         totalAmount.TotalAmount,
		Currency:       "KZT",
		Description:    orderCreateRequest.CreatePaymentRequest.Description,
		Language:       "ru",
		BillingAddress: orderCreateRequest.CreatePaymentRequest.BillingAddress,
		CreditCard:     orderCreateRequest.CreatePaymentRequest.CreditCard,
		Customer:       orderCreateRequest.CreatePaymentRequest.Customer,
	}

	paymentResponse, status, err := o.forteClient.CreatePayment(paymentRequest)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Payment failed: "+err.Error())
	}

	if status != http.StatusOK {
		return NewErrorResponse(ctx, status, "Payment failed")
	}

	orderDomain := orderCreateRequest.ToDomain()
	orderDomain.UserId = jwtClaims.UserId
	statusCode, err = o.ordersUsecase.Save(orderDomain, cartItems, *totalAmount)
	if err != nil {
		return NewSuccessResponse(ctx, statusCode, err.Error(), paymentResponse)
	}

	return NewSuccessResponse(ctx, statusCode, "Order saved successfully", paymentResponse)
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
	if orderUpdateRequest.StreetNum != nil {
		orderDomain.StreetNum = *orderUpdateRequest.StreetNum
	}
	if orderUpdateRequest.Email != nil {
		orderDomain.Email = *orderUpdateRequest.Email
	}
	if orderUpdateRequest.DeliveryMethod != nil {
		orderDomain.DeliveryMethod = *orderUpdateRequest.DeliveryMethod
	}

	statusCode, err = o.ordersUsecase.Update(orderDomain)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Order updated successfully", nil)
}

func (o *OrdersHandler) FindAll(ctx echo.Context) error {

	params := ctx.QueryParams()
	status := params.Get("status")
	limit, err := strconv.Atoi(params.Get("limit"))
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(params.Get("offset"))
	if err != nil {
		offset = 0
	}

	filter := constants.OrderFilter{
		Status: &status,
		Limit:  &limit,
		Offset: &offset,
	}
	orders, statusCode, err := o.ordersUsecase.FindAll(filter)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Orders found", responses.ToArrayOfOrdersResponse(orders))
}

func (o *OrdersHandler) Cancel(ctx echo.Context) error {
	orderId := ctx.Param("id")

	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid order id")
	}

	statusCode, err := o.ordersUsecase.Cancel(orderIdInt)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Order canceled successfully", nil)
}
