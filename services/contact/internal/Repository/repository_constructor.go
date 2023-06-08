package repository

import (
	domain "Assignment/services/contact/internal/Domain"
)

type ContactUseCaseImpl struct {
	repository ContactRepository
}

// Use the repository in your use case methods
func (uc *ContactUseCaseImpl) NewContactRepository(contact *domain.Contact) error {
	// Access repository methods through uc.repository
	return uc.repository.CreateContact(contact)
}
