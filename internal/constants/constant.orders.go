package constants

const (
	Pending   = "pending"
	Shipping  = "shipping"
	Delivered = "delivered"
	Cancelled = "cancelled"
)

type OrderFilter struct {
	Status *string
	Limit  *int
	Offset *int
}
