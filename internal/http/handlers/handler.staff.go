package handlers

import (
	"encoding/json"
	"fmt"
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/http/datatransfers/requests"
	"ga_marketplace/pkg/helpers"
	"ga_marketplace/third_party/aws"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type StaffHandler struct {
	staffUsecase domains.StaffUsecase
	s3Client     *aws.S3Client
}

func NewStaffHandler(staffUsecase domains.StaffUsecase, s3Client *aws.S3Client) StaffHandler {
	return StaffHandler{
		staffUsecase: staffUsecase,
		s3Client:     s3Client,
	}
}

func (h *StaffHandler) Save(ctx echo.Context) error {
	var staffCreateRequest requests.CreateStaffRequest

	if err := helpers.BindAndValidate(ctx, &staffCreateRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	image, err := ctx.FormFile("avatar")
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Avatar is required")
	}

	avatarUrl, err := h.s3Client.UploadFile(image.Filename, image)

	staffDomain := &domains.StaffDomain{
		FullName:         staffCreateRequest.FullName,
		Occupation:       staffCreateRequest.Occupation,
		Experience:       staffCreateRequest.Experience,
		Avatar:           &avatarUrl,
		ServiceId:        staffCreateRequest.ServiceId,
		ServiceAddressId: staffCreateRequest.ServiceAddressId,
	}

	statusCode, err := h.staffUsecase.Save(staffDomain)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusCreated, "Staff created successfully", nil)
}

func (h *StaffHandler) FindAll(ctx echo.Context) error {
	staff, statusCode, err := h.staffUsecase.FindAll()
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Staffs found", staff)
}

func (h *StaffHandler) Update(ctx echo.Context) error {
	var staffUpdateRequest requests.UpdateStaffRequest
	staffId := ctx.Param("id")
	staffIdInt, err := strconv.Atoi(staffId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid staff id")
	}

	if err := helpers.BindAndValidate(ctx, &staffUpdateRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	avatar, _ := ctx.FormFile("avatar")
	var avatarUrl string
	if avatar != nil {
		imageUrl, err := h.s3Client.UploadFile(avatar.Filename, avatar)
		avatarUrl = imageUrl
		if err != nil {
			return NewErrorResponse(ctx, http.StatusBadRequest, "Failed to upload image")
		}
	}

	staff, statusCode, err := h.staffUsecase.FindById(staffIdInt)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	if staffUpdateRequest.FullName != nil {
		staff.FullName = *staffUpdateRequest.FullName
	}
	if staffUpdateRequest.Occupation != nil {
		staff.Occupation = *staffUpdateRequest.Occupation
	}
	if staffUpdateRequest.Experience != nil {
		staff.Experience = *staffUpdateRequest.Experience
	}
	if staffUpdateRequest.ServiceId != nil {
		staff.ServiceId = *staffUpdateRequest.ServiceId
	}
	if staffUpdateRequest.ServiceAddressId != nil {
		staff.ServiceAddressId = *staffUpdateRequest.ServiceAddressId
	}
	if avatar != nil {
		staff.Avatar = &avatarUrl
	}
	if staffUpdateRequest.TimeSlot != nil {
		var timeSlots []domains.TimeSlot

		err := json.Unmarshal([]byte(*staffUpdateRequest.TimeSlot), &timeSlots)
		if err != nil {
			return NewErrorResponse(ctx, http.StatusBadRequest, fmt.Sprintf("Invalid time slot: %v", err))
		}
		staff.TimeSlot = timeSlots
	}
	if staffUpdateRequest.WorkingDays != nil {
		staff.WorkingDays = *staffUpdateRequest.WorkingDays
	}

	statusCode, err = h.staffUsecase.Update(*staff)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Staff updated successfully", nil)
}

func (h *StaffHandler) Delete(ctx echo.Context) error {
	staffId := ctx.Param("id")
	staffIdInt, err := strconv.Atoi(staffId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid staff id")
	}

	statusCode, err := h.staffUsecase.Delete(staffIdInt)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Staff deleted successfully", nil)
}

func (h *StaffHandler) FindByServiceId(ctx echo.Context) error {
	serviceId := ctx.Param("service_id")
	serviceIdInt, err := strconv.Atoi(serviceId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid service id")
	}

	staff, statusCode, err := h.staffUsecase.FindByServiceId(serviceIdInt)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Staffs found", staff)
}

func (h *StaffHandler) FindByServiceAddressId(ctx echo.Context) error {
	serviceAddressId := ctx.Param("service_address_id")
	serviceAddressIdInt, err := strconv.Atoi(serviceAddressId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid service address id")
	}

	staff, statusCode, err := h.staffUsecase.FindByServiceAddressId(serviceAddressIdInt)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Staffs found", staff)
}

func (h *StaffHandler) FindTimeSlotByStaffId(ctx echo.Context) error {
	staffIdParam := ctx.Param("staffId")
	staffId, err := strconv.Atoi(staffIdParam)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid staff ID")
	}

	dateParam := ctx.QueryParam("date")
	if dateParam == "" {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Date parameter is required")
	}

	date, err := time.Parse("2006-01-02", dateParam)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid date format")
	}

	timeSlots, statusCode, err := h.staffUsecase.FindAvailableTimeSlot(staffId, date)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Time slots found", timeSlots)
}
