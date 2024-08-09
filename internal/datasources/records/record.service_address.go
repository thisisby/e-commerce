package records

type ServiceAddress struct {
	Id      int    `db:"id"`
	CityId  int    `db:"city_id"`
	City    Cities `db:"city"`
	Address string `db:"address"`
}
