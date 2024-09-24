package domains

type CharacteristicsDomain struct {
	Id            int
	Name          string
	SubcategoryId int
}

type CharacteristicsRepository interface {
	FindAll() ([]CharacteristicsDomain, error)
	Save(CharacteristicsDomain) error
	Update(CharacteristicsDomain) error
	Delete(id int) error
	FindById(id int) (CharacteristicsDomain, error)
	FindAllBySubcategoryId(subcategoryId int) ([]CharacteristicsDomain, error)
}

type CharacteristicsUsecase interface {
	FindAll() ([]CharacteristicsDomain, int, error)
	Save(CharacteristicsDomain) (int, error)
	Update(CharacteristicsDomain) (int, error)
	Delete(id int) (int, error)
	FindAllBySubcategoryId(subcategoryId int) ([]CharacteristicsDomain, int, error)
	FindById(id int) (CharacteristicsDomain, int, error)
}
