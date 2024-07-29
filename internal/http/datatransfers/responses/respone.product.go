package responses

import (
	"ga_marketplace/internal/business/domains"
	"time"
)

type ProductResponse struct {
	Id              int                  `json:"id"`
	Name            string               `json:"name"`
	Description     string               `json:"description"`
	Price           float64              `json:"price"`
	DiscountedPrice float64              `json:"discounted_price"`
	TotalPrice      *float64             `json:"total_price"`
	IsInCart        bool                 `json:"is_in_cart"`
	IsInWishlist    bool                 `json:"is_in_wishlist"`
	Discount        *DiscountResponse    `json:"discount"`
	Stock           int                  `json:"stock"`
	SubCategoryId   int                  `json:"sub_category_id"`
	SubCategory     *SubcategoryResponse `json:"sub_category"`
	Image           string               `json:"image"`
	Images          []string             `json:"images"`
	CreatedAt       time.Time            `json:"created_at"`
	UpdatedAt       time.Time            `json:"updated_at"`
}

func FromProductDomain(inDom *domains.ProductDomain) ProductResponse {
	return ProductResponse{
		Id:              inDom.Id,
		Name:            inDom.Name,
		Description:     inDom.Description,
		Price:           inDom.Price,
		DiscountedPrice: inDom.DiscountedPrice,
		TotalPrice:      inDom.TotalPrice,
		Discount:        FromDiscountDomain(inDom.Discount),
		IsInCart:        inDom.IsInCart,
		IsInWishlist:    inDom.IsInWishlist,
		Stock:           inDom.Stock,
		SubCategoryId:   inDom.SubcategoryId,
		SubCategory:     FromSubcategoryDomain(inDom.Subcategory),
		Image:           inDom.Image,
		Images:          inDom.Images,
		CreatedAt:       inDom.CreatedAt,
		UpdatedAt:       inDom.UpdatedAt,
	}
}

func ToArrayOfProductResponse(inDom []domains.ProductDomain) []ProductResponse {
	var outDom []ProductResponse

	for _, dom := range inDom {
		outDom = append(outDom, FromProductDomain(&dom))
	}

	return outDom
}
