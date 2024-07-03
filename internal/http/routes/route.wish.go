package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type WishRoute struct {
	wishHandler    handlers.WishHandler
	router         *echo.Group
	db             *sqlx.DB
	authMiddleware middlewares.AuthMiddleware
}

func NewWishRoute(
	db *sqlx.DB,
	router *echo.Group,
	authMiddleware middlewares.AuthMiddleware,
) *WishRoute {
	wishRepo := postgre.NewPostgreWishRepository(db)
	wishUsecase := usecases.NewWishUsecase(wishRepo)
	wishHandler := handlers.NewWishHandler(wishUsecase)

	return &WishRoute{
		wishHandler:    wishHandler,
		router:         router,
		db:             db,
		authMiddleware: authMiddleware,
	}
}

func (r *WishRoute) Register() {
	wishes := r.router.Group("/wishes")

	wishes.Use(r.authMiddleware.Handle)
	wishes.GET("", r.wishHandler.FindAll)
	wishes.POST("", r.wishHandler.SaveWish)
	wishes.DELETE("/:id", r.wishHandler.DeleteWish)
}
