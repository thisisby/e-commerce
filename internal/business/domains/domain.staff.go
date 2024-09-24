package domains

import "time"

type StaffDomain struct {
	Id               int
	FullName         string
	Occupation       string
	Experience       int
	Avatar           *string
	ServiceId        int
	ServiceAddressId int
	TimeSlot         []TimeSlot
	WorkingDays      []string
	ServiceAddress   *string
}

type TimeSlot struct {
	Id          int    `json:"id"`
	Time        string `json:"time"`
	IsAvailable bool   `json:"isAvailable"`
}

type StaffRepository interface {
	FindById(id int) (*StaffDomain, error)
	Save(staff *StaffDomain) error
	FindAll() ([]StaffDomain, error)
	Update(inDom StaffDomain) error
	Delete(id int) error
	FindByServiceId(serviceId int) ([]StaffDomain, error)
	FindByServiceAddressId(serviceAddressId int) ([]StaffDomain, error)
}

type StaffUsecase interface {
	Save(staff *StaffDomain) (int, error)
	FindAll() ([]StaffDomain, int, error)
	Update(inDom StaffDomain) (int, error)
	FindById(id int) (*StaffDomain, int, error)
	Delete(id int) (int, error)
	FindByServiceId(serviceId int) ([]StaffDomain, int, error)
	FindByServiceAddressId(serviceAddressId int) ([]StaffDomain, int, error)
	FindAvailableTimeSlot(staffId int, date time.Time) ([]TimeSlot, int, error)
}
