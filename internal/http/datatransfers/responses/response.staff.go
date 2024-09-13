package responses

import (
	"ga_marketplace/internal/business/domains"
)

type StaffResponse struct {
	Id               int                `json:"id"`
	FullName         string             `json:"full_name"`
	Occupation       string             `json:"occupation"`
	Experience       int                `json:"experience"`
	Avatar           *string            `json:"avatar"`
	ServiceId        int                `json:"service_id"`
	ServiceAddressId int                `json:"service_address_id"`
	TimeSlot         []domains.TimeSlot `json:"time_slot"`
	WorkingDays      []string           `json:"working_days"`
}

func FromStaffDomain(inDom domains.StaffDomain) *StaffResponse {
	if inDom.Id == 0 {
		return nil
	}
	return &StaffResponse{
		Id:               inDom.Id,
		FullName:         inDom.FullName,
		Occupation:       inDom.Occupation,
		Experience:       inDom.Experience,
		Avatar:           inDom.Avatar,
		ServiceId:        inDom.ServiceId,
		ServiceAddressId: inDom.ServiceAddressId,
		TimeSlot:         inDom.TimeSlot,
		WorkingDays:      inDom.WorkingDays,
	}
}

func ToArrayOfStaffResponse(inDoms []domains.StaffDomain) []StaffResponse {
	var staffResponses []StaffResponse

	for _, inDom := range inDoms {
		staffResponses = append(staffResponses, *FromStaffDomain(inDom))
	}

	return staffResponses
}
