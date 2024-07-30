package helpers

import (
	"bytes"
	"errors"
	"fmt"
	"ga_marketplace/internal/constants"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"strings"
)

func BindAndValidate(c echo.Context, req any) error {
	if err := c.Bind(req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {

		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			log.Error(err)
			return errors.New("something went wrong. Please try again later")
		}

		var errorsOut validationErrors
		for _, err := range err.(validator.ValidationErrors) {
			var e error
			switch err.Tag() {
			case "required":
				e = fmt.Errorf("field '%s' cannot be blank", err.Field())
			case "email":
				e = fmt.Errorf("field '%s' must be a valid email address", err.Field())
			case "eth_addr":
				e = fmt.Errorf("field '%s' must  be a valid Ethereum address", err.Field())
			case "len":
				e = fmt.Errorf("field '%s' must be exactly %v characters long", err.Field(), err.Param())
			case "orderstatus":
				e = fmt.Errorf("field '%s' must be one of [pending, shipping, delivered, cancelled]", err.Field())
			case "order_delivery_method":
				e = fmt.Errorf("field '%s' must be one of [pickup, delivery]", err.Field())
			default:
				e = fmt.Errorf("field '%s': '%v' must satisfy '%s' '%v' criteria", err.Field(), err.Value(), err.Tag(), err.Param())
			}
			errorsOut = append(errorsOut, e)
		}

		return errorsOut
	}

	return nil
}

type validationErrors []error

func (v validationErrors) Error() string {

	buff := bytes.NewBufferString("")

	for i := 0; i < len(v); i++ {

		buff.WriteString(v[i].Error())
		buff.WriteString("\n")
	}

	return strings.TrimSpace(buff.String())
}

func OrderStatusValidator(fl validator.FieldLevel) bool {
	status := fl.Field().String()
	switch status {
	case constants.Pending, constants.Shipping, constants.Delivered, constants.Cancelled:
		return true
	}
	return false
}

func OrderDeliveryMethodValidator(fl validator.FieldLevel) bool {
	deliveryMethod := fl.Field().String()
	switch deliveryMethod {
	case constants.DeliveryMethodPickup, constants.DeliveryMethodDelivery:
		return true
	}
	return false
}
