package responses

import "ga_marketplace/internal/business/domains"

type OrdersResponse struct {
	Id              int                    `json:"id"`
	UserId          int                    `json:"user_id"`
	User            UserResponse           `json:"user"`
	OrderDetails    []OrderDetailsResponse `json:"order_details"`
	TotalPrice      float64                `json:"total_price"`
	DiscountedPrice float64                `json:"discounted_price"`
	CityId          int                    `json:"city_id"`
	City            CityResponse           `json:"city"`
	Status          string                 `json:"status"`
	Street          string                 `json:"street"`
	Region          string                 `json:"region"`
	Apartment       string                 `json:"apartment"`
}

type OrderDetailsResponse struct {
	Id        int             `json:"id"`
	OrderId   int             `json:"order_id"`
	ProductId int             `json:"product_id"`
	Product   ProductResponse `json:"product"`
	Quantity  int             `json:"quantity"`
	Price     float64         `json:"price"`
	SubTotal  float64         `json:"sub_total"`
}

func FromOrdersDomain(inDom domains.OrdersDomain) OrdersResponse {
	return OrdersResponse{
		Id:              inDom.Id,
		UserId:          inDom.UserId,
		User:            FromUserDomain(&inDom.User),
		OrderDetails:    ToArrayOfOrderDetailsResponse(inDom.OrderDetails),
		TotalPrice:      inDom.TotalPrice,
		DiscountedPrice: inDom.DiscountedPrice,
		CityId:          inDom.CityId,
		City:            *FromCityDomain(&inDom.City),
		Status:          inDom.Status,
		Street:          inDom.Street,
		Region:          inDom.Region,
		Apartment:       inDom.Apartment,
	}
}

func ToArrayOfOrdersResponse(inDom []domains.OrdersDomain) []OrdersResponse {
	var outDom []OrdersResponse

	for _, dom := range inDom {
		outDom = append(outDom, FromOrdersDomain(dom))
	}

	return outDom
}

func FromOrderDetailsDomain(inDom *domains.OrderDetailsDomain) OrderDetailsResponse {
	return OrderDetailsResponse{
		Id:        inDom.Id,
		OrderId:   inDom.OrderId,
		ProductId: inDom.ProductId,
		Product:   FromProductDomain(&inDom.Product),
		Quantity:  inDom.Quantity,
		Price:     inDom.Price,
		SubTotal:  inDom.SubTotal,
	}
}

func ToArrayOfOrderDetailsResponse(inDom []domains.OrderDetailsDomain) []OrderDetailsResponse {
	var outDom []OrderDetailsResponse

	for _, dom := range inDom {
		outDom = append(outDom, FromOrderDetailsDomain(&dom))
	}

	return outDom
}
