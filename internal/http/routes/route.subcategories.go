package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type SubcategoriesRoute struct {
	subcategoriesHandler handlers.SubcategoriesHandler
	router               *echo.Group
	db                   *sqlx.DB
	authMiddleware       middlewares.AuthMiddleware
	adminMiddleware      middlewares.AuthMiddleware
}

func NewSubcategoriesRoute(
	db *sqlx.DB,
	router *echo.Group,
	authMiddleware middlewares.AuthMiddleware,
	adminMiddleware middlewares.AuthMiddleware,
) *SubcategoriesRoute {

	subcategoriesRepo := postgre.NewPostgreSubcategoriesRepository(db)
	subcategoriesUsecase := usecases.NewSubcategoriesUsecase(subcategoriesRepo)
	subcategoriesHandler := handlers.NewSubcategoriesHandler(subcategoriesUsecase)

	return &SubcategoriesRoute{
		subcategoriesHandler: subcategoriesHandler,
		router:               router,
		db:                   db,
		authMiddleware:       authMiddleware,
		adminMiddleware:      adminMiddleware,
	}
}

func (r *SubcategoriesRoute) Register() {
	subcategories := r.router.Group("/subcategories")
	categories := r.router.Group("/categories")
	admin := r.router.Group("/admin/subcategories")

	subcategories.GET("", r.subcategoriesHandler.FindAll)

	categories.GET("/:category_id/subcategories", r.subcategoriesHandler.FindByCategoryId)

	admin.Use(r.adminMiddleware.Handle)
	admin.POST("", r.subcategoriesHandler.Save)
	admin.PATCH("/:id", r.subcategoriesHandler.Update)
	admin.DELETE("/:id", r.subcategoriesHandler.Delete)
}
