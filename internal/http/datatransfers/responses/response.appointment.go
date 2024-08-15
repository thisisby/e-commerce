package responses

import (
	"ga_marketplace/internal/business/domains"
	"time"
)

type AppointmentResponse struct {
	Id            int                  `json:"id"`
	UserId        int                  `json:"user_id"`
	StaffId       int                  `json:"staff_id"`
	Staff         *StaffResponse       `json:"staff"`
	StartTime     time.Time            `json:"start_time"`
	EndTime       time.Time            `json:"end_time"`
	ServiceItemId int                  `json:"service_item_id"`
	ServiceItem   *ServiceItemResponse `json:"service_item"`
	Comments      *string              `json:"comments"`
	Status        string               `json:"status"`
	FullName      string               `json:"full_name"`
	PhoneNumber   string               `json:"phone_number"`
}

func FromAppointmentDomain(inDom *domains.AppointmentDomain) AppointmentResponse {
	return AppointmentResponse{
		Id:            inDom.Id,
		UserId:        inDom.UserId,
		StaffId:       inDom.StaffId,
		Staff:         FromStaffDomain(*inDom.Staff),
		StartTime:     inDom.StartTime,
		EndTime:       inDom.EndTime,
		ServiceItemId: inDom.ServiceItemId,
		ServiceItem:   FromServiceItemDomain(*inDom.ServiceItemDomain),
		Comments:      inDom.Comments,
		Status:        inDom.Status,
		FullName:      inDom.FullName,
		PhoneNumber:   inDom.PhoneNumber,
	}
}

func ToArrayOfAppointmentResponse(inDoms []domains.AppointmentDomain) []AppointmentResponse {
	var appointmentResponses []AppointmentResponse

	for _, inDom := range inDoms {
		appointmentResponses = append(appointmentResponses, FromAppointmentDomain(&inDom))
	}

	return appointmentResponses
}
