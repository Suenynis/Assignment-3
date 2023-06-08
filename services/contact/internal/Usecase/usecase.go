package usecase

import domain "Assignment/services/contact/internal/Domain"

type ContactUseCase interface {
	CreateContact(contact *domain.Contact) error
	GetContactByID(contactID string) (*domain.Contact, error)
	UpdateContact(contact *domain.Contact) error
	DeleteContact(contactID string) error
}

type GroupUseCase interface {
	CreateGroup(group *domain.Group) error
	GetGroupByID(groupID string) (*domain.Group, error)
	AddContactToGroup(groupID string, contactID string) error
}
