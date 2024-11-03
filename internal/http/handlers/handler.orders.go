package handlers

import (
	"encoding/json"
	"fmt"
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/config"
	"ga_marketplace/internal/constants"
	"ga_marketplace/internal/datasources/caches"
	"ga_marketplace/internal/http/datatransfers/requests"
	"ga_marketplace/internal/http/datatransfers/responses"
	"ga_marketplace/pkg/helpers"
	"ga_marketplace/pkg/jwt"
	"ga_marketplace/third_party/forte"
	"ga_marketplace/third_party/one_c"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

type OrdersHandler struct {
	ordersUsecase    domains.OrdersUsecase
	cartItemsUsecase domains.CartUsecase
	forteClient      *forte.Client
	oneC             *one_c.Client
	redisCache       caches.RedisCache
}

func NewOrdersHandler(
	ordersUsecase domains.OrdersUsecase,
	cartItemsUsecase domains.CartUsecase,
	forteClient *forte.Client,
	oneC *one_c.Client,
	redisCache caches.RedisCache) OrdersHandler {
	return OrdersHandler{
		ordersUsecase:    ordersUsecase,
		cartItemsUsecase: cartItemsUsecase,
		forteClient:      forteClient,
		oneC:             oneC,
		redisCache:       redisCache,
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

	currentUserAgent := ctx.Request().Header.Get("User-Agent")
	if currentUserAgent == "" {
		currentUserAgent = "Default-User-Agent"
	}
	currentAcceptHeader := ctx.Request().Header.Get("Accept")
	if currentAcceptHeader == "" {
		currentAcceptHeader = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"
	}
	currentLanguage := "en-US"

	browser := requests.Browser{
		AcceptHeader:      currentAcceptHeader,
		ScreenWidth:       1920,
		WindowHeight:      1080,
		ScreenColorDepth:  24,
		WindowWidth:       1920,
		JavaEnabled:       true,
		JavascriptEnabled: true,
		Language:          currentLanguage,
		ScreenHeight:      1080,
		TimeZone:          0,
		UserAgent:         currentUserAgent,
		TimeZoneName:      "Asia/Qyzylorda",
	}

	slog.Info("add: ", browser)
	addtitonalData := requests.AdditionalData{
		Browser: browser,
	}
	paymentRequest := requests.CreatePaymentRequest{
		Amount:         totalAmount.TotalAmount * 100,
		Currency:       "KZT",
		Description:    orderCreateRequest.CreatePaymentRequest.Description,
		Language:       "ru",
		ReturnUrl:      orderCreateRequest.CreatePaymentRequest.ReturnUrl,
		Test:           orderCreateRequest.CreatePaymentRequest.Test,
		BillingAddress: orderCreateRequest.CreatePaymentRequest.BillingAddress,
		CreditCard:     orderCreateRequest.CreatePaymentRequest.CreditCard,
		Customer:       orderCreateRequest.CreatePaymentRequest.Customer,
		AdditionalData: addtitonalData,
	}

	paymentResponse, status, err := o.forteClient.CreatePayment(paymentRequest)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Payment failed: "+err.Error())
	}
	paymentData, ok := paymentResponse.(map[string]any)
	if !ok {
		return NewErrorResponse(ctx, http.StatusInternalServerError, "Unexpected response format from payment gateway")
	}

	if status != http.StatusOK {
		return NewErrorResponse(ctx, status, "Payment failed")
	}
	transaction, ok := paymentData["transaction"].(map[string]any)
	if !ok {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Transaction data not found in payment response")
	}

	fmt.Println("Transaction11111: ", transaction)
	transactionStatus := transaction["status"].(string)
	transactionID := transaction["uid"].(string)
	receiptUrl := transaction["receipt_url"].(string)

	if transactionStatus == "failed" {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Payment failed: "+transaction["message"].(string))
	}

	if transactionStatus == "successful" {
		orderDomain := orderCreateRequest.ToDomain()
		orderDomain.UserId = jwtClaims.UserId
		orderDomain.ReceiptUrl = receiptUrl
		statusCode, err = o.ordersUsecase.Save(orderDomain, cartItems, *totalAmount)
		if err != nil {
			return NewSuccessResponse(ctx, statusCode, err.Error(), paymentResponse)
		}
		return NewSuccessResponse(ctx, statusCode, "Order saved successfully", paymentResponse)
	}

	if transactionStatus == "incomplete" {
		go func() {
			for {
				time.Sleep(20 * time.Second)
				status, err := o.checkTransactionStatus(transactionID)
				if err != nil {
					log.Error("Failed to fetch transaction status: ", err)
					return
				}
				if status == "successful" {
					orderDomain := orderCreateRequest.ToDomain()
					orderDomain.UserId = jwtClaims.UserId
					orderDomain.ReceiptUrl = receiptUrl
					statusCode, err = o.ordersUsecase.Save(orderDomain, cartItems, *totalAmount)
					if err != nil {
						log.Error("Failed to save order: ", err)
						return
					}
					log.Info("Order saved successfully in background after successful payment")
					return
				} else if status == "failed" {
					log.Error("Payment failed during polling.")
					return
				}
				log.Info("Polling transaction status...")
			}
		}()
		return NewSuccessResponse(ctx, statusCode, "Order processing started. You will be notified once it is completed.", paymentResponse)
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

func (o *OrdersHandler) checkTransactionStatus(transactionID string) (string, error) {
	url := fmt.Sprintf("https://gateway.fortebank.com/transactions/%s", transactionID)

	slog.Info("Checking transaction status", "transactionID", transactionID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.SetBasicAuth(config.AppConfig.ForteUsername, config.AppConfig.FortePassword)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to fetch transaction status: %w", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode transaction status response: %w", err)
	}

	transactionData, ok := result["transaction"]
	if !ok || transactionData == nil {
		slog.Error("Transaction data not found in response", "response", result)
		return "", nil
	}

	transaction, ok := transactionData.(map[string]interface{})
	if !ok {
		slog.Error("Unexpected format for transaction data", "transactionData", transactionData)
		return "", fmt.Errorf("unexpected format for transaction data")
	}

	status, ok := transaction["status"].(string)
	if !ok {
		slog.Error("Transaction status not found or not a string", "transaction", transaction)
		return "", fmt.Errorf("transaction status not found or invalid format")
	}

	return status, nil
}
