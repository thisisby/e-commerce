package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type OrdersRoute struct {
	orderHandler    handlers.OrdersHandler
	router          *echo.Group
	db              *sqlx.DB
	authMiddleware  middlewares.AuthMiddleware
	adminMiddleware middlewares.AuthMiddleware
}

func NewOrdersRoute(
	db *sqlx.DB,
	router *echo.Group,
	authMiddleware middlewares.AuthMiddleware,
	adminMiddleware middlewares.AuthMiddleware,
) *OrdersRoute {

	cartRepo := postgre.NewPostgreCartsRepository(db)
	productRepo := postgre.NewPostgreProductsRepository(db)
	productStockRepo := postgre.NewPostgreProductStockRepository(db)
	userRepo := postgre.NewPostgreUsersRepository(db)
	cartItemsUsecase := usecases.NewCartsUsecase(cartRepo, userRepo, productRepo)
	ordersRepo := postgre.NewPostgreOrdersRepository(db)
	ordersUsecase := usecases.NewOrdersUsecase(ordersRepo, productStockRepo)
	orderHandler := handlers.NewOrdersHandler(ordersUsecase, cartItemsUsecase)

	return &OrdersRoute{
		orderHandler:    orderHandler,
		router:          router,
		db:              db,
		authMiddleware:  authMiddleware,
		adminMiddleware: adminMiddleware,
	}
}

func (r *OrdersRoute) Register() {
	orders := r.router.Group("/orders")
	admin := r.router.Group("/admin/orders")

	orders.Use(r.authMiddleware.Handle)
	orders.POST("", r.orderHandler.Save)
	orders.GET("", r.orderHandler.FindMyOrders)
	orders.POST("/:id/cancel", r.orderHandler.Cancel)

	admin.Use(r.adminMiddleware.Handle)
	admin.GET("", r.orderHandler.FindAll)
	admin.PATCH("/:id", r.orderHandler.Update)

}
