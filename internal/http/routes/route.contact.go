package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type ContactsRoute struct {
	contactHandler  handlers.ContactHandler
	router          *echo.Group
	db              *sqlx.DB
	authMiddleware  middlewares.AuthMiddleware
	adminMiddleware middlewares.AuthMiddleware
}

func NewContactsRoute(
	db *sqlx.DB,
	router *echo.Group,
	authMiddleware middlewares.AuthMiddleware,
	adminMiddleware middlewares.AuthMiddleware,
) *ContactsRoute {

	contactRepo := postgre.NewPostgreContactsRepository(db)
	contactUsecase := usecases.NewContactsUsecase(contactRepo)
	contactHandler := handlers.NewContactHandler(contactUsecase)

	return &ContactsRoute{
		contactHandler:  *contactHandler,
		router:          router,
		db:              db,
		authMiddleware:  authMiddleware,
		adminMiddleware: adminMiddleware,
	}
}

func (r *ContactsRoute) Register() {
	contacts := r.router.Group("/contacts")
	admin := r.router.Group("/admin/contacts")

	contacts.Use(r.authMiddleware.Handle)
	contacts.GET("", r.contactHandler.FindAll)

	admin.Use(r.adminMiddleware.Handle)
	admin.POST("", r.contactHandler.Save)
	admin.PATCH("/:id", r.contactHandler.Update)
	admin.DELETE("/:id", r.contactHandler.Delete)
}
