package requests

import (
	"ga_marketplace/internal/business/domains"
	"time"
)

type CreateProductStockRequest struct {
	CCode             string    `json:"c_code" validate:"required"`
	Date              time.Time `json:"date" validate:"required"`
	TransactionType   int       `json:"transaction_type" validate:"required"`
	TransactionId     string    `json:"transaction_id" validate:"required"`
	Quantity          int       `json:"quantity" validate:"required"`
	TotalSum          float64   `json:"total_sum" validate:"required"`
	TransactionStatus int       `json:"transaction_status" validate:"required"`
}

type UpdateProductStockRequest struct {
	CCode             *string    `json:"c_code"`
	Date              *time.Time `json:"date"`
	TransactionType   *int       `json:"transaction_type"`
	TransactionId     *string    `json:"transaction_id"`
	Quantity          *int       `json:"quantity"`
	TotalSum          *float64   `json:"total_sum"`
	TransactionStatus *int       `json:"transaction_status"`
}

func (r *CreateProductStockRequest) ToDomain() domains.ProductStockDomain {
	return domains.ProductStockDomain{
		CCode:             r.CCode,
		Date:              r.Date,
		TransactionType:   r.TransactionType,
		TransactionId:     r.TransactionId,
		Quantity:          r.Quantity,
		TotalSum:          r.TotalSum,
		TransactionStatus: r.TransactionStatus,
	}
}
