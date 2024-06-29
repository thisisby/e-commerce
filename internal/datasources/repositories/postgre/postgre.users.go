package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type postgreUsersRepository struct {
	conn *sqlx.DB
}

func NewPostgreUsersRepository(conn *sqlx.DB) domains.UserRepository {
	return &postgreUsersRepository{
		conn: conn,
	}
}

func (p *postgreUsersRepository) FindByPhone(phone string) (*domains.UserDomain, error) {
	query := `
		SELECT u.id, u.name, u.phone, r.name "role.name", u.date_of_birth, u.created_at, u.updated_at 
		FROM users u 
		INNER JOIN roles r ON u.role_id = r.id 
		WHERE phone = $1
		`

	var userRecord records.Users

	err := p.conn.Get(&userRecord, query, phone)
	if err != nil {
		slog.Error("PostgreUsersRepository.FindByPhone: ", err)
		return nil, helpers.PostgresErrorTransform(err)
	}

	return userRecord.ToDomain(), nil
}

func (p *postgreUsersRepository) Save(inDom *domains.UserDomain) error {
	query := `
		INSERT INTO users (name, phone, role_id, date_of_birth, created_at)
		VALUES (:name, :phone, :role_id, :date_of_birth, :created_at)
		`

	userRecord := records.FromUsersDomain(inDom)

	_, err := p.conn.NamedQuery(query, userRecord)
	if err != nil {
		slog.Error("PostgreUsersRepository.Save: ", err)
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreUsersRepository) Update(inDom *domains.UserDomain) error {
	query := `
		UPDATE users 
		SET name = :name, phone = :phone, role_id = :role_id, date_of_birth = :date_of_birth, refresh_token = :refresh_token, updated_at = :updated_at
		WHERE id = :id
		`

	userRecord := records.FromUsersDomain(inDom)

	_, err := p.conn.NamedQuery(query, userRecord)
	if err != nil {
		slog.Error("PostgreUsersRepository.Update: ", err)
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreUsersRepository) FindById(id int) (*domains.UserDomain, error) {
	query := `
		SELECT u.id, u.name, u.phone, r.name "role.name", u.refresh_token, u.date_of_birth, u.created_at, u.updated_at
		FROM users u
		INNER JOIN roles r ON u.role_id = r.id
		WHERE u.id = $1
		`

	var userRecord records.Users

	err := p.conn.Get(&userRecord, query, id)
	if err != nil {
		slog.Error("PostgreUsersRepository.FindById: ", err)
		return nil, helpers.PostgresErrorTransform(err)
	}

	return userRecord.ToDomain(), nil
}
