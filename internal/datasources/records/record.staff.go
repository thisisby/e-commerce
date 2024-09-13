package records

import "github.com/lib/pq"

type StaffRecord struct {
	Id               int            `db:"id"`
	FullName         string         `db:"full_name"`
	Occupation       string         `db:"occupation"`
	Experience       int            `db:"experience"`
	Avatar           *string        `db:"avatar"`
	ServiceId        int            `db:"service_id"`
	ServiceAddressId int            `db:"service_address_id"`
	TimeSlot         string         `db:"time_slot"`
	WorkingDays      pq.StringArray `db:"working_days"`
}
