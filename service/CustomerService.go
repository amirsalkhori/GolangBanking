package service

import (
	"_/Users/amir/Desktop/GoWithRealProject/domain"
	"_/Users/amir/Desktop/GoWithRealProject/errs"
)

type CustomerService interface {
	GetAllCustomer(string)([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomeRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError){
	if status == "active"{
		status = "1"
	}else if status == "deactive"{
		status = "0"
	}else{
		status = ""
	}
	return s.repo.FindAll(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError){
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomeRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}