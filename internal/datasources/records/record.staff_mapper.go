package records

import "ga_marketplace/internal/business/domains"

func (rec *StaffRecord) ToDomain() *domains.StaffDomain {
	return &domains.StaffDomain{
		Id:         rec.Id,
		FullName:   rec.FullName,
		Occupation: rec.Occupation,
		Experience: rec.Experience,
		Avatar:     rec.Avatar,
		ServiceId:  rec.ServiceId,
	}
}

func FromStaffDomain(dom *domains.StaffDomain) *StaffRecord {
	return &StaffRecord{
		Id:         dom.Id,
		FullName:   dom.FullName,
		Occupation: dom.Occupation,
		Experience: dom.Experience,
		Avatar:     dom.Avatar,
		ServiceId:  dom.ServiceId,
	}
}

func ToArrayOfStaffDomain(data []StaffRecord) (result []domains.StaffDomain) {
	for _, v := range data {
		result = append(result, *v.ToDomain())
	}
	return
}
