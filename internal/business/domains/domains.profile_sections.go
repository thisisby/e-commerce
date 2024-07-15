package domains

type ProfileSectionsDomain struct {
	Id              int
	Name            string
	Content         *string
	ParentId        *int
	ProfileSections []ProfileSectionsDomain
}

type ProfileSectionsRepository interface {
	FindAll() ([]ProfileSectionsDomain, error)
	Save(profileSection ProfileSectionsDomain) error
	UpdateById(profileSection ProfileSectionsDomain) error
	FindById(id int) (ProfileSectionsDomain, error)
	DeleteById(id int) error
}

type ProfileSectionsUsecase interface {
	FindAll() ([]ProfileSectionsDomain, int, error)
	FindById(id int) (ProfileSectionsDomain, int, error)
	Save(profileSection ProfileSectionsDomain) (int, error)
	UpdateById(profileSection ProfileSectionsDomain) (int, error)
	DeleteById(id int) (int, error)
}
