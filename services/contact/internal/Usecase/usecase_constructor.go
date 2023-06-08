package usecase

import (
	_ "Assignment/services/contact/internal/Repository"
	repository "Assignment/services/contact/internal/Repository"
	"Assignment/services/contact/internal/domain"
)

type ContactUseCaseImpl struct {
	contactRepository repository.ContactRepository
}

func NewContactUseCase(contactRepository repository.ContactRepository) *ContactUseCaseImpl {
	return &ContactUseCaseImpl{
		contactRepository: contactRepository,
	}
}

func (uc *ContactUseCaseImpl) CreateContact(contact *domain.Contact) error {
	// Implement the logic for creating a contact
	return uc.contactRepository.CreateContact(contact)
}

func (uc *ContactUseCaseImpl) GetContactByID(contactID string) (*domain.Contact, error) {
	// Implement the logic for retrieving a contact by ID
	return uc.contactRepository.GetContactByID(contactID)
}

func (uc *ContactUseCaseImpl) UpdateContact(contact *domain.Contact) error {
	// Implement the logic for updating a contact
	return uc.contactRepository.UpdateContact(contact)
}

func (uc *ContactUseCaseImpl) DeleteContact(contactID string) error {
	// Implement the logic for deleting a contact
	return uc.contactRepository.DeleteContact(contactID)
}
