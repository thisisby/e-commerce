package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type CharacteristicsRoute struct {
	characteristicsHandler handlers.CharacteristicsHandler
	router                 *echo.Group
	db                     *sqlx.DB
	authMiddleware         middlewares.AuthMiddleware
	adminMiddleware        middlewares.AuthMiddleware
}

func NewCharacteristicsRoute(
	db *sqlx.DB,
	router *echo.Group,
	authMiddleware middlewares.AuthMiddleware,
	adminMiddleware middlewares.AuthMiddleware,
) *CharacteristicsRoute {

	characteristicsRepo := postgre.NewPostgreCharacteristicsRepository(db)
	characteristicsUsecase := usecases.NewCharacteristicsUsecase(characteristicsRepo)
	characteristicsHandler := handlers.NewCharacteristicsHandler(characteristicsUsecase)

	return &CharacteristicsRoute{
		characteristicsHandler: *characteristicsHandler,
		router:                 router,
		db:                     db,
		authMiddleware:         authMiddleware,
		adminMiddleware:        adminMiddleware,
	}
}

func (r *CharacteristicsRoute) Register() {
	characteristics := r.router.Group("/characteristics")
	admin := r.router.Group("/admin/characteristics")
	subcategory := r.router.Group("/subcategories/:subcategoryId/characteristics")

	characteristics.GET("", r.characteristicsHandler.FindAll)
	subcategory.GET("", r.characteristicsHandler.FindAllBySubcategoryId)

	admin.Use(r.adminMiddleware.Handle)
	admin.POST("", r.characteristicsHandler.Save)
	admin.PATCH("/:id", r.characteristicsHandler.Update)
	admin.DELETE("/:id", r.characteristicsHandler.Delete)
}
