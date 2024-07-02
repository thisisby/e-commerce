package usecases

import (
	"errors"
	"fmt"
	"ga_marketplace/internal/business/domains"
	"ga_marketplace/internal/constants"
	"ga_marketplace/pkg/helpers"
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

func (c *cartItemsUsecase) FindByUserId(id int) (outDom []domains.CartItemsDomain, statusCode int, err error) {
	carts, err := c.cartRepo.FindByUserId(id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(carts) == 0 {
		return nil, http.StatusNotFound, fmt.Errorf("cart_items is empty")
	}

	return carts, http.StatusOK, nil
}

func (c *cartItemsUsecase) Save(inDom *domains.CartItemsDomain) (statusCode int, err error) {

	// Check if product_id exists
	_, err = c.productRepo.FindById(inDom.ProductId)
	if err != nil {
		if errors.Is(err, constants.ErrRowNotFound) {
			return http.StatusNotFound, constants.ErrProductNotFound
		}
		return http.StatusInternalServerError, err
	}

	// Check if user_id exists
	_, err = c.userRepo.FindById(inDom.UserId)
	if err != nil {
		if errors.Is(err, constants.ErrRowNotFound) {
			return http.StatusNotFound, constants.ErrUserNotFound
		}
		return http.StatusInternalServerError, err
	}

	inDom.CreatedAt = helpers.GetCurrentTime()
	err = c.cartRepo.Save(inDom)
	if err != nil {
		if errors.Is(err, constants.ErrForeignKeyViolation) {
			return http.StatusBadRequest, err
		}
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (c *cartItemsUsecase) Delete(id int, userId int) (statusCode int, err error) {

	// Check if cart_items exists
	_, err = c.cartRepo.FindById(id)
	if err != nil {
		if errors.Is(err, constants.ErrRowNotFound) {
			return http.StatusNotFound, constants.ErrCartItemsNotFound
		}
		return http.StatusInternalServerError, err
	}

	err = c.cartRepo.Delete(id, userId)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (c *cartItemsUsecase) FindAll(userId int, isAdmin bool) (outDom []domains.CartItemsDomain, statusCode int, err error) {
	if isAdmin {
		carts, err := c.cartRepo.FindAll()
		if err != nil {
			return nil, http.StatusInternalServerError, err
		}

		if len(carts) == 0 {
			return nil, http.StatusNotFound, fmt.Errorf("cart_items is empty")
		}

		return carts, http.StatusOK, nil
	}

	carts, err := c.cartRepo.FindByUserId(userId)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(carts) == 0 {
		return nil, http.StatusNotFound, fmt.Errorf("cart_items is empty")
	}

	return carts, http.StatusOK, nil
}
