package requests

import (
	"ga_marketplace/internal/business/domains"
	"time"
)

type AppointmentCreateRequest struct {
	StaffId       int       `json:"staff_id" validate:"required"`
	StartTime     time.Time `json:"start_time" validate:"required"`
	ServiceItemId int       `json:"service_item_id" validate:"required"`
	Comments      string    `json:"comments"`
}

func (r *AppointmentCreateRequest) ToDomain() domains.AppointmentDomain {
	return domains.AppointmentDomain{
		StaffId:       r.StaffId,
		StartTime:     r.StartTime,
		ServiceItemId: r.ServiceItemId,
		Comments:      &r.Comments,
	}
}

type AppointmentUpdateRequest struct {
	StaffId       *int       `json:"staff_id"`
	StartTime     *time.Time `json:"start_time"`
	ServiceItemId *int       `json:"service_item_id"`
	Comments      *string    `json:"comments"`
	Status        *string    `json:"status"`
}

type AppointmentChangeTimeRequest struct {
	StartTime time.Time `json:"start_time" validate:"required"`
}
