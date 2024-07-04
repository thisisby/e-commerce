package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/caches"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type CartsRoute struct {
	cartItemsHandler handlers.CartItemsHandler
	router           *echo.Group
	db               *sqlx.DB
	redisCache       caches.RedisCache
	authMiddleware   middlewares.AuthMiddleware
}

func NewCartsRoute(
	db *sqlx.DB,
	router *echo.Group,
	redisCache caches.RedisCache,
	authMiddleware middlewares.AuthMiddleware,
) *CartsRoute {
	cartItemsRepo := postgre.NewPostgreCartsRepository(db)
	productsRepo := postgre.NewPostgreProductsRepository(db)
	usersRepo := postgre.NewPostgreUsersRepository(db)
	cartItemsUsecase := usecases.NewCartsUsecase(cartItemsRepo, usersRepo, productsRepo)
	cartItemsHandler := handlers.NewCartsHandler(cartItemsUsecase, redisCache)
	return &CartsRoute{
		cartItemsHandler: cartItemsHandler,
		router:           router,
		db:               db,
		redisCache:       redisCache,
		authMiddleware:   authMiddleware,
	}
}

func (r *CartsRoute) Register() {
	// carts
	cartItems := r.router.Group("/cart_items")

	cartItems.Use(r.authMiddleware.Handle)
	cartItems.GET("", r.cartItemsHandler.GetAllMyCartItems)
	cartItems.POST("", r.cartItemsHandler.SaveToMyCartItems)
	cartItems.DELETE("/:id", r.cartItemsHandler.DeleteMyCartItem)
	cartItems.PATCH("/:id", r.cartItemsHandler.UpdateMyCartItem)

	//users.Use(r.adminMiddleware.Handle)
	//users.GET("/:id/cart_items", r.cartItemsHandler.GetCartItemsByUserId)

}
