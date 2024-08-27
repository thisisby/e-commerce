package records

import "ga_marketplace/internal/business/domains"

func (r *ProductStock) ToDomain() domains.ProductStockDomain {
	return domains.ProductStockDomain{
		Id:                r.Id,
		CCode:             r.CCode,
		Date:              r.Date,
		TransactionType:   r.TransactionType,
		TransactionId:     r.TransactionId,
		Quantity:          r.Quantity,
		TotalSum:          r.TotalSum,
		TransactionStatus: r.TransactionStatus,
	}
}

func FromProductStockDomain(domain domains.ProductStockDomain) ProductStock {
	return ProductStock{
		Id:                domain.Id,
		CCode:             domain.CCode,
		Date:              domain.Date,
		TransactionType:   domain.TransactionType,
		TransactionId:     domain.TransactionId,
		Quantity:          domain.Quantity,
		TotalSum:          domain.TotalSum,
		TransactionStatus: domain.TransactionStatus,
	}
}
