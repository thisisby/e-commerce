package records

type Brands struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Info string `db:"info"`
}
