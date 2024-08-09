package records

type SubServiceRecord struct {
	Id        int    `db:"id"`
	Name      string `db:"name"`
	ServiceId int    `db:"service_id"`
}
