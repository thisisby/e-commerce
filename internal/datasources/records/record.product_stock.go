package records

import "time"

type ProductStock struct {
	TransactionId string    `db:"transaction_id"`
	CustomerId    int       `db:"customer_id"`
	Date          time.Time `db:"date"`
	Active        bool      `db:"active"`
}

type ProductStockItem struct {
	TransactionId   string  `db:"transaction_id"`
	ProductCode     string  `db:"product_code"`
	Quantity        int     `db:"quantity"`
	Amount          float64 `db:"amount"`
	TransactionType int     `db:"transaction_type"`
}
