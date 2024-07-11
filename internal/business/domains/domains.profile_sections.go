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
}

type ProfileSectionsUsecase interface {
	FindAll() ([]ProfileSectionsDomain, int, error)
	Save(profileSection ProfileSectionsDomain) (int, error)
}
