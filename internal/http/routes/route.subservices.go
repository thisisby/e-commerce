package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type SubservicesRoutes struct {
	SubServicesHandler handlers.SubServicesHandler
	router             *echo.Group
	db                 *sqlx.DB
	authMiddleware     middlewares.AuthMiddleware
	adminMiddleware    middlewares.AuthMiddleware
}

func NewSubservicesRoutes(
	db *sqlx.DB,
	router *echo.Group,
	authMiddleware middlewares.AuthMiddleware,
	adminMiddleware middlewares.AuthMiddleware,
) *SubservicesRoutes {

	subServicesRepo := postgre.NewPostgreSubservicesRepository(db)
	subServicesUsecase := usecases.NewSubServicesUsecase(subServicesRepo)
	subServicesHandler := handlers.NewSubServicesHandler(subServicesUsecase)

	return &SubservicesRoutes{
		SubServicesHandler: *subServicesHandler,
		router:             router,
		db:                 db,
		authMiddleware:     authMiddleware,
		adminMiddleware:    adminMiddleware,
	}
}

func (r *SubservicesRoutes) Register() {
	services := r.router.Group("/services")
	subservices := r.router.Group("/subservices")
	admin := r.router.Group("/admin/subservices")

	subservices.Use(r.authMiddleware.Handle)
	subservices.GET("", r.SubServicesHandler.FindAll)

	services.Use(r.authMiddleware.Handle)
	services.GET("/:service_id/subservices", r.SubServicesHandler.FindAllByServiceId)

	admin.Use(r.adminMiddleware.Handle)
	admin.POST("", r.SubServicesHandler.Save)
	admin.PATCH("/:id", r.SubServicesHandler.Update)
	admin.DELETE("/:id", r.SubServicesHandler.Delete)
}
