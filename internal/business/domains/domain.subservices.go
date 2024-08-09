package domains

type SubServicesDomain struct {
	Id        int
	Name      string
	ServiceId int
}

type SubServicesRepository interface {
	FindAll() ([]SubServicesDomain, error)
	Save(SubServicesDomain) error
	Update(SubServicesDomain) error
	Delete(id int) error
	FindAllByServiceId(serviceId int) ([]SubServicesDomain, error)
	FindById(id int) (SubServicesDomain, error)
}

type SubServicesUsecase interface {
	FindAll() ([]SubServicesDomain, int, error)
	Save(SubServicesDomain) (int, error)
	Update(SubServicesDomain) (int, error)
	Delete(id int) (int, error)
	FindAllByServiceId(serviceId int) ([]SubServicesDomain, int, error)
	FindById(id int) (SubServicesDomain, int, error)
}
