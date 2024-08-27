package records

import "time"

type ProductStock struct {
	Id                int       `db:"id"`
	CCode             string    `db:"c_code"`
	Date              time.Time `db:"date"`
	TransactionType   int       `db:"transaction_type"`
	TransactionId     string    `db:"transaction_id"`
	Quantity          int       `db:"quantity"`
	TotalSum          float64   `db:"total_sum"`
	TransactionStatus int       `db:"transaction_status"`
}
