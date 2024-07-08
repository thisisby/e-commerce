package records

import "ga_marketplace/internal/business/domains"

func (d *Discounts) ToDiscountsDomain() *domains.DiscountsDomain {
	return &domains.DiscountsDomain{
		Id:        d.Id,
		ProductId: d.ProductId,
		Discount:  d.Discount,
		StartDate: d.StartDate,
		EndDate:   d.EndDate,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func FromDiscountsDomain(inDom *domains.DiscountsDomain) Discounts {
	return Discounts{
		Id:        inDom.Id,
		ProductId: inDom.ProductId,
		Discount:  inDom.Discount,
		StartDate: inDom.StartDate,
		EndDate:   inDom.EndDate,
		CreatedAt: inDom.CreatedAt,
		UpdatedAt: inDom.UpdatedAt,
	}
}
