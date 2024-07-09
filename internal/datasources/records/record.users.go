package records

import "time"

type Users struct {
	Id           int       `db:"id"`
	Name         string    `db:"name,omitempty"`
	Phone        string    `db:"phone,omitempty"`
	Role         Roles     `db:"role"`
	RoleId       int       `db:"role_id,omitempty"`
	CountryId    int       `db:"country_id,omitempty"`
	Street       *string   `db:"street,omitempty"`
	Region       *string   `db:"region,omitempty"`
	Apartment    *string   `db:"apartment,omitempty"`
	RefreshToken string    `db:"refresh_token,omitempty"`
	DateOfBirth  time.Time `db:"date_of_birth,omitempty"`
	CreatedAt    time.Time `db:"created_at,omitempty"`
	UpdatedAt    time.Time `db:"updated_at,omitempty"`
}
