package domain

type Customer struct {
	Id string `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
	ZipCode string `json:"zip"`
	DateOfBirth string `json:"dateOfBirth"`
	Status string `json:"status"`
}

//Find All Customers
type CustomeRepository interface {
	FindAll() ([]Customer, error)
}

