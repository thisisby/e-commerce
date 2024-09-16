package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"ga_marketplace/internal/http/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type ProfileSectionsRoute struct {
	profileSectionsHandler handlers.ProfileSectionsHandler
	router                 *echo.Group
	db                     *sqlx.DB
	authMiddleware         middlewares.AuthMiddleware
	adminAuthMiddleware    middlewares.AuthMiddleware
}

func NewProfileSectionsRoute(
	db *sqlx.DB,
	router *echo.Group,
	authMiddleware middlewares.AuthMiddleware,
	adminAuthMiddleware middlewares.AuthMiddleware,
) *ProfileSectionsRoute {
	profileSectionRepo := postgre.NewPostgreProfileSectionsRepository(db)
	profileSectionsUsecase := usecases.NewProfileSectionsUsecase(profileSectionRepo)
	profileSectionsHandler := handlers.NewProfileSectionsHandler(profileSectionsUsecase)

	return &ProfileSectionsRoute{
		profileSectionsHandler: profileSectionsHandler,
		router:                 router,
		db:                     db,
		authMiddleware:         authMiddleware,
		adminAuthMiddleware:    adminAuthMiddleware,
	}
}

func (r *ProfileSectionsRoute) Register() {
	profileSections := r.router.Group("/profile-sections")

	profileSections.GET("", r.profileSectionsHandler.FindAll)

	admin := r.router.Group("/admin/profile-sections")

	admin.Use(r.adminAuthMiddleware.Handle)
	admin.POST("", r.profileSectionsHandler.Save)
	admin.PATCH("/:id", r.profileSectionsHandler.UpdateById)
	admin.DELETE("/:id", r.profileSectionsHandler.DeleteById)
}
