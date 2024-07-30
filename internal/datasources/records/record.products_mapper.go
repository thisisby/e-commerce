package records

import "ga_marketplace/internal/business/domains"

func (p *Products) ToDomain() *domains.ProductDomain {
	var discountDomain *domains.DiscountsDomain
	if p.Discount != nil {
		discountDomain = p.Discount.ToDiscountsDomain()
	}
	return &domains.ProductDomain{
		Id:              p.Id,
		Name:            p.Name,
		Description:     p.Description,
		Price:           p.Price,
		DiscountedPrice: p.DiscountedPrice,
		TotalPrice:      p.TotalPrice,
		Discount:        discountDomain,
		SubcategoryId:   p.SubcategoryId,
		Subcategory:     p.Subcategory.ToDomain(),
		BrandId:         p.BrandId,
		Brand:           p.Brand.ToDomain(),
		Image:           p.Image,
		Images:          p.Images,
		IsInCart:        p.IsInCart,
		IsInWishlist:    p.IsInWishlist,
		Stock:           p.Stock,
		CreatedAt:       p.CreatedAt,
		UpdatedAt:       p.UpdatedAt,
	}
}

func FromProductDomain(inDom *domains.ProductDomain) Products {
	return Products{
		Id:            inDom.Id,
		Name:          inDom.Name,
		Description:   inDom.Description,
		Price:         inDom.Price,
		SubcategoryId: inDom.SubcategoryId,
		BrandId:       inDom.BrandId,
		Image:         inDom.Image,
		Images:        inDom.Images,
		Stock:         inDom.Stock,
		CreatedAt:     inDom.CreatedAt,
		UpdatedAt:     inDom.UpdatedAt,
	}
}

func ToArrayOfProductsDomain(inRec []Products) []domains.ProductDomain {
	var outDom []domains.ProductDomain

	for _, rec := range inRec {
		outDom = append(outDom, *rec.ToDomain())
	}

	return outDom
}
