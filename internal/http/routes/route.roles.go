package routes

import (
	"ga_marketplace/internal/business/usecases"
	"ga_marketplace/internal/datasources/repositories/postgre"
	"ga_marketplace/internal/http/handlers"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type RolesRoute struct {
	rolesHandler handlers.RolesHandler
	router       *echo.Group
	db           *sqlx.DB
}

func NewRolesRoute(db *sqlx.DB, router *echo.Group) *RolesRoute {
	rolesRepo := postgre.NewPostgreRolesRepository(db)
	rolesUsecase := usecases.NewRolesUsecase(rolesRepo)
	rolesHandler := handlers.NewRolesHandler(rolesUsecase)

	return &RolesRoute{
		rolesHandler: rolesHandler,
		router:       router,
		db:           db,
	}
}

func (r *RolesRoute) Register() {

}
