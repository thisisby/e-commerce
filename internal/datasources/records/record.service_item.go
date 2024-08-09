package records

type ServiceItem struct {
	Id           int     `db:"id"`
	Title        string  `db:"title"`
	Duration     int     `db:"duration"`
	Description  string  `db:"description"`
	Price        float64 `db:"price"`
	SubServiceId int     `db:"subservice_id"`
}
