package domains

import "time"

type AppointmentDomain struct {
	Id                int
	UserId            int
	User              *UserDomain
	StaffId           int
	Staff             *StaffDomain
	StartTime         time.Time
	EndTime           time.Time
	ServiceItemId     int
	ServiceItemDomain *ServiceItemDomain
	Comments          *string
	Status            string
}

type AppointmentRepository interface {
	FindAll() ([]AppointmentDomain, error)
	FindByUserId(userId int) ([]AppointmentDomain, error)
	Save(AppointmentDomain) error
	Update(AppointmentDomain) error
	Delete(id int) error
	FindById(id int) (AppointmentDomain, error)
	IsOverlapping(appointmentId int, staffId int, startTime time.Time, endTime time.Time) (bool, error)
}

type AppointmentUsecase interface {
	FindAll() ([]AppointmentDomain, int, error)
	FindByUserId(userId int) ([]AppointmentDomain, int, error)
	Save(AppointmentDomain) (int, error)
	Update(AppointmentDomain) (int, error)
	Delete(id int) (int, error)
	FindById(id int) (AppointmentDomain, int, error)
	ChangeTime(AppointmentDomain) (int, error)
}
