package domains

type FilialAddressesDomain struct {
	Id        int
	Street    string
	Region    string
	Apartment string
	StreetNum string
	CityId    int
	City      *CityDomain
}

type FilialAddressesRepository interface {
	FindAll() ([]FilialAddressesDomain, error)
	FindByUserId(userId int) ([]FilialAddressesDomain, error)
	Save(FilialAddressesDomain) error
	Update(FilialAddressesDomain, int) error
	Delete(id int) error
}

type FilialAddressesDomainUsecase interface {
	FindAll() ([]FilialAddressesDomain, int, error)
	FindByUserId(userId int) ([]FilialAddressesDomain, int, error)
	Save(FilialAddressesDomain) (int, error)
	Update(FilialAddressesDomain, int) (int, error)
	Delete(id int) (int, error)
}
