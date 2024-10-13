package domains

type FaqDomain struct {
	Id       int
	Question string
	Answer   string
}

type FaqRepository interface {
	FindAll() ([]FaqDomain, error)
	Save(FaqDomain) error
	Update(FaqDomain, int) error
	Delete(id int) error
}

type FaqUsecase interface {
	FindAll() ([]FaqDomain, int, error)
	Save(FaqDomain) (int, error)
	Update(FaqDomain, int) (int, error)
	Delete(id int) (int, error)
}
