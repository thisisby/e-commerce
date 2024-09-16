package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/constants"
	"ga_marketplace/pkg/helpers"
	"ga_marketplace/third_party/one_c"
	"log/slog"
	"net/http"
	"time"
)

type ordersUsecase struct {
	ordersRepo       domains.OrdersRepository
	productStockRepo domains.ProductStockRepository
	usersRepo        domains.UserRepository
	OnceCClient      one_c.Client
}

func NewOrdersUsecase(
	ordersRepo domains.OrdersRepository,
	productStockRepo domains.ProductStockRepository,
	usersRepo domains.UserRepository,
	oneCClient one_c.Client,
) domains.OrdersUsecase {
	return &ordersUsecase{
		ordersRepo:       ordersRepo,
		productStockRepo: productStockRepo,
		usersRepo:        usersRepo,
		OnceCClient:      oneCClient,
	}
}

func (o *ordersUsecase) Save(orders domains.OrdersDomain, cartItems []domains.CartItemsDomain, totalAmount domains.CartItemTotalAmount) (int, error) {

	orders.Status = constants.OrderPending
	orders.TotalPrice = totalAmount.TotalAmount + totalAmount.TotalDiscount
	orders.DiscountedPrice = totalAmount.TotalAmount

	for _, detail := range cartItems {
		orders.OrderDetails = append(orders.OrderDetails, domains.OrderDetailsDomain{
			ProductId:    detail.ProductId,
			Quantity:     detail.Quantity,
			Price:        detail.Product.DiscountedPrice,
			SubTotal:     detail.Product.DiscountedPrice * float64(detail.Quantity),
			ProductCCode: detail.Product.CCode,
		})
	}

	transactionId := helpers.GenerateUUID()

	for {
		transactionIdExists, err := o.productStockRepo.IsTransactionIdExist(transactionId)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		if !transactionIdExists {
			break
		}

		transactionId = helpers.GenerateUUID()
	}

	user, err := o.usersRepo.FindById(orders.UserId)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	var productsItems []one_c.Products
	for _, item := range orders.OrderDetails {
		productsItems = append(productsItems, one_c.Products{
			Quantity:  item.Quantity,
			Amount:    item.SubTotal,
			ProductId: item.ProductCCode,
		})
	}

	productSales := one_c.ProductStock{
		CustomerId:      user.Phone,
		TransactionId:   transactionId,
		Active:          true,
		TransactionDate: time.Now(),
		Products:        productsItems,
	}

	err = o.OnceCClient.CreateProductStockRequest(productSales)
	if err != nil {
		slog.Error("OrdersUsecase.Save: failed to create product stock in 1C: ", err)
		return http.StatusInternalServerError, err
	}

	orderId, err := o.ordersRepo.Save(orders)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	productStockDomain := domains.ProductStockDomain{
		TransactionId: transactionId,
		CustomerId:    user.Phone,
		Date:          time.Now(),
		Active:        true,
		OrderId:       orderId,
	}

	for _, detail := range orders.OrderDetails {
		productStockDomain.Items = append(productStockDomain.Items, domains.ProductStockItemDomain{
			TransactionId:   transactionId,
			ProductCode:     detail.ProductCCode,
			Quantity:        detail.Quantity,
			Amount:          detail.SubTotal,
			TransactionType: constants.ProductStockTransactionTypeOut,
		})
	}

	err = o.productStockRepo.Save(productStockDomain)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (o *ordersUsecase) FindByUserId(userId int, statusParam string) ([]domains.OrdersDomain, int, error) {
	orders, err := o.ordersRepo.FindByUserId(userId, statusParam)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(orders) == 0 {
		return nil, http.StatusNotFound, errors.New("orders not found")
	}

	return orders, http.StatusOK, nil
}

func (o *ordersUsecase) Update(orders domains.OrdersDomain) (int, error) {
	err := o.ordersRepo.Update(orders)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (o *ordersUsecase) FindById(id int) (domains.OrdersDomain, int, error) {
	order, err := o.ordersRepo.FindById(id)
	if err != nil {
		return domains.OrdersDomain{}, http.StatusInternalServerError, err
	}

	if order.Id == 0 {
		return domains.OrdersDomain{}, http.StatusNotFound, errors.New("order not found")
	}

	return order, http.StatusOK, nil
}

func (o *ordersUsecase) FindAll(filter constants.OrderFilter) ([]domains.OrdersDomain, int, error) {
	orders, err := o.ordersRepo.FindAll(filter)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(orders) == 0 {
		return nil, http.StatusNotFound, errors.New("orders not found")
	}

	return orders, http.StatusOK, nil
}

func (o *ordersUsecase) Cancel(id int) (int, error) {
	order, err := o.ordersRepo.FindById(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if order.Status != constants.OrderPending {
		return http.StatusBadRequest, errors.New("order cannot be cancelled at this state")
	}

	err = o.ordersRepo.Cancel(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
