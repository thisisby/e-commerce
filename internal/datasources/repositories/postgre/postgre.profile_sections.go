package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
)

type postgreProfileSectionsRepository struct {
	conn *sqlx.DB
}

func NewPostgreProfileSectionsRepository(conn *sqlx.DB) domains.ProfileSectionsRepository {
	return &postgreProfileSectionsRepository{
		conn: conn,
	}
}

func (p *postgreProfileSectionsRepository) FindAll() ([]domains.ProfileSectionsDomain, error) {
	query := `
		SELECT id, name, content, parent_id
		FROM profile_sections
		`

	var profileSectionsRecords []records.ProfileSections

	err := p.conn.Select(&profileSectionsRecords, query)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToProfileSectionDomains(profileSectionsRecords), nil
}

func (p *postgreProfileSectionsRepository) Save(profileSection domains.ProfileSectionsDomain) error {
	query := `
		INSERT INTO profile_sections (name, content, parent_id)
		VALUES (:name, :content, :parent_id)
`
	profileSectionRecord := records.FromProfileSectionDomain(profileSection)

	_, err := p.conn.NamedQuery(query, profileSectionRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}
