package records

type Contacts struct {
	Id    int    `db:"id"`
	Title string `db:"title"`
	Value string `db:"value"`
}
