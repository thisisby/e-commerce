package records

import "time"

type CartItems struct {
	Id        int       `db:"id"`
	User      Users     `db:"user"`
	UserId    int       `db:"user_id"`
	Product   Products  `db:"product"`
	ProductId int       `db:"product_id"`
	Quantity  int       `db:"quantity"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
