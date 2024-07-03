package responses

import (
	"ga_marketplace/internal/business/domains"
	"time"
)

type WishResponse struct {
	Id        int             `json:"id"`
	UserId    int             `json:"user_id"`
	ProductId int             `json:"product_id"`
	Product   ProductResponse `json:"product"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

func FromWishDomain(inDom *domains.WishDomain) WishResponse {
	return WishResponse{
		Id:        inDom.Id,
		UserId:    inDom.UserId,
		ProductId: inDom.ProductId,
		Product:   FromProductDomain(&inDom.Product),
		CreatedAt: inDom.CreatedAt,
		UpdatedAt: inDom.UpdatedAt,
	}
}

func ToArrayOfWishResponse(inDom []domains.WishDomain) []WishResponse {
	var wishes []WishResponse
	for _, dom := range inDom {
		wishes = append(wishes, FromWishDomain(&dom))
	}
	return wishes
}
