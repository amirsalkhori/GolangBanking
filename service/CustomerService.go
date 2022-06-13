package service

import "_/Users/amir/Desktop/GoWithRealProject/domain"

type CustomerService interface {
	GetAllCustomer()([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomeRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error){
	return s.repo.FindAll()
}

func NewCustomerService(repository domain.CustomeRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}