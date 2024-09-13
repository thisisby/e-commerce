package postgre

import (
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/datasources/records"
	"ga_marketplace/pkg/helpers"
	"github.com/jmoiron/sqlx"
	"time"
)

type postgreAppointmentsRepository struct {
	conn *sqlx.DB
}

func NewPostgreAppointmentsRepository(conn *sqlx.DB) domains.AppointmentRepository {
	return &postgreAppointmentsRepository{
		conn: conn,
	}
}

func (p *postgreAppointmentsRepository) FindAll() ([]domains.AppointmentDomain, error) {
	query := `
		SELECT
			a.id, a.user_id, a.staff_id, a.start_time, a.end_time, a.service_item_id, a.comments, a.status, a.full_name, a.phone_number,
			s.id "staff.id", s.full_name "staff.full_name", s.occupation "staff.occupation", s.experience "staff.experience", s.avatar "staff.avatar", s.service_id "staff.service_id", s.service_address_id "staff.service_address_id",
			si.id "service_item.id", si.title "service_item.title", si.price "service_item.price", si.duration "service_item.duration", si.description "service_item.description", si.subservice_id "service_item.subservice_id"
		FROM appointments a	
		JOIN staff s ON a.staff_id = s.id
		JOIN service_items si ON a.service_item_id = si.id
	`

	var appointmentRecord []records.Appointment
	err := p.conn.Select(&appointmentRecord, query)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfAppointmentDomain(appointmentRecord), nil
}

func (p *postgreAppointmentsRepository) FindByUserId(userId int) ([]domains.AppointmentDomain, error) {
	query := `
		SELECT
			a.id, a.user_id, a.staff_id, a.start_time, a.end_time, a.service_item_id, a.comments, a.status, a.full_name, a.phone_number,
			s.id "staff.id", s.full_name "staff.full_name", s.occupation "staff.occupation", s.experience "staff.experience", s.avatar "staff.avatar", s.service_id "staff.service_id", s.service_address_id "staff.service_address_id",
			si.id "service_item.id", si.title "service_item.title", si.price "service_item.price", si.duration "service_item.duration", si.description "service_item.description", si.subservice_id "service_item.subservice_id"
		FROM appointments a	
		JOIN users u ON a.user_id = u.id
		JOIN staff s ON a.staff_id = s.id
		JOIN service_items si ON a.service_item_id = si.id
		WHERE a.user_id = $1	
	`

	var appointmentRecord []records.Appointment
	err := p.conn.Select(&appointmentRecord, query, userId)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfAppointmentDomain(appointmentRecord), nil
}

func (p *postgreAppointmentsRepository) Save(domain domains.AppointmentDomain) error {
	query := `
		INSERT INTO appointments (user_id, staff_id, start_time, end_time, service_item_id, comments, status, full_name, phone_number)
		VALUES (:user_id, :staff_id, :start_time, :end_time, :service_item_id, :comments, :status, :full_name, :phone_number)
	`

	appointmentRecord := records.FromAppointmentDomain(domain)
	_, err := p.conn.NamedExec(query, appointmentRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreAppointmentsRepository) Update(domain domains.AppointmentDomain) error {
	query := `
		UPDATE appointments
		SET 
		    user_id = :user_id, 
		    staff_id = :staff_id, 
		    start_time = :start_time, 
		    end_time = :end_time, 
		    service_item_id = :service_item_id,
		    comments = :comments, 
		    status = :status,
		    full_name = :full_name,
		    phone_number = :phone_number
		WHERE id = :id
	`

	appointmentRecord := records.FromAppointmentDomain(domain)

	_, err := p.conn.NamedExec(query, appointmentRecord)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreAppointmentsRepository) Delete(id int) error {
	query := `
		DELETE FROM appointments
		WHERE id = $1
	`

	_, err := p.conn.Exec(query, id)
	if err != nil {
		return helpers.PostgresErrorTransform(err)
	}

	return nil
}

func (p *postgreAppointmentsRepository) FindById(id int) (domains.AppointmentDomain, error) {
	query := `
		SELECT 
			a.id, a.user_id, a.staff_id, a.start_time, a.end_time, a.service_item_id, a.comments, a.status, a.full_name, a.phone_number,
			si.id "service_item.id", si.title "service_item.title", si.price "service_item.price", si.duration "service_item.duration", si.description "service_item.description", si.subservice_id "service_item.subservice_id"
		FROM appointments a
		JOIN service_items si ON a.service_item_id = si.id
		WHERE a.id = $1
`

	var appointmentRecord records.Appointment
	err := p.conn.Get(&appointmentRecord, query, id)
	if err != nil {
		return domains.AppointmentDomain{}, helpers.PostgresErrorTransform(err)
	}

	return *appointmentRecord.ToDomain(), nil
}

func (p *postgreAppointmentsRepository) IsOverlapping(appointmentId int, staffId int, startTime time.Time, endTime time.Time) (bool, error) {
	query := `
			SELECT COUNT(*) 
			FROM appointments 
			WHERE 
				staff_id = $1 AND 
				start_time < $3 AND
				end_time > $2 AND
				id != $4
	`

	var count int
	err := p.conn.QueryRow(query, staffId, startTime, endTime, appointmentId).Scan(&count)
	if err != nil {
		return false, helpers.PostgresErrorTransform(err)
	}

	return count > 0, nil
}

func (p *postgreAppointmentsRepository) FindAllByStaffId(staffId int) ([]domains.AppointmentDomain, error) {
	query := `
		SELECT
			a.id, a.user_id, a.staff_id, a.start_time, a.end_time, a.service_item_id, a.comments, a.status, a.full_name, a.phone_number,
			si.id "service_item.id", si.title "service_item.title", si.price "service_item.price", si.duration "service_item.duration", si.description "service_item.description", si.subservice_id "service_item.subservice_id"
		FROM appointments a	
		JOIN service_items si ON a.service_item_id = si.id
		WHERE a.staff_id = $1	
	`

	var appointmentRecord []records.Appointment
	err := p.conn.Select(&appointmentRecord, query, staffId)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfAppointmentDomain(appointmentRecord), nil
}

func (p *postgreAppointmentsRepository) FindAllByStaffIdAndDate(staffId int, date time.Time) ([]domains.AppointmentDomain, error) {
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 999999999, date.Location())

	query := `
		SELECT
			a.id, a.user_id, a.staff_id, a.start_time, a.end_time, a.service_item_id, a.comments, a.status, a.full_name, a.phone_number,
			s.id "staff.id", s.full_name "staff.full_name", s.occupation "staff.occupation", s.experience "staff.experience", s.avatar "staff.avatar", s.service_id "staff.service_id", s.service_address_id "staff.service_address_id"
		FROM appointments a	
		JOIN staff s ON a.staff_id = s.id
		WHERE a.staff_id = $1
		AND a.start_time >= $2 AND a.start_time <= $3
	`

	var appointmentRecord []records.Appointment
	err := p.conn.Select(&appointmentRecord, query, staffId, startOfDay, endOfDay)
	if err != nil {
		return nil, helpers.PostgresErrorTransform(err)
	}

	return records.ToArrayOfAppointmentDomain(appointmentRecord), nil
}
