package records

import (
	"encoding/json"
	"ga_marketplace/internal/business/domains"
	"log/slog"
)

func (rec *StaffRecord) ToDomain() *domains.StaffDomain {
	if rec == nil || rec.Id == 0 {
		return nil
	}

	var timeSlot []domains.TimeSlot
	err := json.Unmarshal([]byte(rec.TimeSlot), &timeSlot)
	if err != nil {
		slog.Error("error unmarshal time slot", err)
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
		TimeSlot:         timeSlot,
		WorkingDays:      rec.WorkingDays,
	}
}

func FromStaffDomain(dom *domains.StaffDomain) *StaffRecord {

	timeSlot, err := json.Marshal(dom.TimeSlot)
	if err != nil {
		slog.Error("error marshal time slot", err)
		return nil
	}

	return &StaffRecord{
		Id:               dom.Id,
		FullName:         dom.FullName,
		Occupation:       dom.Occupation,
		Experience:       dom.Experience,
		Avatar:           dom.Avatar,
		ServiceId:        dom.ServiceId,
		ServiceAddressId: dom.ServiceAddressId,
		TimeSlot:         string(timeSlot),
		WorkingDays:      dom.WorkingDays,
	}
}

func ToArrayOfStaffDomain(data []StaffRecord) (result []domains.StaffDomain) {
	for _, v := range data {
		result = append(result, *v.ToDomain())
	}
	return
}
