package utils

import (
	"ga_marketplace/pkg/helpers"
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	newValidator := validator.New()
	newValidator.RegisterValidation("orderstatus", helpers.OrderStatusValidator)
	newValidator.RegisterValidation("order_delivery_method", helpers.OrderDeliveryMethodValidator)
	return &Validator{
		validator: newValidator,
	}
}

func (v *Validator) Validate(in any) error {
	return v.validator.Struct(in)
}
