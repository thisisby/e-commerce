package domains

type CityDomain struct {
	Id                   int
	Name                 string
	DeliveryDurationDays int
}

type CitiesRepository interface {
	FindAll() ([]CityDomain, error)
	FindById(id int) (CityDomain, error)
	Save(city CityDomain) error
	Update(city CityDomain) error
	Delete(id int) error
}

type CitiesUsecase interface {
	FindAll() ([]CityDomain, int, error)
	FindById(id int) (CityDomain, int, error)
	Save(city CityDomain) (int, error)
	Update(city CityDomain) (int, error)
	Delete(id int) (int, error)
}
