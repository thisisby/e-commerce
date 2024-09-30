package domains

type PersonalAddressesDomain struct {
	Id        int
	UserId    int
	User      *UserDomain
	Street    string
	Region    string
	Apartment string
	StreetNum string
	CityId    int
	City      *CityDomain
}

type PersonalAddressesRepository interface {
	FindAll() ([]PersonalAddressesDomain, error)
	FindByUserId(userId int) ([]PersonalAddressesDomain, error)
	Save(PersonalAddressesDomain) error
	Update(PersonalAddressesDomain, int) error
	Delete(id int) error
}

type PersonalAddressesUsecase interface {
	FindAll() ([]PersonalAddressesDomain, int, error)
	FindByUserId(userId int) ([]PersonalAddressesDomain, int, error)
	Save(PersonalAddressesDomain) (int, error)
	Update(PersonalAddressesDomain, int) (int, error)
	Delete(id int) (int, error)
}
