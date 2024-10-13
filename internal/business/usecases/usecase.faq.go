package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"net/http"
)

type faqUsecase struct {
	faqRepository domains.FaqRepository
}

func NewFaqUsecase(faqRepo domains.FaqRepository) domains.FaqUsecase {
	return &faqUsecase{
		faqRepository: faqRepo,
	}
}

func (f *faqUsecase) FindAll() ([]domains.FaqDomain, int, error) {
	faq, err := f.faqRepository.FindAll()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(faq) == 0 {
		return nil, http.StatusNotFound, errors.New("faq not found")
	}

	return faq, http.StatusOK, nil
}

func (f *faqUsecase) Save(domain domains.FaqDomain) (int, error) {
	err := f.faqRepository.Save(domain)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (f *faqUsecase) Update(domain domains.FaqDomain, id int) (int, error) {
	err := f.faqRepository.Update(domain, id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (f *faqUsecase) Delete(id int) (int, error) {
	err := f.faqRepository.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
