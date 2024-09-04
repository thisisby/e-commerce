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
	productHandler   handlers.ProductHandler
	s3Client         *aws.S3Client
	router           *echo.Group
	db               *sqlx.DB
	authMiddleware   middlewares.AuthMiddleware
	adminMilddleware middlewares.AuthMiddleware
}

func NewProductRoute(
	db *sqlx.DB,
	router *echo.Group,
	s3Client *aws.S3Client,
	authMiddleware middlewares.AuthMiddleware,
	adminMiddleware middlewares.AuthMiddleware,
) *ProductRoute {
	productRepo := postgre.NewPostgreProductsRepository(db)
	productUsecase := usecases.NewProductsUsecase(productRepo)
	productHandler := handlers.NewProductHandler(productUsecase, s3Client)

	return &ProductRoute{
		productHandler:   productHandler,
		router:           router,
		db:               db,
		s3Client:         s3Client,
		authMiddleware:   authMiddleware,
		adminMilddleware: adminMiddleware,
	}
}

func (r *ProductRoute) Register() {
	products := r.router.Group("/products")
	productsBySubcategories := r.router.Group("/subcategories")
	productsByBrands := r.router.Group("/brands")
	admin := r.router.Group("/admin/products")

	products.Use(r.authMiddleware.Handle)
	products.GET("", r.productHandler.FindAllForMe)
	products.GET("/:id", r.productHandler.FindByIdForUser)

	productsBySubcategories.Use(r.authMiddleware.Handle)
	productsBySubcategories.GET("/:subcategory_id/products", r.productHandler.FindBySubCategoryId)

	productsByBrands.Use(r.authMiddleware.Handle)
	productsByBrands.GET("/:brand_id/products", r.productHandler.FindByBrandId)

	admin.Use(r.adminMilddleware.Handle)
	admin.POST("", r.productHandler.Save)
	admin.POST("/from1c", r.productHandler.SaveFrom1c)
	admin.PATCH("/from1c/:c_code", r.productHandler.UpdateFrom1c)
	admin.PATCH("/:id", r.productHandler.UpdateById)
}
