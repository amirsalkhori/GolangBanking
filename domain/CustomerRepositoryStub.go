package domain

//Mock Adapter
type CustomerRepositoryStub struct {
	customers []Customer
}
//Return all customers
func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customers, nil
}

//Responsible for create new customer
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1", Name: "Amir", City:"Tehran", Zipcode: "BN11AA", DateOfBirth: "1991/01/15", Status: "1"},
		{Id: "2", Name: "Ali", City:"Tehran", Zipcode: "BN11AA", DateOfBirth: "1991/01/15", Status: "1"},
	}

	return CustomerRepositoryStub{customers}
}