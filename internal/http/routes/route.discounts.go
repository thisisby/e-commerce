package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type DiscountRoute struct {
	discountHandler handlers.DiscountHandler
	router          *echo.Group
	db              *sqlx.DB
	adminMiddleware middlewares.AuthMiddleware
}

func NewDiscountRoute(
	db *sqlx.DB,
	router *echo.Group,
	adminMiddleware middlewares.AuthMiddleware,
) *DiscountRoute {
	discountRepo := postgre.NewPostgreDiscountRepository(db)
	discountUsecase := usecases.NewDiscountsUsecase(discountRepo)
	discountHandler := handlers.NewDiscountHandler(discountUsecase)

	return &DiscountRoute{
		discountHandler: discountHandler,
		router:          router,
		db:              db,
		adminMiddleware: adminMiddleware,
	}
}

func (r *DiscountRoute) Register() {
	discounts := r.router.Group("/discounts")
	products := r.router.Group("/products")

	discounts.Use(r.adminMiddleware.Handle)
	discounts.POST("", r.discountHandler.Save)

	products.Use(r.adminMiddleware.Handle)
	products.DELETE("/:product_id/discounts", r.discountHandler.DeleteByProductId)
}
