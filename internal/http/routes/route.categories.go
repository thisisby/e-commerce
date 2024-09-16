package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type CategoriesRoute struct {
	categoriesHandler handlers.CategoriesHandler
	router            *echo.Group
	db                *sqlx.DB
	authMiddleware    middlewares.AuthMiddleware
	adminMiddleware   middlewares.AuthMiddleware
}

func NewCategoriesRoute(
	db *sqlx.DB,
	router *echo.Group,
	authMiddleware middlewares.AuthMiddleware,
	adminMiddleware middlewares.AuthMiddleware,
) *CategoriesRoute {

	categoriesRepo := postgre.NewPostgreCategoriesRepository(db)
	categoriesUsecase := usecases.NewCategoriesUsecase(categoriesRepo)
	categoriesHandler := handlers.NewCategoriesHandler(categoriesUsecase)

	return &CategoriesRoute{
		categoriesHandler: categoriesHandler,
		router:            router,
		db:                db,
		authMiddleware:    authMiddleware,
		adminMiddleware:   adminMiddleware,
	}
}

func (r *CategoriesRoute) Register() {
	categories := r.router.Group("/categories")
	admin := r.router.Group("/admin/categories")

	categories.GET("", r.categoriesHandler.FindAll)

	admin.Use(r.adminMiddleware.Handle)
	admin.POST("", r.categoriesHandler.Save)
	admin.PATCH("/:id", r.categoriesHandler.Update)
	admin.DELETE("/:id", r.categoriesHandler.Delete)
}
