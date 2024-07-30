package constants

const (
	Pending   = "pending"
	Shipping  = "shipping"
	Delivered = "delivered"
	Cancelled = "cancelled"
)

const (
	DeliveryMethodPickup   = "pickup"
	DeliveryMethodDelivery = "delivery"
)

type OrderFilter struct {
	Status *string
	Limit  *int
	Offset *int
}
