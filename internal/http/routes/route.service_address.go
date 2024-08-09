package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type ServiceAddressRoute struct {
	serviceAddressHandler handlers.ServiceAddressHandler
	router                *echo.Group
	db                    *sqlx.DB
	authMiddleware        middlewares.AuthMiddleware
	adminMiddleware       middlewares.AuthMiddleware
}

func NewServiceAddressRoute(
	db *sqlx.DB,
	router *echo.Group,
	authMiddleware middlewares.AuthMiddleware,
	adminMiddleware middlewares.AuthMiddleware,
) *ServiceAddressRoute {

	serviceAddressRepo := postgre.NewPostgreServiceAddressRepository(db)
	serviceAddressUsecase := usecases.NewServiceAddressUsecase(serviceAddressRepo)
	serviceAddressHandler := handlers.NewServiceAddressHandler(serviceAddressUsecase)

	return &ServiceAddressRoute{
		serviceAddressHandler: *serviceAddressHandler,
		router:                router,
		db:                    db,
		authMiddleware:        authMiddleware,
		adminMiddleware:       adminMiddleware,
	}
}

func (r *ServiceAddressRoute) Register() {
	serviceAddress := r.router.Group("/service-address")
	admin := r.router.Group("/admin/service-address")
	cities := r.router.Group("/cities")

	serviceAddress.Use(r.authMiddleware.Handle)
	serviceAddress.GET("", r.serviceAddressHandler.FindAll)

	cities.Use(r.authMiddleware.Handle)
	cities.GET("/:city_id/service-address", r.serviceAddressHandler.FindByCityId)

	admin.Use(r.adminMiddleware.Handle)
	admin.POST("", r.serviceAddressHandler.Save)
	admin.PATCH("/:id", r.serviceAddressHandler.Update)
	admin.DELETE("/:id", r.serviceAddressHandler.Delete)
}
