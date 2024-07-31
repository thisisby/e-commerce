package records

type Cities struct {
	Id                   int    `db:"id"`
	Name                 string `db:"name"`
	DeliveryDurationDays int    `db:"delivery_duration_days"`
}
