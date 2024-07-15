package routes

import (
	"ga_marketplace/internal/http/handlers"
	"github.com/labstack/echo/v4"
)

type HealthCheckRoute struct {
	healthCheckHandler handlers.HealthCheckHandler
	router             *echo.Group
}

func NewHealthCheckRoute(router *echo.Group) *HealthCheckRoute {
	healthCheckHandler := handlers.NewHealthCheckHandler()

	return &HealthCheckRoute{
		healthCheckHandler: healthCheckHandler,
		router:             router,
	}
}

func (r *HealthCheckRoute) Register() {
	r.router.GET("/health-check", r.healthCheckHandler.HealthCheck)
}
