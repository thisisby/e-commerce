package domains

type BrandsDomain struct {
	Id   int
	Name string
	Info string
}

type BrandsRepository interface {
	FindAll() ([]BrandsDomain, error)
	Save(BrandsDomain) error
	Update(BrandsDomain) error
	Delete(id int) error
}

type BrandsUsecase interface {
	FindAll() ([]BrandsDomain, int, error)
	Save(BrandsDomain) (int, error)
	Update(BrandsDomain) (int, error)
	Delete(id int) (int, error)
}
