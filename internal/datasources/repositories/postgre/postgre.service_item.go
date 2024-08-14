package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
)

type postgreServiceItemRepository struct {
	conn *sqlx.DB
}

func NewPostgreServiceItemRepository(conn *sqlx.DB) domains.ServiceItemRepository {
	return &postgreServiceItemRepository{
		conn: conn,
	}
}

func (p *postgreServiceItemRepository) FindAll() ([]domains.ServiceItemDomain, error) {
	query := `SELECT 
				id,
				title,
				duration,	
				description,	
				price,	
				subservice_id
			   FROM 
				service_items
			`

	var serviceItems []records.ServiceItem
	err := p.conn.Select(&serviceItems, query)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfServiceItemDomain(serviceItems), nil
}

func (p *postgreServiceItemRepository) FindById(id int) (domains.ServiceItemDomain, error) {
	query := `SELECT 
				id,
				title,
				duration,	
				description,	
				price,	
				subservice_id
			   FROM 
				service_items
			   WHERE 
				id = $1
			`

	var serviceItem records.ServiceItem
	err := p.conn.Get(&serviceItem, query, id)
	if err != nil {
		return domains.ServiceItemDomain{}, helpers.PostgresErrorTransform(err)
	}

	return *serviceItem.ToDomain(), nil
}

func (p *postgreServiceItemRepository) FindBySubServiceId(subServiceId int) ([]domains.ServiceItemDomain, error) {
	query := `SELECT 
				id,
				title,
				duration,	
				description,	
				price,	
				subservice_id
			   FROM 
				service_items
			   WHERE 
				subservice_id = $1
			`

	var serviceItems []records.ServiceItem
	err := p.conn.Select(&serviceItems, query, subServiceId)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfServiceItemDomain(serviceItems), nil
}

func (p *postgreServiceItemRepository) Update(serviceItem domains.ServiceItemDomain) error {
	query := `	UPDATE service_items 
				SET 
                         title = :title,
                         duration = :duration, 
                         description = :description,
                         price = :price, 
                         subservice_id = :subservice_id
                     WHERE 
                         id = :id`

	serviceItemRecord := records.FromServiceItemDomain(serviceItem)
	_, err := p.conn.NamedQuery(query, serviceItemRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreServiceItemRepository) Delete(id int) error {
	query := `DELETE FROM service_items WHERE id = $1`

	_, err := p.conn.Exec(query, id)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreServiceItemRepository) Save(serviceItem domains.ServiceItemDomain) error {
	query := `INSERT INTO service_items (title, duration, description, price, subservice_id) VALUES (:title, :duration, :description, :price, :subservice_id)`

	serviceItemRecord := records.FromServiceItemDomain(serviceItem)
	_, err := p.conn.NamedQuery(query, serviceItemRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}
