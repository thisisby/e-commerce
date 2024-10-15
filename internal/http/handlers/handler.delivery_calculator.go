package handlers

import (
	"ga_marketplace/internal/http/datatransfers/requests"
	"ga_marketplace/pkg/helpers"
	"ga_marketplace/third_party/cdek"
	"github.com/labstack/echo/v4"
	"net/http"
)

type DeliveryCalculatorHandler struct {
	cdekClient *cdek.Client
}

func NewDeliveryCalculatorHandler(cdekClient *cdek.Client) DeliveryCalculatorHandler {
	return DeliveryCalculatorHandler{
		cdekClient: cdekClient,
	}
}

func (h *DeliveryCalculatorHandler) Calculate(ctx echo.Context) error {
	var deliveryCalculatorRequest requests.DeliveryCalculatorRequest

	if err := helpers.BindAndValidate(ctx, &deliveryCalculatorRequest); err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	fromCityCode, err := h.cdekClient.GetCityCode(deliveryCalculatorRequest.FromLocation.City)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
	toCityCode, err := h.cdekClient.GetCityCode(deliveryCalculatorRequest.ToLocation.City)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	out, err := h.cdekClient.CalculateDelivery(fromCityCode, toCityCode, deliveryCalculatorRequest)
	if err != nil {
		return NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}

	return NewSuccessResponse(ctx, http.StatusOK, "Calculated!", out)

}
