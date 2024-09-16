package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type BrandsRoute struct {
	brandHandler    handlers.BrandHandler
	router          *echo.Group
	db              *sqlx.DB
	authMiddleware  middlewares.AuthMiddleware
	adminMiddleware middlewares.AuthMiddleware
}

func NewBrandsRoute(
	db *sqlx.DB,
	router *echo.Group,
	authMiddleware middlewares.AuthMiddleware,
	adminMiddleware middlewares.AuthMiddleware,
) *BrandsRoute {

	brandRepo := postgre.NewPostgreBrandsRepository(db)
	brandUsecase := usecases.NewBrandsUsecase(brandRepo)
	brandHandler := handlers.NewBrandHandler(brandUsecase)

	return &BrandsRoute{
		brandHandler:    brandHandler,
		router:          router,
		db:              db,
		authMiddleware:  authMiddleware,
		adminMiddleware: adminMiddleware,
	}
}

func (r *BrandsRoute) Register() {
	brands := r.router.Group("/brands")
	admin := r.router.Group("/admin/brands")

	brands.GET("", r.brandHandler.FindAll)

	admin.Use(r.adminMiddleware.Handle)
	admin.POST("", r.brandHandler.Save)
	admin.PATCH("/:id", r.brandHandler.Update)
	admin.DELETE("/:id", r.brandHandler.Delete)
}
