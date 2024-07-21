package usecases

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/constants"
	"net/http"
)

type ordersUsecase struct {
	ordersRepo domains.OrdersRepository
}

func NewOrdersUsecase(ordersRepo domains.OrdersRepository) domains.OrdersUsecase {
	return &ordersUsecase{
		ordersRepo: ordersRepo,
	}
}

func (o *ordersUsecase) Save(orders domains.OrdersDomain, cartItems []domains.CartItemsDomain, totalAmount domains.CartItemTotalAmount) (int, error) {

	orders.Status = constants.Pending
	orders.TotalPrice = totalAmount.TotalAmount + totalAmount.TotalDiscount
	orders.DiscountedPrice = totalAmount.TotalAmount

	for _, detail := range cartItems {
		orders.OrderDetails = append(orders.OrderDetails, domains.OrderDetailsDomain{
			ProductId: detail.ProductId,
			Quantity:  detail.Quantity,
			Price:     detail.Product.DiscountedPrice,
			SubTotal:  detail.Product.DiscountedPrice * float64(detail.Quantity),
		})
	}

	err := o.ordersRepo.Save(orders)
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

	return order, http.StatusOK, nil
}
