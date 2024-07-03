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

type wishUsecase struct {
	wishRepo domains.WishRepository
}

func NewWishUsecase(wishRepo domains.WishRepository) domains.WishUsecase {
	return &wishUsecase{
		wishRepo: wishRepo,
	}
}

func (w *wishUsecase) FindByUserId(id int) (outDom []domains.WishDomain, statusCode int, err error) {

	// Find wish by user id
	wishes, err := w.wishRepo.FindByUserId(id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return wishes, http.StatusOK, nil
}

func (w *wishUsecase) Save(inDom *domains.WishDomain) (statusCode int, err error) {

	wishExists, err := w.wishRepo.FindByUserIdAndProductId(inDom.UserId, inDom.ProductId)
	if err != nil {
		slog.Error("wishUsecase.Save.FindByUserIdAndProductId", err)
		if !errors.Is(err, constants.ErrRowNotFound) {
			return http.StatusInternalServerError, err
		}
	}
	if wishExists != nil {
		return http.StatusBadRequest, fmt.Errorf("wish already exists")
	}

	inDom.CreatedAt = helpers.GetCurrentTime()
	err = w.wishRepo.Save(inDom)
	if err != nil {
		slog.Error("wishUsecase.Save", err)
		if errors.Is(err, constants.ErrForeignKeyViolation) {
			return http.StatusBadRequest, fmt.Errorf("product not found")
		}
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (w *wishUsecase) Delete(id int, userId int) (statusCode int, err error) {

	err = w.wishRepo.Delete(id, userId)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
