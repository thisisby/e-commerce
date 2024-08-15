package handlers

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/constants"
	"ga_marketplace/internal/http/datatransfers/requests"
	"ga_marketplace/pkg/helpers"
	"ga_marketplace/pkg/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type AppointmentHandler struct {
	appointmentUsecase domains.AppointmentUsecase
}

func NewAppointmentHandler(appointmentUsecase domains.AppointmentUsecase) *AppointmentHandler {
	return &AppointmentHandler{
		appointmentUsecase: appointmentUsecase,
	}
}

func (h *AppointmentHandler) CreateAppointment(ctx echo.Context) error {
	var appointmentCreateRequest requests.AppointmentCreateRequest
	jwtClaims := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims)

	err := helpers.BindAndValidate(ctx, &appointmentCreateRequest)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	appointmentDomain := appointmentCreateRequest.ToDomain()
	appointmentDomain.UserId = jwtClaims.UserId
	appointmentDomain.Status = constants.AppointmentStatusPending

	statusCode, err := h.appointmentUsecase.Save(appointmentDomain)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Appointment created successfully", nil)
}

func (h *AppointmentHandler) UpdateAppointment(ctx echo.Context) error {
	var appointmentUpdateRequest requests.AppointmentUpdateRequest
	appointmentId := ctx.Param("id")

	appointmentIdInt, err := strconv.Atoi(appointmentId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid appointment id")
	}

	err = helpers.BindAndValidate(ctx, &appointmentUpdateRequest)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	appointment, statusCode, err := h.appointmentUsecase.FindById(appointmentIdInt)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	if appointmentUpdateRequest.Status != nil {
		appointment.Status = *appointmentUpdateRequest.Status
	}
	if appointmentUpdateRequest.StaffId != nil {
		appointment.StaffId = *appointmentUpdateRequest.StaffId
	}
	if appointmentUpdateRequest.StartTime != nil {
		appointment.StartTime = *appointmentUpdateRequest.StartTime
	}
	if appointmentUpdateRequest.ServiceItemId != nil {
		appointment.ServiceItemId = *appointmentUpdateRequest.ServiceItemId
	}
	if appointmentUpdateRequest.Comments != nil {
		appointment.Comments = appointmentUpdateRequest.Comments
	}
	if appointmentUpdateRequest.FullName != nil {
		appointment.FullName = *appointmentUpdateRequest.FullName
	}
	if appointmentUpdateRequest.PhoneNumber != nil {
		appointment.PhoneNumber = *appointmentUpdateRequest.PhoneNumber
	}

	statusCode, err = h.appointmentUsecase.Update(appointment)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Appointment updated successfully", nil)
}

func (h *AppointmentHandler) DeleteAppointment(ctx echo.Context) error {
	appointmentId := ctx.Param("id")

	appointmentIdInt, err := strconv.Atoi(appointmentId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid appointment id")
	}

	statusCode, err := h.appointmentUsecase.Delete(appointmentIdInt)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Appointment deleted successfully", nil)
}

func (h *AppointmentHandler) FindAllAppointments(ctx echo.Context) error {
	appointments, statusCode, err := h.appointmentUsecase.FindAll()
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Appointments found", appointments)
}

func (h *AppointmentHandler) FindMyAppointments(ctx echo.Context) error {
	jwtClaims := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims)

	appointments, statusCode, err := h.appointmentUsecase.FindByUserId(jwtClaims.UserId)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Appointments found", appointments)
}

func (h *AppointmentHandler) ChangeTime(ctx echo.Context) error {
	jwtClaims := ctx.Get(constants.CtxAuthenticatedUserKey).(jwt.JWTCustomClaims)
	var appointmentChangeTimeRequest requests.AppointmentChangeTimeRequest
	appointmentId := ctx.Param("id")

	appointmentIdInt, err := strconv.Atoi(appointmentId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid appointment id")
	}

	err = helpers.BindAndValidate(ctx, &appointmentChangeTimeRequest)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	appointment, statusCode, err := h.appointmentUsecase.FindById(appointmentIdInt)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}
	if appointment.UserId != jwtClaims.UserId {
		return NewErrorResponse(ctx, http.StatusForbidden, "You are not allowed to change this appointment")
	}

	appointment.StartTime = appointmentChangeTimeRequest.StartTime
	statusCode, err = h.appointmentUsecase.ChangeTime(appointment)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Appointment time changed successfully", nil)
}

func (h *AppointmentHandler) FindAllAppointmentsByStaffId(ctx echo.Context) error {
	staffId := ctx.Param("staff_id")

	staffIdInt, err := strconv.Atoi(staffId)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, "Invalid staff id")
	}

	appointments, statusCode, err := h.appointmentUsecase.FindAllByStaffId(staffIdInt)
	if err != nil {
		return NewErrorResponse(ctx, statusCode, err.Error())
	}

	return NewSuccessResponse(ctx, statusCode, "Appointments found", appointments)
}
