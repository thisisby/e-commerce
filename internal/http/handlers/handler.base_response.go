package handlers

import "github.com/labstack/echo/v4"

type BaseResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Payload any    `json:"payload"`
}

func NewSuccessResponse(ctx echo.Context, statusCode int, message string, payload any) error {
	return ctx.JSON(statusCode, BaseResponse{
		Status:  statusCode,
		Message: message,
		Payload: payload,
	})
}

func NewErrorResponse(ctx echo.Context, statusCode int, message string) error {
	return ctx.JSON(statusCode, BaseResponse{
		Status:  statusCode,
		Message: message,
	})
}
