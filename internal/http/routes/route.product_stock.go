package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type ProductStockRoute struct {
	productStockHandler handlers.ProductStockHandler
	router              *echo.Group
	db                  *sqlx.DB
	authMiddleware      middlewares.AuthMiddleware
	adminMiddleware     middlewares.AuthMiddleware
}

func NewProductStockRoute(
	db *sqlx.DB,
	router *echo.Group,
	authMiddleware middlewares.AuthMiddleware,
	adminMiddleware middlewares.AuthMiddleware,
) *ProductStockRoute {
	productStockRepo := postgre.NewPostgreProductStockRepository(db)
	productStockUsecase := usecases.NewProductStockUsecase(productStockRepo)

	productStockHandler := handlers.NewProductStockHandler(productStockUsecase)

	return &ProductStockRoute{
		productStockHandler: productStockHandler,
		router:              router,
		db:                  db,
		authMiddleware:      authMiddleware,
		adminMiddleware:     adminMiddleware,
	}
}

func (r *ProductStockRoute) Register() {
	adminRoute := r.router.Group("/admin/product-stocks")

	adminRoute.Use(r.adminMiddleware.Handle)
	adminRoute.POST("", r.productStockHandler.Save)
	adminRoute.PATCH("/:id", r.productStockHandler.Update)
}
