package domain
import (
	"_/Users/amir/Desktop/GoWithRealProject/errs"
)

type Customer struct {
	Id string `json:"id" db:"customer_id"`
	Name string `json:"name"`
	City string `json:"city"`
	Zipcode string `json:"zipcode" db:"zip_code"`
	DateOfBirth string `json:"dateOfBirth" db:"date_of_birth"`
	Status string `json:"status"`
}

//Find All Customers
type CustomeRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError) 
}

