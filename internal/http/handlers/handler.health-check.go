package handlers

import "github.com/labstack/echo/v4"

type HealthCheckHandler struct {
}

func NewHealthCheckHandler() HealthCheckHandler {
	return HealthCheckHandler{}
}

func (h *HealthCheckHandler) HealthCheck(ctx echo.Context) error {
	return NewSuccessResponse(ctx, 200, "OK", nil)
}
