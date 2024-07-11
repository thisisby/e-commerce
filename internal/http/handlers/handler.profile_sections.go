package handlers

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/http/datatransfers/requests"
	"ga_marketplace/pkg/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProfileSectionsHandler struct {
	profileSectionsUsecase domains.ProfileSectionsUsecase
}

func NewProfileSectionsHandler(profileSectionsUsecase domains.ProfileSectionsUsecase) ProfileSectionsHandler {
	return ProfileSectionsHandler{
		profileSectionsUsecase: profileSectionsUsecase,
	}
}

func (p *ProfileSectionsHandler) FindAll(ctx echo.Context) error {
	profileSections, statusCode, err := p.profileSectionsUsecase.FindAll()
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Profile sections fetched successfully", profileSections)
}

func (p *ProfileSectionsHandler) Save(ctx echo.Context) error {
	var profileSectionCreateRequest requests.ProfileSectionCreateRequest

	err := helpers.BindAndValidate(ctx, &profileSectionCreateRequest)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	profileSectionDomain := profileSectionCreateRequest.ToDomain()
	statusCode, err := p.profileSectionsUsecase.Save(*profileSectionDomain)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Profile section saved successfully", nil)
}
