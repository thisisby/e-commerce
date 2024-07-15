package responses

import (
	"ga_marketplace/internal/business/domains"
	"time"
)

type CartItemsResponse struct {
	Id        int             `json:"id"`
	UserId    int             `json:"user_id"`
	ProductId int             `json:"product_id"`
	Product   ProductResponse `json:"product"`
	Quantity  int             `json:"quantity"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type CartItemsAdminResponse struct {
	CartItemsResponse
	User UserResponse `json:"user"`
}

type CartItemTotalAmountResponse struct {
	TotalAmount   float64 `json:"total_amount"`
	TotalDiscount float64 `json:"total_discount"`
}

func FromCartItemsAdminDomain(inDom *domains.CartItemsDomain) CartItemsAdminResponse {
	return CartItemsAdminResponse{
		CartItemsResponse: CartItemsResponse{
			Id:        inDom.Id,
			UserId:    inDom.UserId,
			ProductId: inDom.ProductId,
			Product:   FromProductDomain(&inDom.Product),
			Quantity:  inDom.Quantity,
			CreatedAt: inDom.CreatedAt,
			UpdatedAt: inDom.UpdatedAt,
		},
		User: FromUserDomain(&inDom.User),
	}
}

func ToArrayOfCartItemsAdminResponse(inDom []domains.CartItemsDomain) []CartItemsAdminResponse {
	var carts []CartItemsAdminResponse
	for _, dom := range inDom {
		carts = append(carts, FromCartItemsAdminDomain(&dom))
	}
	return carts

}

func FromCartItemsDomain(inDom *domains.CartItemsDomain) CartItemsResponse {
	return CartItemsResponse{
		Id:        inDom.Id,
		UserId:    inDom.UserId,
		ProductId: inDom.ProductId,
		Product:   FromProductDomain(&inDom.Product),
		Quantity:  inDom.Quantity,
		CreatedAt: inDom.CreatedAt,
		UpdatedAt: inDom.UpdatedAt,
	}
}

func ToArrayOfCartItemsResponse(inDom []domains.CartItemsDomain) []CartItemsResponse {
	var carts []CartItemsResponse
	for _, dom := range inDom {
		carts = append(carts, FromCartItemsDomain(&dom))
	}
	return carts
}
