package domains

type AttributesDomain struct {
	Id                int
	Name              string
	CharacteristicsId int
}

type AttributesRepository interface {
	FindAll() ([]AttributesDomain, error)
	Save(AttributesDomain) error
	Update(AttributesDomain) error
	Delete(id int) error
	FindById(id int) (AttributesDomain, error)
	FindAllByCharacteristicsId(characteristicsId int) ([]AttributesDomain, error)
}

type AttributesUsecase interface {
	FindAll() ([]AttributesDomain, int, error)
	Save(AttributesDomain) (int, error)
	Update(AttributesDomain) (int, error)
	Delete(id int) (int, error)
	FindAllByCharacteristicsId(characteristicsId int) ([]AttributesDomain, int, error)
	FindById(id int) (AttributesDomain, int, error)
}
