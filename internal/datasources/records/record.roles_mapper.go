package records

import "ga_marketplace/internal/business/domains"

func (r *Roles) ToDomain() *domains.RoleDomain {
	return &domains.RoleDomain{
		Id:   r.Id,
		Name: r.Name,
	}
}
