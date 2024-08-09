package responses

import "ga_marketplace/internal/business/domains"

type StaffResponse struct {
	Id         int     `json:"id"`
	FullName   string  `json:"full_name"`
	Occupation string  `json:"occupation"`
	Experience int     `json:"experience"`
	Avatar     *string `json:"avatar"`
	ServiceId  int     `json:"service_id"`
}

func FromStaffDomain(inDom domains.StaffDomain) StaffResponse {
	return StaffResponse{
		Id:         inDom.Id,
		FullName:   inDom.FullName,
		Occupation: inDom.Occupation,
		Experience: inDom.Experience,
		Avatar:     inDom.Avatar,
		ServiceId:  inDom.ServiceId,
	}
}

func ToArrayOfStaffResponse(inDoms []domains.StaffDomain) []StaffResponse {
	var staffResponses []StaffResponse

	for _, inDom := range inDoms {
		staffResponses = append(staffResponses, FromStaffDomain(inDom))
	}

	return staffResponses
}
