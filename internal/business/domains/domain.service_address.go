package domains

type ServiceAddressDomain struct {
	Id      int
	CityId  int
	City    CityDomain
	Address string
}

type ServiceAddressRepository interface {
	FindAll() ([]ServiceAddressDomain, error)
	FindById(id int) (ServiceAddressDomain, error)
	Save(serviceAddress ServiceAddressDomain) error
	Update(serviceAddress ServiceAddressDomain) error
	Delete(id int) error
	FindAllByCityId(cityId int) ([]ServiceAddressDomain, error)
}

type ServiceAddressUsecase interface {
	FindAll() ([]ServiceAddressDomain, int, error)
	FindById(id int) (ServiceAddressDomain, int, error)
	Save(serviceAddress ServiceAddressDomain) (int, error)
	Update(serviceAddress ServiceAddressDomain) (int, error)
	Delete(id int) (int, error)
	FindAllByCityId(cityId int) ([]ServiceAddressDomain, int, error)
}
