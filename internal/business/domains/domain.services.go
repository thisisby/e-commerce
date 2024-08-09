package domains

type ServicesDomain struct {
	Id   int
	Name string
}

type ServicesRepository interface {
	FindAll() ([]ServicesDomain, error)
	Save(ServicesDomain) error
	Update(ServicesDomain) error
	Delete(id int) error
}

type ServicesUsecase interface {
	FindAll() ([]ServicesDomain, int, error)
	Save(ServicesDomain) (int, error)
	Update(ServicesDomain) (int, error)
	Delete(id int) (int, error)
}
