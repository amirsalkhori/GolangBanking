package service

import (
	"_/Users/amir/Desktop/GoWithRealProject/domain"
	"_/Users/amir/Desktop/GoWithRealProject/dto"
	"_/Users/amir/Desktop/GoWithRealProject/errs"
)

type CustomerService interface {
	GetAllCustomer(string)([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
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

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError){
	c, err := s.repo.ById(id)
	if err != nil{
		return nil, err
	}
	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomeRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
