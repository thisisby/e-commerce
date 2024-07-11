package records

import "ga_marketplace/internal/business/domains"

func (r *ProfileSections) ToDomain() *domains.ProfileSectionsDomain {
	return &domains.ProfileSectionsDomain{
		Id:       r.Id,
		Name:     r.Name,
		Content:  r.Content,
		ParentId: r.ParentId,
	}
}

func FromProfileSectionDomain(domain domains.ProfileSectionsDomain) ProfileSections {
	return ProfileSections{
		Id:       domain.Id,
		Name:     domain.Name,
		Content:  domain.Content,
		ParentId: domain.ParentId,
	}
}

func ToProfileSectionDomains(records []ProfileSections) []domains.ProfileSectionsDomain {
	var sectionsDomains []domains.ProfileSectionsDomain
	sectionMap := make(map[int][]domains.ProfileSectionsDomain)

	for _, record := range records {
		domain := record.ToDomain()
		if record.ParentId != nil {
			sectionMap[*record.ParentId] = append(sectionMap[*record.ParentId], *domain)
		} else {
			sectionsDomains = append(sectionsDomains, *domain)
		}
	}

	var buildTree func(parent domains.ProfileSectionsDomain) domains.ProfileSectionsDomain
	buildTree = func(parent domains.ProfileSectionsDomain) domains.ProfileSectionsDomain {
		children, exists := sectionMap[parent.Id]
		if exists {
			for _, child := range children {
				childWithSubsections := buildTree(child)
				parent.ProfileSections = append(parent.ProfileSections, childWithSubsections)
			}
		}
		return parent
	}

	for i, section := range sectionsDomains {
		sectionsDomains[i] = buildTree(section)
	}

	return sectionsDomains
}
