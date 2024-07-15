package handlers

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/http/datatransfers/requests"
	"ga_marketplace/pkg/helpers"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"strconv"
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

func (p *ProfileSectionsHandler) UpdateById(ctx echo.Context) error {
	var profileSectionUpdateRequest requests.ProfileSectionUpdateRequest
	profileSectionId := ctx.Param("id")

	err := helpers.BindAndValidate(ctx, &profileSectionUpdateRequest)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	profileSectionIdInt, err := strconv.Atoi(profileSectionId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid profile section id")
	}
	profile, statusCode, err := p.profileSectionsUsecase.FindById(profileSectionIdInt)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	slog.Info("profileSectionUpdateRequest", profileSectionUpdateRequest)
	slog.Info("profile", profile)
	if profileSectionUpdateRequest.Name != nil {
		profile.Name = *profileSectionUpdateRequest.Name
	}
	if profileSectionUpdateRequest.Content != nil {
		profile.Content = profileSectionUpdateRequest.Content
	}
	if profileSectionUpdateRequest.ParentId != nil {
		profile.ParentId = profileSectionUpdateRequest.ParentId
	}

	slog.Info("profile", profile)

	statusCode, err = p.profileSectionsUsecase.UpdateById(profile)

	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Profile section updated successfully", nil)
}

func (p *ProfileSectionsHandler) DeleteById(ctx echo.Context) error {
	profileSectionId := ctx.Param("id")
	profileSectionIdInt, err := strconv.Atoi(profileSectionId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid profile section id")
	}

	statusCode, err := p.profileSectionsUsecase.DeleteById(profileSectionIdInt)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Profile section deleted successfully", nil)
}
