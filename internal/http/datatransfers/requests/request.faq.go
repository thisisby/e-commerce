package requests

import "ga_marketplace/internal/business/domains"

type CreateFaqRequest struct {
	Question string `json:"question" validate:"required"`
	Answer   string `json:"answer" validate:"required"`
}

func (c *CreateFaqRequest) ToDomain() domains.FaqDomain {
	return domains.FaqDomain{
		Question: c.Question,
		Answer:   c.Answer,
	}
}
