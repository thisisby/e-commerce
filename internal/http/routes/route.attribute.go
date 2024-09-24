package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type AttributeRoute struct {
	attributeHandler handlers.AttributesHandler
	router           *echo.Group
	db               *sqlx.DB
	authMiddleware   middlewares.AuthMiddleware
	adminMiddleware  middlewares.AuthMiddleware
}

func NewAttributeRoute(
	db *sqlx.DB,
	router *echo.Group,
	authMiddleware middlewares.AuthMiddleware,
	adminMiddleware middlewares.AuthMiddleware,
) *AttributeRoute {

	attributeRepo := postgre.NewPostgreAttributesRepository(db)
	attributeUsecase := usecases.NewAttributesUsecase(attributeRepo)
	attributeHandler := handlers.NewAttributesHandler(attributeUsecase)

	return &AttributeRoute{
		attributeHandler: *attributeHandler,
		router:           router,
		db:               db,
		authMiddleware:   authMiddleware,
		adminMiddleware:  adminMiddleware,
	}
}

func (r *AttributeRoute) Register() {
	attributes := r.router.Group("/attributes")
	admin := r.router.Group("/admin/attributes")
	characteristics := r.router.Group("/characteristics/:characteristicId/attributes")

	attributes.GET("", r.attributeHandler.FindAll)
	characteristics.GET("", r.attributeHandler.FindAllByCharacteristicsId)

	admin.Use(r.adminMiddleware.Handle)
	admin.POST("", r.attributeHandler.Save)
	admin.PATCH("/:id", r.attributeHandler.Update)
	admin.DELETE("/:id", r.attributeHandler.Delete)
}
