package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"ga_marketplace/third_party/aws"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type ProductRoute struct {
	productHandler handlers.ProductHandler
	s3Client       *aws.S3Client
	router         *echo.Group
	db             *sqlx.DB
	authMiddleware middlewares.AuthMiddleware
}

func NewProductRoute(
	db *sqlx.DB,
	router *echo.Group,
	s3Client *aws.S3Client,
	authMiddleware middlewares.AuthMiddleware,
) *ProductRoute {
	productRepo := postgre.NewPostgreProductsRepository(db)
	productUsecase := usecases.NewProductsUsecase(productRepo)
	productHandler := handlers.NewProductHandler(productUsecase, s3Client)

	return &ProductRoute{
		productHandler: productHandler,
		router:         router,
		db:             db,
		s3Client:       s3Client,
		authMiddleware: authMiddleware,
	}
}

func (r *ProductRoute) Register() {
	products := r.router.Group("/products")

	products.Use(r.authMiddleware.Handle)
	products.POST("", r.productHandler.Save)
	products.GET("", r.productHandler.FindAllForMe)
}
