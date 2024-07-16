package domains

type CountryDomain struct {
	Id   int
	Name string
}

type CountriesRepository interface {
	FindAll() ([]CountryDomain, error)
	FindById(id int) (CountryDomain, error)
	Save(country CountryDomain) error
	Update(country CountryDomain) error
	Delete(id int) error
}

type CountriesUsecase interface {
	FindAll() ([]CountryDomain, int, error)
	FindById(id int) (CountryDomain, int, error)
	Save(country CountryDomain) (int, error)
	Update(country CountryDomain) (int, error)
	Delete(id int) (int, error)
}
