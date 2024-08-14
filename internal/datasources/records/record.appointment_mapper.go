package records

import "ga_marketplace/internal/business/domains"

func (r *Appointment) ToDomain() *domains.AppointmentDomain {
	return &domains.AppointmentDomain{
		Id:                r.Id,
		UserId:            r.UserId,
		User:              r.User.ToDomain(),
		StaffId:           r.StaffId,
		Staff:             r.Staff.ToDomain(),
		StartTime:         r.StartTime,
		EndTime:           r.EndTime,
		ServiceItemId:     r.ServiceItemId,
		ServiceItemDomain: r.ServiceItem.ToDomain(),
		Comments:          r.Comments,
		Status:            r.Status,
	}
}

func FromAppointmentDomain(domain domains.AppointmentDomain) Appointment {
	return Appointment{
		Id:            domain.Id,
		UserId:        domain.UserId,
		StaffId:       domain.StaffId,
		StartTime:     domain.StartTime,
		EndTime:       domain.EndTime,
		ServiceItemId: domain.ServiceItemId,
		Comments:      domain.Comments,
		Status:        domain.Status,
	}
}

func ToArrayOfAppointmentDomain(data []Appointment) (result []domains.AppointmentDomain) {
	for _, v := range data {
		result = append(result, *v.ToDomain())
	}
	return
}
