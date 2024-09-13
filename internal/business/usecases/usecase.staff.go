package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"net/http"
	"time"
)

type staffUsecase struct {
	staffRepo       domains.StaffRepository
	appointmentRepo domains.AppointmentRepository
}

func NewStaffUsecase(staffRepo domains.StaffRepository, appointmentRepo domains.AppointmentRepository) domains.StaffUsecase {
	return &staffUsecase{
		staffRepo:       staffRepo,
		appointmentRepo: appointmentRepo,
	}
}

func (s *staffUsecase) Save(staff *domains.StaffDomain) (int, error) {
	err := s.staffRepo.Save(staff)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (s *staffUsecase) FindAll() ([]domains.StaffDomain, int, error) {
	staffs, err := s.staffRepo.FindAll()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(staffs) == 0 {
		return nil, http.StatusNotFound, errors.New("staffs not found")
	}

	return staffs, http.StatusOK, nil
}

func (s *staffUsecase) Update(inDom domains.StaffDomain) (int, error) {
	err := s.staffRepo.Update(inDom)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (s *staffUsecase) FindById(id int) (*domains.StaffDomain, int, error) {
	staff, err := s.staffRepo.FindById(id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if staff == nil {
		return nil, http.StatusNotFound, errors.New("staff not found")
	}

	return staff, http.StatusOK, nil
}

func (s *staffUsecase) Delete(id int) (int, error) {
	err := s.staffRepo.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (s *staffUsecase) FindByServiceId(serviceId int) ([]domains.StaffDomain, int, error) {
	staffs, err := s.staffRepo.FindByServiceId(serviceId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(staffs) == 0 {
		return nil, http.StatusNotFound, errors.New("staffs not found")
	}

	return staffs, http.StatusOK, nil
}

func (s *staffUsecase) FindByServiceAddressId(serviceAddressId int) ([]domains.StaffDomain, int, error) {
	staffs, err := s.staffRepo.FindByServiceAddressId(serviceAddressId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(staffs) == 0 {
		return nil, http.StatusNotFound, errors.New("staffs not found")
	}

	return staffs, http.StatusOK, nil
}

func (s *staffUsecase) FindAvailableTimeSlot(staffId int, date time.Time) ([]domains.TimeSlot, int, error) {
	staff, err := s.staffRepo.FindById(staffId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if staff == nil {
		return nil, http.StatusNotFound, errors.New("staff not found")
	}

	dayOfWeek := date.Weekday().String()
	isWorkingDay := false
	for _, workingDay := range staff.WorkingDays {
		if workingDay == dayOfWeek {
			isWorkingDay = true
			break
		}
	}

	if !isWorkingDay {
		return nil, http.StatusNotFound, errors.New("staff is not working on this day")
	}

	staffAppointments, err := s.appointmentRepo.FindAllByStaffIdAndDate(staffId, date)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	timeSlots := make([]domains.TimeSlot, len(staff.TimeSlot))
	copy(timeSlots, staff.TimeSlot)

	for _, appointment := range staffAppointments {
		for i := range timeSlots {
			slotTime, _ := time.Parse("15:04", timeSlots[i].Time)
			appointmentStart := appointment.StartTime.Format("15:04")
			appointmentEnd := appointment.EndTime.Format("15:04")

			if slotTime.Format("15:04") >= appointmentStart && slotTime.Format("15:04") < appointmentEnd {
				timeSlots[i].IsAvailable = false
			}
		}
	}

	return timeSlots, http.StatusOK, nil
}
