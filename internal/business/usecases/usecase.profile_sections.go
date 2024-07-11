package usecases

import (
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

	return profileSections, http.StatusOK, nil
}

func (uc *ProfileSectionsUsecase) Save(profileSection domains.ProfileSectionsDomain) (int, error) {
	err := uc.profileSectionRepo.Save(profileSection)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}
