package postgre

import (
	"ga_marketplace/internal/business/domains"
	"github.com/jmoiron/sqlx"
)

type postgreRolesRepository struct {
	conn *sqlx.DB
}

func NewPostgreRolesRepository(conn *sqlx.DB) domains.RoleRepository {
	return &postgreRolesRepository{
		conn: conn,
	}
}
