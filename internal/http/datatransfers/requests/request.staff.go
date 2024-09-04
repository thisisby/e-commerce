package requests

import "time"

type CreateStaffRequest struct {
	FullName         string    `form:"full_name" validate:"required"`
	Occupation       string    `form:"occupation" validate:"required"`
	Experience       int       `form:"experience" validate:"required"`
	ServiceId        int       `form:"service_id" validate:"required"`
	ServiceAddressId int       `form:"service_address_id" validate:"required"`
	StartTime        time.Time `form:"start_time" validate:"required"`
	EndTime          time.Time `form:"end_time" validate:"required"`
}

type UpdateStaffRequest struct {
	FullName         *string    `form:"full_name"`
	Occupation       *string    `form:"occupation"`
	Experience       *int       `form:"experience"`
	ServiceId        *int       `form:"service_id"`
	ServiceAddressId *int       `form:"service_address_id"`
	StartTime        *time.Time `form:"start_time"`
	EndTime          *time.Time `form:"end_time"`
}
