package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
)

type postgreContactsRepository struct {
	conn *sqlx.DB
}

func NewPostgreContactsRepository(conn *sqlx.DB) domains.ContactRepository {
	return &postgreContactsRepository{
		conn: conn,
	}
}

func (p *postgreContactsRepository) FindAll() ([]domains.ContactDomain, error) {
	query := `
		SELECT id, title, value
		FROM contacts
	`

	var contacts []records.Contacts

	err := p.conn.Select(&contacts, query)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfContactDomain(contacts), nil
}

func (p *postgreContactsRepository) Save(contact domains.ContactDomain) error {
	query := `INSERT INTO contacts (title, value) VALUES (:title, :value)`
	contactRecord := records.FromContactDomain(&contact)

	_, err := p.conn.NamedQuery(query, contactRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreContactsRepository) Update(contact domains.ContactDomain) error {
	query := `UPDATE contacts SET title = :title, value = :value WHERE id = :id`
	contactRecord := records.FromContactDomain(&contact)

	_, err := p.conn.NamedQuery(query, contactRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreContactsRepository) FindById(id int) (domains.ContactDomain, error) {
	query := `
		SELECT id, title, value
		FROM contacts
		WHERE id = $1
	`

	var contact records.Contacts

	err := p.conn.Get(&contact, query, id)
	if err != nil {
		return domains.ContactDomain{}, helpers.PostgresErrorTransform(err)
	}

	return *contact.ToDomain(), nil
}

func (p *postgreContactsRepository) Delete(id int) error {
	query := `DELETE FROM contacts WHERE id = $1`

	_, err := p.conn.Exec(query, id)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}
