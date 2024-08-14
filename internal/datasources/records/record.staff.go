package records

type StaffRecord struct {
	Id               int     `db:"id"`
	FullName         string  `db:"full_name"`
	Occupation       string  `db:"occupation"`
	Experience       int     `db:"experience"`
	Avatar           *string `db:"avatar"`
	ServiceId        int     `db:"service_id"`
	ServiceAddressId int     `db:"service_address_id"`
}
