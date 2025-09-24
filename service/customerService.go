package service

import (
	"github.com/Omkar2020/MICROSERVICES/domain"
)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomerByID(id string) (*domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomerByID(id string) (*domain.Customer, error) {
	return s.repo.FindByID(id)
}

func NewCustomerService(repository domain.CustomerRepository) *DefaultCustomerService {
	return &DefaultCustomerService{repo: repository}
}
