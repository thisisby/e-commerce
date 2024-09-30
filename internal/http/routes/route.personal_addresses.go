package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type PersonalAddresses struct {
	PersonalAddressesHandler handlers.PersonalAddressesHandler
	router                   *echo.Group
	db                       *sqlx.DB
	authMiddleware           middlewares.AuthMiddleware
	adminMiddleware          middlewares.AuthMiddleware
}

func NewPersonalAddressesRoute(
	db *sqlx.DB,
	router *echo.Group,
	authMiddleware middlewares.AuthMiddleware,
	adminMiddleware middlewares.AuthMiddleware,
) *PersonalAddresses {

	personalAddressRepo := postgre.NewPostgrePersonalAddresses(db)
	personalAddressUsecase := usecases.NewPersonalAddressesUsecase(personalAddressRepo)
	personalAddressHandler := handlers.NewPersonalAddressesHandler(personalAddressUsecase)

	return &PersonalAddresses{
		PersonalAddressesHandler: *personalAddressHandler,
		router:                   router,
		db:                       db,
		authMiddleware:           authMiddleware,
		adminMiddleware:          adminMiddleware,
	}
}

func (r *PersonalAddresses) Register() {
	personalAddresses := r.router.Group("/personal_addresses")
	admin := r.router.Group("/admin/personal_addresses")

	me := r.router.Group("/me/personal_addresses")

	personalAddresses.Use(r.authMiddleware.Handle)
	personalAddresses.GET("", r.PersonalAddressesHandler.FindAll)

	me.Use(r.authMiddleware.Handle)
	me.GET("", r.PersonalAddressesHandler.FindByUserId)

	admin.Use(r.adminMiddleware.Handle)
	admin.POST("", r.PersonalAddressesHandler.Save)
	admin.PUT("/:id", r.PersonalAddressesHandler.Update)
	admin.DELETE("/:id", r.PersonalAddressesHandler.Delete)
}
