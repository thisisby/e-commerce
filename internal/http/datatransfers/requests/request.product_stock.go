package requests

import (
	"ga_marketplace/internal/business/domains"
	"time"
)

type CreateProductStockRequest struct {
	Date          time.Time `json:"date" validate:"required"`
	TransactionId string    `json:"transaction_id_mp" validate:"required"`
	CustomerId    string    `json:"customer_id" validate:"required"`
	Active        bool      `json:"active"`
	Items         []CreateProductStockItemRequest
}

type CreateProductStockItemRequest struct {
	ProductCode     string  `json:"product_id" validate:"required"`
	Quantity        int     `json:"quantity" validate:"required"`
	Amount          float64 `json:"amount" validate:"required"`
	TransactionType int     `json:"transaction_type" validate:"required"`
}

type UpdateProductStockRequest struct {
	Date          *time.Time `json:"date"`
	TransactionId *string    `json:"transaction_id_mp"`
	CustomerId    *string    `json:"customer_id"`
	Active        *bool      `json:"active"`
}

type UpdateProductStockItemRequest struct {
	ProductCode     *string  `json:"product_id"`
	Quantity        *int     `json:"quantity"`
	Amount          *float64 `json:"amount"`
	TransactionType *int     `json:"transaction_type"`
}

func ConvertToProductStockDomain(req CreateProductStockRequest) domains.ProductStockDomain {
	items := make([]domains.ProductStockItemDomain, len(req.Items))

	for i, item := range req.Items {
		items[i] = domains.ProductStockItemDomain{
			TransactionId:   req.TransactionId,
			ProductCode:     item.ProductCode,
			Quantity:        item.Quantity,
			Amount:          item.Amount,
			TransactionType: item.TransactionType,
		}
	}

	return domains.ProductStockDomain{
		TransactionId: req.TransactionId,
		Date:          req.Date,
		Active:        req.Active,
		Items:         items,
		CustomerId:    req.CustomerId,
	}
}
