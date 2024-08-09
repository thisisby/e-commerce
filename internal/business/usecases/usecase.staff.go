package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"net/http"
)

type staffUsecase struct {
	staffRepo domains.StaffRepository
}

func NewStaffUsecase(staffRepo domains.StaffRepository) domains.StaffUsecase {
	return &staffUsecase{
		staffRepo: staffRepo,
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
