package responses

import (
	"ga_marketplace/internal/business/domains"
	"time"
)

type ProductResponse struct {
	Id              int                  `json:"id"`
	Name            string               `json:"name"`
	Description     string               `json:"description"`
	Ingredients     string               `json:"ingredients"`
	Article         string               `json:"article"`
	CCode           string               `json:"c_code"`
	EdIzm           string               `json:"ed_izm"`
	Price           float64              `json:"price"`
	Weight          *float64             `json:"weight"`
	DiscountedPrice float64              `json:"discounted_price"`
	TotalPrice      *float64             `json:"total_price"`
	IsInCart        bool                 `json:"is_in_cart"`
	IsInWishlist    int                  `json:"is_in_wishlist"`
	Discount        *DiscountResponse    `json:"discount"`
	SubCategoryId   int                  `json:"sub_category_id"`
	SubCategory     *SubcategoryResponse `json:"sub_category"`
	BrandId         int                  `json:"brand_id"`
	Brand           *BrandResponse       `json:"brand"`
	Image           string               `json:"image"`
	Images          []string             `json:"images"`
	Stock           int                  `json:"stock"`
	CreatedAt       time.Time            `json:"created_at"`
	UpdatedAt       time.Time            `json:"updated_at"`
}

func FromProductDomain(inDom *domains.ProductDomain) *ProductResponse {
	return &ProductResponse{
		Id:              inDom.Id,
		Name:            inDom.Name,
		Description:     inDom.Description,
		Ingredients:     inDom.Ingredients,
		Article:         inDom.Article,
		CCode:           inDom.CCode,
		EdIzm:           inDom.EdIzm,
		Price:           inDom.Price,
		DiscountedPrice: inDom.DiscountedPrice,
		TotalPrice:      inDom.TotalPrice,
		Discount:        FromDiscountDomain(inDom.Discount),
		IsInCart:        inDom.IsInCart,
		IsInWishlist:    inDom.IsInWishlist,
		SubCategoryId:   inDom.SubcategoryId,
		Weight:          inDom.Weight,
		SubCategory:     FromSubcategoryDomain(inDom.Subcategory),
		BrandId:         inDom.BrandId,
		Brand:           FromBrandDomain(inDom.Brand),
		Image:           inDom.Image,
		Images:          inDom.Images,
		Stock:           inDom.Stock,
		CreatedAt:       inDom.CreatedAt,
		UpdatedAt:       inDom.UpdatedAt,
	}
}

func ToArrayOfProductResponse(inDom []domains.ProductDomain) []ProductResponse {
	var outDom []ProductResponse

	for _, dom := range inDom {
		outDom = append(outDom, *FromProductDomain(&dom))
	}

	return outDom
}
