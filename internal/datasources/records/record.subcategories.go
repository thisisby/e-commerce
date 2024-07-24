package records

type SubcategoriesRecord struct {
	Id         int    `db:"id"`
	Name       string `db:"name"`
	CategoryId int    `db:"category_id"`
}
