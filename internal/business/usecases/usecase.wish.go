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
	wishes, err := w.wishRepo.FindByUserId(id)
	if err != nil {
		slog.Error("wishUsecase.FindAllByUserId", err)
		return nil, http.StatusInternalServerError, err
	}

	if len(wishes) == 0 {
		return nil, http.StatusNotFound, constants.ErrWishNotFound
	}

	return wishes, http.StatusOK, nil
}

func (w *wishUsecase) Save(inDom *domains.WishDomain) (statusCode int, err error) {
	wishExists, err := w.wishRepo.FindByUserIdAndProductId(inDom.UserId, inDom.ProductId)
	if err != nil {
		if !errors.Is(err, constants.ErrRowNotFound) {
			slog.Error("wishUsecase.Save", err)
			return http.StatusInternalServerError, err
		}
	}

	if wishExists != nil {
		return http.StatusBadRequest, constants.ErrWishExists
	}

	inDom.CreatedAt = helpers.GetCurrentTime()
	err = w.wishRepo.Save(inDom)
	if err != nil {
		if errors.Is(err, constants.ErrForeignKeyViolation) {
			return http.StatusBadRequest, fmt.Errorf("product id or user id not found")
		}
		slog.Error("wishUsecase.Save", err)
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (w *wishUsecase) DeleteByIdAndUserId(id int, userId int) (statusCode int, err error) {
	err = w.wishRepo.DeleteByIdAndUserId(id, userId)
	if err != nil {
		slog.Error("wishUsecase.DeleteByIdAndUserId", err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
