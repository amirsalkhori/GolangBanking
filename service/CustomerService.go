package service

import "_/Users/amir/Desktop/GoWithRealProject/domain"

type CustomerService interface {
	GetAllCustomer()([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomeRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error){
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, error){
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomeRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}