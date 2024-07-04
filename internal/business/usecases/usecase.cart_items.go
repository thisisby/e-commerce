package usecases

import (
	"errors"
	"fmt"
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/constants"
	"ga_marketplace/pkg/helpers"
	"log/slog"
	"net/http"
)

type cartItemsUsecase struct {
	cartRepo    domains.CartItemsRepository
	userRepo    domains.UserRepository
	productRepo domains.ProductRepository
}

func NewCartsUsecase(
	cartRepo domains.CartItemsRepository,
	userRepo domains.UserRepository,
	productRepo domains.ProductRepository,
) domains.CartUsecase {
	return &cartItemsUsecase{
		cartRepo:    cartRepo,
		userRepo:    userRepo,
		productRepo: productRepo,
	}
}

func (c *cartItemsUsecase) FindAllByUserId(id int) (outDom []domains.CartItemsDomain, statusCode int, err error) {
	carts, err := c.cartRepo.FindAllByUserId(id)
	if err != nil {
		slog.Error("cartItemsUsecase.FindAllByUserId", err)
		return nil, http.StatusInternalServerError, err
	}

	if len(carts) == 0 {
		return nil, http.StatusNotFound, fmt.Errorf("cart_items is empty")
	}

	return carts, http.StatusOK, nil
}

func (c *cartItemsUsecase) Save(inDom *domains.CartItemsDomain) (statusCode int, err error) {

	// Check if cart_items exists
	cartItemExists, err := c.cartRepo.FindByUserIdAndProductId(inDom.UserId, inDom.ProductId)
	if err != nil {
		if !errors.Is(err, constants.ErrRowNotFound) {
			slog.Error("cartItemsUsecase.Save", err)
			return http.StatusInternalServerError, err
		}
	}

	if cartItemExists != nil {
		return http.StatusBadRequest, constants.ErrCartItemsExists
	}

	inDom.CreatedAt = helpers.GetCurrentTime()
	err = c.cartRepo.Save(inDom)
	if err != nil {
		if errors.Is(err, constants.ErrForeignKeyViolation) {
			return http.StatusBadRequest, fmt.Errorf("product id not found")
		}
		slog.Error("cartItemsUsecase.Save", err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (c *cartItemsUsecase) DeleteByIdAndUserId(id int, userId int) (statusCode int, err error) {

	err = c.cartRepo.DeleteByIdAndUserId(id, userId)
	if err != nil {
		slog.Error("cartItemsUsecase.DeleteByIdAndUserId", err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (c *cartItemsUsecase) UpdateByIdAndUserId(id int, userId int, cart *domains.CartItemsDomain) (statusCode int, err error) {
	cart.UpdatedAt = helpers.GetCurrentTime()
	cart.Id = id
	cart.UserId = userId
	err = c.cartRepo.UpdateByIdAndUserId(cart)
	if err != nil {
		slog.Error("cartItemsUsecase.UpdateByIdAndUserId", err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
