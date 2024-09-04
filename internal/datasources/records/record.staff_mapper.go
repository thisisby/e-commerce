package records

import "ga_marketplace/internal/business/domains"

func (rec *StaffRecord) ToDomain() *domains.StaffDomain {
	if rec == nil || rec.Id == 0 {
		return nil
	}
	return &domains.StaffDomain{
		Id:               rec.Id,
		FullName:         rec.FullName,
		Occupation:       rec.Occupation,
		Experience:       rec.Experience,
		Avatar:           rec.Avatar,
		ServiceId:        rec.ServiceId,
		ServiceAddressId: rec.ServiceAddressId,
		StartTime:        rec.StartTime,
		EndTime:          rec.EndTime,
	}
}

func FromStaffDomain(dom *domains.StaffDomain) *StaffRecord {
	return &StaffRecord{
		Id:               dom.Id,
		FullName:         dom.FullName,
		Occupation:       dom.Occupation,
		Experience:       dom.Experience,
		Avatar:           dom.Avatar,
		ServiceId:        dom.ServiceId,
		ServiceAddressId: dom.ServiceAddressId,
		StartTime:        dom.StartTime,
		EndTime:          dom.EndTime,
	}
}

func ToArrayOfStaffDomain(data []StaffRecord) (result []domains.StaffDomain) {
	for _, v := range data {
		result = append(result, *v.ToDomain())
	}
	return
}
