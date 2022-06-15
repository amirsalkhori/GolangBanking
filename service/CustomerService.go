package service

import (
	"_/Users/amir/Desktop/GoWithRealProject/domain"
	"_/Users/amir/Desktop/GoWithRealProject/errs"
)

type CustomerService interface {
	GetAllCustomer()([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomeRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, *errs.AppError){
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError){
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomeRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}