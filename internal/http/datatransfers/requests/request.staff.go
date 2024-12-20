package requests

type CreateStaffRequest struct {
	FullName         string `form:"full_name" validate:"required"`
	Occupation       string `form:"occupation" validate:"required"`
	Experience       int    `form:"experience" validate:"required"`
	ServiceId        int    `form:"service_id" validate:"required"`
	ServiceAddressId int    `form:"service_address_id" validate:"required"`
}

type UpdateStaffRequest struct {
	FullName         *string   `form:"full_name"`
	Occupation       *string   `form:"occupation"`
	Experience       *int      `form:"experience"`
	ServiceId        *int      `form:"service_id"`
	ServiceAddressId *int      `form:"service_address_id"`
	TimeSlot         *string   `form:"time_slot"`
	WorkingDays      *[]string `form:"working_days"`
}
