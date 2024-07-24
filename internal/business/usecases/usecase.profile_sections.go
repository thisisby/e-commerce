package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"net/http"
)

type ProfileSectionsUsecase struct {
	profileSectionRepo domains.ProfileSectionsRepository
}

func NewProfileSectionsUsecase(profileSectionRepo domains.ProfileSectionsRepository) domains.ProfileSectionsUsecase {
	return &ProfileSectionsUsecase{
		profileSectionRepo: profileSectionRepo,
	}
}

func (uc *ProfileSectionsUsecase) FindAll() ([]domains.ProfileSectionsDomain, int, error) {
	profileSections, err := uc.profileSectionRepo.FindAll()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(profileSections) == 0 {
		return nil, http.StatusNotFound, errors.New("profile sections not found")
	}

	return profileSections, http.StatusOK, nil
}

func (uc *ProfileSectionsUsecase) Save(profileSection domains.ProfileSectionsDomain) (int, error) {
	err := uc.profileSectionRepo.Save(profileSection)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (uc *ProfileSectionsUsecase) UpdateById(profileSection domains.ProfileSectionsDomain) (int, error) {
	err := uc.profileSectionRepo.UpdateById(profileSection)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (uc *ProfileSectionsUsecase) FindById(id int) (domains.ProfileSectionsDomain, int, error) {
	profileSection, err := uc.profileSectionRepo.FindById(id)
	if err != nil {
		return domains.ProfileSectionsDomain{}, http.StatusInternalServerError, err
	}

	if profileSection.Id == 0 {
		return domains.ProfileSectionsDomain{}, http.StatusNotFound, errors.New("profile section not found")
	}

	return profileSection, http.StatusOK, nil
}

func (uc *ProfileSectionsUsecase) DeleteById(id int) (int, error) {
	err := uc.profileSectionRepo.DeleteById(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
