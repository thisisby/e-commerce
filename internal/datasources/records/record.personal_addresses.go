package records

type PersonalAddresses struct {
	Id        int     `db:"id"`
	UserId    int     `db:"user_id"`
	Street    string  `db:"street"`
	Region    string  `db:"region"`
	Apartment string  `db:"apartment"`
	StreetNum string  `db:"street_num"`
	CityId    int     `db:"city_id"`
	City      *Cities `db:"city"`
	User      *Users  `db:"user"`
}
