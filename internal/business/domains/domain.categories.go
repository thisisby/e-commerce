package domains

type CategoriesDomain struct {
	Id   int
	Name string
}

type CategoriesRepository interface {
	FindAll() ([]CategoriesDomain, error)
	Save(CategoriesDomain) error
	Update(CategoriesDomain) error
	Delete(id int) error
}

type CategoriesUsecase interface {
	FindAll() ([]CategoriesDomain, int, error)
	Save(CategoriesDomain) (int, error)
	Update(CategoriesDomain) (int, error)
	Delete(id int) (int, error)
}
