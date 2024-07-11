package records

type ProfileSections struct {
	Id       int     `db:"id"`
	Name     string  `db:"name"`
	Content  *string `db:"content"`
	ParentId *int    `db:"parent_id"`
}
