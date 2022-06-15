package domain

import (
	"database/sql"
	"log"
	"time"
	"_/Users/amir/Desktop/GoWithRealProject/errs"
	
	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {

	findAllSql := "SELECT * FROM customers"
	
	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying customer table" + err.Error())
		return nil, errs.NewUnexpectedError("Unxepected database error")
	}

	customers := make([]Customer, 0)
	for rows.Next(){
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customers table" + err.Error())
			return nil, errs.NewUnexpectedError("Unxepected database error")
		} 
		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById (id string) (*Customer, *errs.AppError){
	customerSql := "SELECT * FROM customers where customer_id = ?"
	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	if err != nil{
		if err == sql.ErrNoRows{
			return nil, errs.NewNotFoundError("Customer not found")
		}else{
			log.Println("Error while scaning customer " + err.Error())
		
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil

}

func NewCustomerRepositoryDb() CustomerRepositoryDb {

	client, err := sql.Open("mysql", "root:changeMe@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}