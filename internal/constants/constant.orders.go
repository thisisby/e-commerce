package constants

const (
	OrderPending   = "pending"
	OrderShipping  = "shipping"
	OrderCompleted = "completed"
	OrderDelivered = "delivered"
	OrderCancelled = "cancelled"
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
