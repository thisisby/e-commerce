package domains

type ContactDomain struct {
	Id    int
	Title string
	Value string
}

type ContactRepository interface {
	FindAll() ([]ContactDomain, error)
	Save(contact ContactDomain) error
	Update(contact ContactDomain) error
	FindById(id int) (ContactDomain, error)
	Delete(id int) error
}

type ContactUsecase interface {
	FindAll() ([]ContactDomain, int, error)
	Save(contact ContactDomain) (int, error)
	Update(contact ContactDomain) (int, error)
	FindById(id int) (ContactDomain, int, error)
	Delete(id int) (int, error)
}
