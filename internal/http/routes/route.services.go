package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type ServicesRoute struct {
	servicesHandler handlers.ServicesHandler
	router          *echo.Group
	db              *sqlx.DB
	authMiddleware  middlewares.AuthMiddleware
	adminMiddleware middlewares.AuthMiddleware
}

func NewServicesRoute(
	db *sqlx.DB,
	router *echo.Group,
	authMiddleware middlewares.AuthMiddleware,
	adminMiddleware middlewares.AuthMiddleware,
) *ServicesRoute {

	servicesRepo := postgre.NewPostgreServicesRepository(db)
	servicesUsecase := usecases.NewServicesUsecase(servicesRepo)
	servicesHandler := handlers.NewServicesHandler(servicesUsecase)

	return &ServicesRoute{
		servicesHandler: *servicesHandler,
		router:          router,
		db:              db,
		authMiddleware:  authMiddleware,
		adminMiddleware: adminMiddleware,
	}
}

func (r *ServicesRoute) Register() {
	services := r.router.Group("/services")
	admin := r.router.Group("/admin/services")

	services.Use(r.authMiddleware.Handle)
	services.GET("", r.servicesHandler.FindAll)

	admin.Use(r.adminMiddleware.Handle)
	admin.POST("", r.servicesHandler.Save)
	admin.PATCH("/:id", r.servicesHandler.Update)
	admin.DELETE("/:id", r.servicesHandler.Delete)
}
