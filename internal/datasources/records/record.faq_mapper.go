package records

import "ga_marketplace/internal/business/domains"

func (f *Faq) ToDomain() domains.FaqDomain {
	return domains.FaqDomain{
		Id:       f.Id,
		Question: f.Question,
		Answer:   f.Answer,
	}
}

func ToArrayOfFaqDomain(f []Faq) []domains.FaqDomain {
	var res []domains.FaqDomain
	for _, v := range f {
		res = append(res, v.ToDomain())
	}

	return res
}

func FromFaqDomain(d domains.FaqDomain) Faq {
	return Faq{
		Id:       d.Id,
		Question: d.Question,
		Answer:   d.Answer,
	}
}
