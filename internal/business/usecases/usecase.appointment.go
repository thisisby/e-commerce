package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"net/http"
	"time"
)

type appointmentUsecase struct {
	appointmentRepository domains.AppointmentRepository
	serviceItemUsecase    domains.ServiceItemUsecase
}

func NewAppointmentUsecase(
	appointmentRepository domains.AppointmentRepository,
	serviceItemUsecase domains.ServiceItemUsecase,
) domains.AppointmentUsecase {
	return &appointmentUsecase{
		appointmentRepository: appointmentRepository,
		serviceItemUsecase:    serviceItemUsecase,
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

	return appointments, http.StatusOK, nil
}

func (a *appointmentUsecase) Save(domain domains.AppointmentDomain) (int, error) {
	serviceItem, statusCode, err := a.serviceItemUsecase.FindById(domain.ServiceItemId)
	if err != nil {
		return statusCode, err
	}

	domain.EndTime = domain.StartTime.Add(time.Minute * time.Duration(serviceItem.Duration))

	// check if the new time is overlapping with other appointments
	isOverlapping, err := a.appointmentRepository.IsOverlapping(0, domain.StaffId, domain.StartTime, domain.EndTime)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if isOverlapping {
		return http.StatusBadRequest, errors.New("appointment time is overlapping with other appointments")
	}

	err = a.appointmentRepository.Save(domain)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (a *appointmentUsecase) Update(domain domains.AppointmentDomain) (int, error) {
	domain.EndTime = domain.StartTime.Add(time.Minute * time.Duration(domain.ServiceItemDomain.Duration))

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