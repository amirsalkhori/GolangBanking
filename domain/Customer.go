package domain

import (
	"_/Users/amir/Desktop/GoWithRealProject/dto"
	"_/Users/amir/Desktop/GoWithRealProject/errs"
)

type Customer struct {
	Id string `db:"customer_id"`
	Name string 
	City string 
	Zipcode string `db:"zip_code"`
	DateOfBirth string `json:"dateOfBirth" db:"date_of_birth"`
	Status string
}

//Find All Customers
type CustomeRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError) 
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "deactive"
	}

	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse{
	
	return dto.CustomerResponse{
		Id: 	 	 	c.Id,
		Name: 	 	    c.Name,
		City: 	 		c.City,
		Zipcode: 		c.Zipcode,
		DateOfBirth:  	c.DateOfBirth,
		Status: 		c.statusAsText(),
	}
}
