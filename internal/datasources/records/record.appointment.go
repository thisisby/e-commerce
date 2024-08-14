package records

import "time"

type Appointment struct {
	Id            int          `db:"id"`
	UserId        int          `db:"user_id"`
	User          *Users       `db:"user"`
	StaffId       int          `db:"staff_id"`
	Staff         *StaffRecord `db:"staff"`
	StartTime     time.Time    `db:"start_time"`
	EndTime       time.Time    `db:"end_time"`
	ServiceItemId int          `db:"service_item_id"`
	ServiceItem   *ServiceItem `db:"service_item"`
	Comments      *string      `db:"comments"`
	Status        string       `db:"status"`
}
