package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"log/slog"
	"net/http"
	"time"
)

type appointmentUsecase struct {
	appointmentRepository domains.AppointmentRepository
	serviceItemUsecase    domains.ServiceItemUsecase
	staffUsecase          domains.StaffUsecase
}

func NewAppointmentUsecase(
	appointmentRepository domains.AppointmentRepository,
	serviceItemUsecase domains.ServiceItemUsecase,
	staffUsecase domains.StaffUsecase,
) domains.AppointmentUsecase {
	return &appointmentUsecase{
		appointmentRepository: appointmentRepository,
		serviceItemUsecase:    serviceItemUsecase,
		staffUsecase:          staffUsecase,
	}
}

func (a *appointmentUsecase) FindAll() ([]domains.AppointmentDomain, int, error) {
	appointments, err := a.appointmentRepository.FindAll()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(appointments) == 0 {
		return nil, http.StatusNotFound, nil
	}

	return appointments, http.StatusOK, nil
}

func (a *appointmentUsecase) FindByUserId(userId int) ([]domains.AppointmentDomain, int, error) {
	appointments, err := a.appointmentRepository.FindByUserId(userId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(appointments) == 0 {
		return nil, http.StatusNotFound, nil
	}

	slog.Info("appointments usecase", appointments)

	return appointments, http.StatusOK, nil
}

func (a *appointmentUsecase) Save(domain domains.AppointmentDomain) (int, error) {
	serviceItem, statusCode, err := a.serviceItemUsecase.FindById(domain.ServiceItemId)
	if err != nil {
		return statusCode, err
	}

	domain.EndTime = domain.StartTime.Add(time.Minute * time.Duration(serviceItem.Duration))

	isOverlapping, err := a.appointmentRepository.IsOverlapping(0, domain.StaffId, domain.StartTime, domain.EndTime)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if isOverlapping {
		return http.StatusBadRequest, errors.New("appointment time is overlapping with other appointments")
	}

	availableTimeSlots, statusCode, err := a.staffUsecase.FindAvailableTimeSlot(domain.StaffId, domain.StartTime)
	if err != nil {
		return statusCode, err
	}

	appointmentStart := domain.StartTime.Format("15:04")
	appointmentEnd := domain.EndTime.Format("15:04")

	for _, slot := range availableTimeSlots {
		slotTime, _ := time.Parse("15:04", slot.Time)

		if slotTime.Format("15:04") >= appointmentStart && slotTime.Format("15:04") < appointmentEnd && !slot.IsAvailable {
			return http.StatusBadRequest, errors.New("requested time slot is unavailable")
		}
	}

	err = a.appointmentRepository.Save(domain)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (a *appointmentUsecase) Update(domain domains.AppointmentDomain) (int, error) {
	domain.EndTime = domain.StartTime.Add(time.Minute * time.Duration(domain.ServiceItemDomain.Duration))

	isOverlapping, err := a.appointmentRepository.IsOverlapping(domain.Id, domain.StaffId, domain.StartTime, domain.EndTime)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if isOverlapping {
		return http.StatusBadRequest, errors.New("appointment time is overlapping with other appointments")
	}

	availableTimeSlots, statusCode, err := a.staffUsecase.FindAvailableTimeSlot(domain.StaffId, domain.StartTime)
	if err != nil {
		return statusCode, err
	}

	appointmentStart := domain.StartTime.Format("15:04")
	appointmentEnd := domain.EndTime.Format("15:04")

	for _, slot := range availableTimeSlots {
		slotTime, _ := time.Parse("15:04", slot.Time)

		if slotTime.Format("15:04") >= appointmentStart && slotTime.Format("15:04") < appointmentEnd && !slot.IsAvailable {
			return http.StatusBadRequest, errors.New("requested time slot is unavailable")
		}
	}

	err = a.appointmentRepository.Update(domain)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (a *appointmentUsecase) Delete(id int) (int, error) {
	err := a.appointmentRepository.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (a *appointmentUsecase) FindById(id int) (domains.AppointmentDomain, int, error) {
	appointment, err := a.appointmentRepository.FindById(id)
	if err != nil {
		return domains.AppointmentDomain{}, http.StatusInternalServerError, err
	}

	if appointment.Id == 0 {
		return domains.AppointmentDomain{}, http.StatusNotFound, nil
	}

	return appointment, http.StatusOK, nil
}

func (a *appointmentUsecase) ChangeTime(domain domains.AppointmentDomain) (int, error) {
	serviceItem, statusCode, err := a.serviceItemUsecase.FindById(domain.ServiceItemId)
	if err != nil {
		return statusCode, err
	}

	domain.EndTime = domain.StartTime.Add(time.Minute * time.Duration(serviceItem.Duration))

	// check if the new time is overlapping with other appointments
	isOverlapping, err := a.appointmentRepository.IsOverlapping(domain.Id, domain.StaffId, domain.StartTime, domain.EndTime)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if isOverlapping {
		return http.StatusBadRequest, errors.New("appointment time is overlapping with other appointments")
	}

	err = a.appointmentRepository.Update(domain)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (a *appointmentUsecase) FindAllByStaffId(staffId int) ([]domains.AppointmentDomain, int, error) {
	appointments, err := a.appointmentRepository.FindAllByStaffId(staffId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(appointments) == 0 {
		return nil, http.StatusNotFound, nil
	}

	return appointments, http.StatusOK, nil
}
