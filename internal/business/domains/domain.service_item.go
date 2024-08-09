package domains

type ServiceItemDomain struct {
	Id           int
	Title        string
	Duration     int
	Description  string
	Price        float64
	SubServiceId int
}

type ServiceItemRepository interface {
	FindAll() ([]ServiceItemDomain, error)
	FindById(id int) (ServiceItemDomain, error)
	FindBySubServiceId(subServiceId int) ([]ServiceItemDomain, error)
	Update(serviceItem ServiceItemDomain) error
	Delete(id int) error
	Save(serviceItem ServiceItemDomain) error
}

type ServiceItemUsecase interface {
	FindAll() ([]ServiceItemDomain, int, error)
	FindById(id int) (ServiceItemDomain, int, error)
	FindBySubServiceId(subServiceId int) ([]ServiceItemDomain, int, error)
	Update(serviceItem ServiceItemDomain) (int, error)
	Delete(id int) (int, error)
	Save(serviceItem ServiceItemDomain) (int, error)
}
