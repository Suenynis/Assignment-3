package delivery

import (
	domain "Assignment/services/contact/internal/Domain"
	usecase "Assignment/services/contact/internal/Usecase"
	"net/http"
)

type ContactDelivery interface {
	// Методы доставки контакта (например, HTTP обработчики)
	CreateContact(contact *domain.Contact) error
	GetContactByID(contactID string) (*domain.Contact, error)
	UpdateContact(contact *domain.Contact) error
	DeleteContact(contactID string) error
}

type ContactDeliveryImpl struct {
	// Зависимости доставки контакта, если таковые есть
	useCase usecase.ContactUseCase
}

func (i ContactDeliveryImpl) CreateContact(writer http.ResponseWriter, request *http.Request) {

}

func (i ContactDeliveryImpl) GetContactByID(writer http.ResponseWriter, request *http.Request) {

}

func NewContactDelivery(useCase usecase.ContactUseCase) *ContactDeliveryImpl {
	// Создание и инициализация экземпляра Delivery контакта
	return &ContactDeliveryImpl{
		useCase: useCase,
	}
}
