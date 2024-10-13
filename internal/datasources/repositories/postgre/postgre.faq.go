package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
)

type postgreFaqRepository struct {
	Conn *sqlx.DB
}

func NewPostgreFaqRepository(Conn *sqlx.DB) domains.FaqRepository {
	return &postgreFaqRepository{
		Conn: Conn,
	}
}

func (p *postgreFaqRepository) FindAll() ([]domains.FaqDomain, error) {
	query := `SELECT * FROM faq`

	var faq []records.Faq

	err := p.Conn.Select(&faq, query)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfFaqDomain(faq), nil
}

func (p *postgreFaqRepository) Save(domain domains.FaqDomain) error {
	query := `
		INSERT INTO faq (question, answer) 
		VALUES (:question, :answer)
	`

	faq := records.FromFaqDomain(domain)

	_, err := p.Conn.NamedQuery(query, faq)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreFaqRepository) Update(domain domains.FaqDomain, id int) error {
	query := `
		UPDATE faq 
		SET 
		    question = :question, 
			answer = :answer 
		WHERE id = :id
	`

	faq := records.FromFaqDomain(domain)
	faq.Id = id

	_, err := p.Conn.NamedQuery(query, faq)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreFaqRepository) Delete(id int) error {
	query := `DELETE FROM faq WHERE id = $1`

	_, err := p.Conn.Exec(query, id)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}
