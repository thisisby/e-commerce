package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type CountriesRoute struct {
	countryHandler  handlers.CountriesHandler
	router          *echo.Group
	db              *sqlx.DB
	authMiddleware  middlewares.AuthMiddleware
	adminMiddleware middlewares.AuthMiddleware
}

func NewCountriesRoute(
	db *sqlx.DB,
	router *echo.Group,
	authMiddleware middlewares.AuthMiddleware,
	adminMiddleware middlewares.AuthMiddleware,
) *CountriesRoute {

	countryRepo := postgre.NewPostgreCountriesRepository(db)
	countryUsecase := usecases.NewCountriesUsecase(countryRepo)
	countryHandler := handlers.NewCountriesHandler(countryUsecase)

	return &CountriesRoute{
		countryHandler:  countryHandler,
		router:          router,
		db:              db,
		authMiddleware:  authMiddleware,
		adminMiddleware: adminMiddleware,
	}
}

func (r *CountriesRoute) Register() {
	countries := r.router.Group("/countries")
	admin := r.router.Group("/admin/countries")

	countries.GET("", r.countryHandler.FindAll)
	countries.GET("/:id", r.countryHandler.FindById)

	admin.Use(r.adminMiddleware.Handle)
	admin.POST("", r.countryHandler.Save)
	admin.PATCH("/:id", r.countryHandler.Update)
	admin.DELETE("/:id", r.countryHandler.Delete)
}
