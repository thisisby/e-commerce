package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type ServiceItemRoute struct {
	serviceItemHandler handlers.ServiceItemHandler
	router             *echo.Group
	db                 *sqlx.DB
	authMiddleware     middlewares.AuthMiddleware
	adminMiddleware    middlewares.AuthMiddleware
}

func NewServiceItemRoute(
	db *sqlx.DB,
	router *echo.Group,
	authMiddleware middlewares.AuthMiddleware,
	adminMiddleware middlewares.AuthMiddleware,
) *ServiceItemRoute {

	serviceItemRepo := postgre.NewPostgreServiceItemRepository(db)
	serviceItemUsecase := usecases.NewServiceItemUsecase(serviceItemRepo)
	serviceItemHandler := handlers.NewServiceItemHandler(serviceItemUsecase)

	return &ServiceItemRoute{
		serviceItemHandler: *serviceItemHandler,
		router:             router,
		db:                 db,
		authMiddleware:     authMiddleware,
		adminMiddleware:    adminMiddleware,
	}
}

func (r *ServiceItemRoute) Register() {
	serviceItems := r.router.Group("/service_items")
	subservices := r.router.Group("/subservices")
	admin := r.router.Group("/admin/service_items")

	subservices.Use(r.authMiddleware.Handle)
	subservices.GET("/:subservice_id/service_items", r.serviceItemHandler.FindBySubserviceId)

	serviceItems.Use(r.authMiddleware.Handle)
	serviceItems.GET("", r.serviceItemHandler.FindAll)

	admin.Use(r.adminMiddleware.Handle)
	admin.POST("", r.serviceItemHandler.Save)
	admin.PATCH("/:id", r.serviceItemHandler.Update)
	admin.DELETE("/:id", r.serviceItemHandler.Delete)
}
