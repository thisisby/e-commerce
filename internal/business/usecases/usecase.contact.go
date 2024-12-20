package usecases

import (
	"errors"
	"ga_marketplace/internal/business/domains"
	"net/http"
)

type contactsUsecase struct {
	contactsRepo domains.ContactRepository
}

func NewContactsUsecase(contactsRepo domains.ContactRepository) domains.ContactUsecase {
	return &contactsUsecase{
		contactsRepo: contactsRepo,
	}
}

func (c *contactsUsecase) FindAll() ([]domains.ContactDomain, int, error) {
	contacts, err := c.contactsRepo.FindAll()
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(contacts) == 0 {
		return nil, http.StatusNotFound, errors.New("contacts not found")
	}

	return contacts, http.StatusOK, nil
}

func (c *contactsUsecase) Save(contact domains.ContactDomain) (int, error) {
	err := c.contactsRepo.Save(contact)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusCreated, nil
}

func (c *contactsUsecase) Update(contact domains.ContactDomain) (int, error) {
	err := c.contactsRepo.Update(contact)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (c *contactsUsecase) FindById(id int) (domains.ContactDomain, int, error) {
	contact, err := c.contactsRepo.FindById(id)
	if err != nil {
		return domains.ContactDomain{}, http.StatusInternalServerError, err
	}

	if contact.Id == 0 {
		return domains.ContactDomain{}, http.StatusNotFound, errors.New("contact not found")
	}

	return contact, http.StatusOK, nil
}

func (c *contactsUsecase) Delete(id int) (int, error) {
	err := c.contactsRepo.Delete(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
