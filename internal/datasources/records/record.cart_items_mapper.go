package records

import "ga_marketplace/internal/business/domains"

func (c *CartItems) ToDomain() *domains.CartItemsDomain {
	return &domains.CartItemsDomain{
		Id:        c.Id,
		UserId:    c.UserId,
		ProductId: c.ProductId,
		Product:   *c.Product.ToDomain(),
		Quantity:  c.Quantity,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func FromCartsDomain(inDom *domains.CartItemsDomain) CartItems {
	return CartItems{
		Id:        inDom.Id,
		UserId:    inDom.UserId,
		ProductId: inDom.ProductId,
		Quantity:  inDom.Quantity,
		CreatedAt: inDom.CreatedAt,
		UpdatedAt: inDom.UpdatedAt,
	}
}

func ToArrayOfCartItemsDomain(inRec []CartItems) []domains.CartItemsDomain {
	var carts []domains.CartItemsDomain
	for _, rec := range inRec {
		carts = append(carts, *rec.ToDomain())
	}
	return carts
}
