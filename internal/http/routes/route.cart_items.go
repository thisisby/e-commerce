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
	cartHandler     handlers.CartsHandler
	router          *echo.Group
	db              *sqlx.DB
	redisCache      caches.RedisCache
	authMiddleware  middlewares.AuthMiddleware
	adminMiddleware middlewares.AuthMiddleware
}

func NewCartsRoute(
	db *sqlx.DB,
	router *echo.Group,
	redisCache caches.RedisCache,
	authMiddleware middlewares.AuthMiddleware,
	adminMiddleware middlewares.AuthMiddleware,
) *CartsRoute {
	cartItemsRepo := postgre.NewPostgreCartsRepository(db)
	productsRepo := postgre.NewPostgreProductsRepository(db)
	usersRepo := postgre.NewPostgreUsersRepository(db)
	cartItemsUsecase := usecases.NewCartsUsecase(cartItemsRepo, usersRepo, productsRepo)
	cartItemsHandler := handlers.NewCartsHandler(cartItemsUsecase, redisCache)

	return &CartsRoute{
		cartHandler:     cartItemsHandler,
		router:          router,
		db:              db,
		redisCache:      redisCache,
		authMiddleware:  authMiddleware,
		adminMiddleware: adminMiddleware,
	}
}

func (r *CartsRoute) Register() {
	// carts
	cartItems := r.router.Group("/cart_items")
	users := r.router.Group("/users")

	cartItems.Use(r.authMiddleware.Handle)
	cartItems.GET("", r.cartHandler.FindAll)
	cartItems.POST("", r.cartHandler.SaveCart)
	cartItems.DELETE("/:id", r.cartHandler.DeleteCart)

	users.Use(r.adminMiddleware.Handle)
	users.GET("/:id/carts", r.cartHandler.GetCartsByUserId)

}
