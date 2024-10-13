package responses

import "ga_marketplace/internal/business/domains"

type FaqResponse struct {
	Id       int    `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

func FromFaqDomain(domain domains.FaqDomain) FaqResponse {
	return FaqResponse{
		Id:       domain.Id,
		Question: domain.Question,
		Answer:   domain.Answer,
	}
}

func ToArrayOfFaqResponse(faqs []domains.FaqDomain) []FaqResponse {
	var result []FaqResponse
	for _, v := range faqs {
		result = append(result, FromFaqDomain(v))
	}
	return result
}
