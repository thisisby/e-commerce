package records

type Faq struct {
	Id       int    `db:"id"`
	Question string `db:"question"`
	Answer   string `db:"answer"`
}
