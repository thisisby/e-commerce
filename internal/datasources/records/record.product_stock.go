package records

import "time"

type ProductStock struct {
	TransactionId string    `db:"transaction_id"`
	CustomerId    string    `db:"customer_id"`
	Date          time.Time `db:"date"`
	Active        bool      `db:"active"`
	OrderId       *int      `db:"order_id"`
}

type ProductStockItem struct {
	TransactionId   string  `db:"transaction_id"`
	ProductCode     string  `db:"product_code"`
	Quantity        int     `db:"quantity"`
	Amount          float64 `db:"amount"`
	TransactionType int     `db:"transaction_type"`
}
