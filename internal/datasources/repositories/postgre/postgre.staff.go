package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
)

type postgreStaffRepository struct {
	conn *sqlx.DB
}

func NewPostgreStaffRepository(conn *sqlx.DB) domains.StaffRepository {
	return &postgreStaffRepository{
		conn: conn,
	}
}

func (p *postgreStaffRepository) FindById(id int) (*domains.StaffDomain, error) {
	query := `SELECT * FROM staff WHERE id = $1`

	var staffRecord records.StaffRecord

	err := p.conn.Get(&staffRecord, query, id)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return staffRecord.ToDomain(), nil
}

func (p *postgreStaffRepository) Save(staff *domains.StaffDomain) error {
	query := `INSERT INTO staff (full_name, occupation, experience, avatar, service_id) VALUES (:full_name, :occupation, :experience, :avatar, :service_id)`

	staffRecord := records.FromStaffDomain(staff)

	_, err := p.conn.NamedQuery(query, staffRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreStaffRepository) FindAll() ([]domains.StaffDomain, error) {
	query := `SELECT * FROM staff`

	var staffRecords []records.StaffRecord

	err := p.conn.Select(&staffRecords, query)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfStaffDomain(staffRecords), nil
}

func (p *postgreStaffRepository) Update(inDom domains.StaffDomain) error {
	query := `UPDATE staff SET full_name = :full_name, occupation = :occupation, experience = :experience, avatar = :avatar, service_id = :service_id WHERE id = :id`

	staffRecord := records.FromStaffDomain(&inDom)

	_, err := p.conn.NamedQuery(query, staffRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreStaffRepository) Delete(id int) error {
	query := `DELETE FROM staff WHERE id = $1`

	_, err := p.conn.Exec(query, id)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreStaffRepository) FindByServiceId(serviceId int) ([]domains.StaffDomain, error) {
	query := `SELECT * FROM staff WHERE service_id = $1`

	var staffRecords []records.StaffRecord

	err := p.conn.Select(&staffRecords, query, serviceId)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfStaffDomain(staffRecords), nil
}
