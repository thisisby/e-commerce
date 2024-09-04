package records

import "ga_marketplace/internal/business/domains"

func (r *ProductStock) ToDomain() domains.ProductStockDomain {
	return domains.ProductStockDomain{
		TransactionId: r.TransactionId,
		CustomerId:    r.CustomerId,
		Date:          r.Date,
		Active:        r.Active,
	}
}

func FromProductStockDomain(domain domains.ProductStockDomain) ProductStock {
	return ProductStock{
		TransactionId: domain.TransactionId,
		CustomerId:    domain.CustomerId,
		Date:          domain.Date,
		Active:        domain.Active,
	}
}

func (r *ProductStockItem) ToDomain() domains.ProductStockItemDomain {
	return domains.ProductStockItemDomain{
		TransactionId:   r.TransactionId,
		ProductCode:     r.ProductCode,
		Quantity:        r.Quantity,
		Amount:          r.Amount,
		TransactionType: r.TransactionType,
	}
}

func FromProductStockItemDomain(domain domains.ProductStockItemDomain) ProductStockItem {
	return ProductStockItem{
		TransactionId:   domain.TransactionId,
		ProductCode:     domain.ProductCode,
		Quantity:        domain.Quantity,
		Amount:          domain.Amount,
		TransactionType: domain.TransactionType,
	}
}
