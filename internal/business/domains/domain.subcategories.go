package domains

type SubcategoriesDomain struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	CategoryId int    `json:"category_id"`
}

type SubcategoriesRepository interface {
	FindAll() ([]SubcategoriesDomain, error)
	FindAllByCategoryId(categoryId int) ([]SubcategoriesDomain, error)
	Save(subcategoriesDomain SubcategoriesDomain) error
	Update(subcategoriesDomain SubcategoriesDomain) error
	Delete(id int) error
	FindById(id int) (SubcategoriesDomain, error)
}

type SubcategoriesUsecase interface {
	FindAll() ([]SubcategoriesDomain, int, error)
	FindAllByCategoryId(categoryId int) ([]SubcategoriesDomain, int, error)
	Save(subcategoriesDomain SubcategoriesDomain) (int, error)
	Update(subcategoriesDomain SubcategoriesDomain) (int, error)
	Delete(id int) (int, error)
	FindById(id int) (SubcategoriesDomain, int, error)
}
