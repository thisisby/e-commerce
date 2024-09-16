package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type CitiesRoute struct {
	cityHandler     handlers.CitiesHandler
	router          *echo.Group
	db              *sqlx.DB
	authMiddleware  middlewares.AuthMiddleware
	adminMiddleware middlewares.AuthMiddleware
}

func NewCitiesRoute(
	db *sqlx.DB,
	router *echo.Group,
	authMiddleware middlewares.AuthMiddleware,
	adminMiddleware middlewares.AuthMiddleware,
) *CitiesRoute {

	cityRepo := postgre.NewPostgreCityRepository(db)
	cityUsecase := usecases.NewCitiesUsecase(cityRepo)
	cityHandler := handlers.NewCitiesHandler(cityUsecase)

	return &CitiesRoute{
		cityHandler:     cityHandler,
		router:          router,
		db:              db,
		authMiddleware:  authMiddleware,
		adminMiddleware: adminMiddleware,
	}
}

func (r *CitiesRoute) Register() {
	cities := r.router.Group("/cities")
	admin := r.router.Group("/admin/cities")

	cities.GET("", r.cityHandler.FindAll)
	cities.GET("/:id", r.cityHandler.FindById)

	admin.Use(r.adminMiddleware.Handle)
	admin.POST("", r.cityHandler.Save)
	admin.PATCH("/:id", r.cityHandler.Update)
	admin.DELETE("/:id", r.cityHandler.Delete)
}
