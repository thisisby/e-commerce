package routes

import (
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"ga_marketplace/third_party/cdek"
	"github.com/labstack/echo/v4"
)

type DeliveryCalculatorRoute struct {
	deliveryCalculatorHandler handlers.DeliveryCalculatorHandler
	router                    *echo.Group
	authMiddleware            middlewares.AuthMiddleware
}

func NewDeliveryCalculatorRoute(
	route *echo.Group,
	authMiddleware middlewares.AuthMiddleware,
	cdekClient *cdek.Client,
) *DeliveryCalculatorRoute {

	deliveryCalculatorHandler := handlers.NewDeliveryCalculatorHandler(cdekClient)
	return &DeliveryCalculatorRoute{
		deliveryCalculatorHandler: deliveryCalculatorHandler,
		router:                    route,
		authMiddleware:            authMiddleware,
	}
}

func (r *DeliveryCalculatorRoute) Register() {
	deliveryCalculator := r.router.Group("/delivery-calculator")

	deliveryCalculator.Use(r.authMiddleware.Handle)
	deliveryCalculator.POST("", r.deliveryCalculatorHandler.Calculate)
}
