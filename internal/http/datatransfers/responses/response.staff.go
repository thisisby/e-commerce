package responses

import (
	"ga_marketplace/internal/business/domains"
	"time"
)

type StaffResponse struct {
	Id               int       `json:"id"`
	FullName         string    `json:"full_name"`
	Occupation       string    `json:"occupation"`
	Experience       int       `json:"experience"`
	Avatar           *string   `json:"avatar"`
	ServiceId        int       `json:"service_id"`
	ServiceAddressId int       `json:"service_address_id"`
	StartTime        time.Time `json:"start_time"`
	EndTime          time.Time `json:"end_time"`
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
		StartTime:        inDom.StartTime,
		EndTime:          inDom.EndTime,
	}
}

func ToArrayOfStaffResponse(inDoms []domains.StaffDomain) []StaffResponse {
	var staffResponses []StaffResponse

	for _, inDom := range inDoms {
		staffResponses = append(staffResponses, *FromStaffDomain(inDom))
	}

	return staffResponses
}
