package records

import "ga_marketplace/internal/business/domains"

func (p *Products) ToDomain() *domains.ProductDomain {
	return &domains.ProductDomain{
		Id:              p.Id,
		Name:            p.Name,
		Description:     p.Description,
		Price:           p.Price,
		DiscountedPrice: p.DiscountedPrice,
		TotalPrice:      p.TotalPrice,
		Discount:        *p.Discount.ToDiscountsDomain(),
		CreatedAt:       p.CreatedAt,
		UpdatedAt:       p.UpdatedAt,
	}
}

func FromProductDomain(inDom *domains.ProductDomain) Products {
	return Products{
		Id:              inDom.Id,
		Name:            inDom.Name,
		Description:     inDom.Description,
		Price:           inDom.Price,
		DiscountedPrice: inDom.DiscountedPrice,
		TotalPrice:      inDom.TotalPrice,
		CreatedAt:       inDom.CreatedAt,
		UpdatedAt:       inDom.UpdatedAt,
	}
}
