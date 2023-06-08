package repository

import domain "Assignment/services/contact/internal/Domain"

type ContactRepository interface {
	CreateContact(contact *domain.Contact) error
	GetContactByID(contactID string) (*domain.Contact, error)
	UpdateContact(contact *domain.Contact) error
	DeleteContact(contactID string) error
}
