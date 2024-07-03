package records

import "ga_marketplace/internal/business/domains"

func (w *Wish) ToDomain() *domains.WishDomain {
	return &domains.WishDomain{
		Id:        w.Id,
		UserId:    w.UserId,
		ProductId: w.ProductId,
		Product:   *w.Product.ToDomain(),
		CreatedAt: w.CreatedAt,
		UpdatedAt: w.UpdatedAt,
	}
}

func FromWishDomain(inDom *domains.WishDomain) Wish {
	return Wish{
		Id:        inDom.Id,
		UserId:    inDom.UserId,
		ProductId: inDom.ProductId,
		CreatedAt: inDom.CreatedAt,
		UpdatedAt: inDom.UpdatedAt,
	}
}

func ToArrayOfWishDomain(inRec []Wish) []domains.WishDomain {
	var wishes []domains.WishDomain
	for _, rec := range inRec {
		wishes = append(wishes, *rec.ToDomain())
	}
	return wishes
}
