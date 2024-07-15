package requests

import "ga_marketplace/internal/business/domains"

type ProfileSectionCreateRequest struct {
	Name     string  `json:"name" validate:"required"`
	Content  *string `json:"content"`
	ParentId *int    `json:"parent_id"`
}

type ProfileSectionUpdateRequest struct {
	Name     *string `json:"name"`
	Content  *string `json:"content"`
	ParentId *int    `json:"parent_id"`
}

func (p *ProfileSectionUpdateRequest) ToDomain() *domains.ProfileSectionsDomain {
	return &domains.ProfileSectionsDomain{
		Name:     *p.Name,
		Content:  p.Content,
		ParentId: p.ParentId,
	}
}

func (p *ProfileSectionCreateRequest) ToDomain() *domains.ProfileSectionsDomain {
	return &domains.ProfileSectionsDomain{
		Name:     p.Name,
		Content:  p.Content,
		ParentId: p.ParentId,
	}
}
