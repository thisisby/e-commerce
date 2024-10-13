package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type FilialAddress struct {
	FilialAddressesHandler handlers.FilialAddressesHandler
	router                 *echo.Group
	db                     *sqlx.DB
	authMiddleware         middlewares.AuthMiddleware
	adminMiddleware        middlewares.AuthMiddleware
}

func NewFilialAddressesRoute(
	db *sqlx.DB,
	router *echo.Group,
	authMiddleware middlewares.AuthMiddleware,
	adminMiddleware middlewares.AuthMiddleware,
) *FilialAddress {

	filialAddressRepo := postgre.NewPostgreFilialAddresses(db)
	filialAddressUsecase := usecases.NewFilialAddressesUsecase(filialAddressRepo)
	filialAddressHandler := handlers.NewFilialAddressesHandler(filialAddressUsecase)

	return &FilialAddress{
		FilialAddressesHandler: *filialAddressHandler,
		router:                 router,
		db:                     db,
		authMiddleware:         authMiddleware,
		adminMiddleware:        adminMiddleware,
	}
}

func (r *FilialAddress) Register() {
	personalAddresses := r.router.Group("/filial_addresses")
	admin := r.router.Group("/admin/filial_addresses")

	personalAddresses.Use(r.authMiddleware.Handle)
	personalAddresses.GET("", r.FilialAddressesHandler.FindAll)

	admin.Use(r.adminMiddleware.Handle)
	admin.POST("", r.FilialAddressesHandler.Save)
	admin.PUT("/:id", r.FilialAddressesHandler.Update)
	admin.DELETE("/:id", r.FilialAddressesHandler.Delete)
}
