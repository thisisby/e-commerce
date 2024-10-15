package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type FaqRouter struct {
	FaqHandler      handlers.FaqHandler
	router          *echo.Group
	db              *sqlx.DB
	authMiddleware  middlewares.AuthMiddleware
	adminMiddleware middlewares.AuthMiddleware
}

func NewFaqRouter(
	db *sqlx.DB,
	router *echo.Group,
	authMiddleware middlewares.AuthMiddleware,
	adminMiddleware middlewares.AuthMiddleware,
) *FaqRouter {

	faqRepo := postgre.NewPostgreFaqRepository(db)
	faqUsecase := usecases.NewFaqUsecase(faqRepo)
	faqHandler := handlers.NewFaqHandler(faqUsecase)

	return &FaqRouter{
		FaqHandler:      *faqHandler,
		router:          router,
		db:              db,
		authMiddleware:  authMiddleware,
		adminMiddleware: adminMiddleware,
	}
}

func (r *FaqRouter) Register() {
	faq := r.router.Group("/faq")
	admin := r.router.Group("/admin/faq")

	faq.GET("", r.FaqHandler.FindAll)

	admin.Use(r.adminMiddleware.Handle)
	admin.POST("", r.FaqHandler.Save)
	admin.PUT("/:id", r.FaqHandler.Update)
	admin.DELETE("/:id", r.FaqHandler.Delete)
}
