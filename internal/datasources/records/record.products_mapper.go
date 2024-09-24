package records

import "ga_marketplace/internal/business/domains"

func (p *Products) ToDomain() *domains.ProductDomain {
	if p == nil || p.Id == 0 {
		return nil
	}
	var discountDomain *domains.DiscountsDomain
	if p.Discount != nil {
		discountDomain = p.Discount.ToDiscountsDomain()
	}

	var characteristic string
	if p.Characteristic != nil {
		characteristic = *p.Characteristic
	}
	return &domains.ProductDomain{
		Id:                  p.Id,
		Name:                p.Name,
		Description:         p.Description,
		Price:               p.Price,
		Ingredients:         p.Ingredients,
		Article:             p.Article,
		CCode:               p.CCode,
		EdIzm:               p.EdIzm,
		DiscountedPrice:     p.DiscountedPrice,
		TotalPrice:          p.TotalPrice,
		Discount:            discountDomain,
		Weight:              p.Weight,
		SubcategoryId:       p.SubcategoryId,
		Subcategory:         p.Subcategory.ToDomain(),
		BrandId:             p.BrandId,
		Brand:               p.Brand.ToDomain(),
		Image:               p.Image,
		Images:              p.Images,
		IsInCart:            p.IsInCart,
		IsInWishlist:        p.IsInWishlist,
		Stock:               p.Stock,
		Attributes:          p.Attributes,
		CountryOfProduction: p.CountryOfProduction,
		Volume:              p.Volume,
		Sex:                 p.Sex,
		Characteristic:      characteristic,
		CreatedAt:           p.CreatedAt,
		UpdatedAt:           p.UpdatedAt,
	}
}

func FromProductDomain(inDom *domains.ProductDomain) Products {
	return Products{
		Id:                  inDom.Id,
		Name:                inDom.Name,
		Description:         inDom.Description,
		Price:               inDom.Price,
		SubcategoryId:       inDom.SubcategoryId,
		BrandId:             inDom.BrandId,
		Weight:              inDom.Weight,
		Volume:              inDom.Volume,
		Sex:                 inDom.Sex,
		CountryOfProduction: inDom.CountryOfProduction,
		Image:               inDom.Image,
		Images:              inDom.Images,
		CreatedAt:           inDom.CreatedAt,
		UpdatedAt:           inDom.UpdatedAt,
	}
}

func ToArrayOfProductsDomain(inRec []Products) []domains.ProductDomain {
	var outDom []domains.ProductDomain

	for _, rec := range inRec {
		outDom = append(outDom, *rec.ToDomain())
	}

	return outDom
}
