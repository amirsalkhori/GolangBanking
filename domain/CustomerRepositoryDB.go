package domain

import (
	"database/sql"
	// "log"
	"_/Users/amir/Desktop/GoWithRealProject/errs"
	"_/Users/amir/Desktop/GoWithRealProject/logger"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	// var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)

	if status == ""{
		findAllSql := "SELECT * FROM customers"	
		err = d.client.Select(&customers, findAllSql)
		// rows, err = d.client.Query(findAllSql )

	}else{
		findAllSql := "SELECT * FROM customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	 	// rows, err = d.client.Query(findAllSql, status)
	}
	
	if err != nil {
		logger.Error("Error while querying customer table" + err.Error())
		return nil, errs.NewUnexpectedError("Unxepected database error")
	}

	// err = sqlx.StructScan(rows, &customers)
	// if err != nil {
	// 	logger.Error("Error while scanning customers table" + err.Error())
	// 	return nil, errs.NewUnexpectedError("Unxepected database error")
	// } 

	// for rows.Next(){
	// 	var c Customer
	// 	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	// 	if err != nil {
	// 		logger.Error("Error while scanning customers table" + err.Error())
	// 		return nil, errs.NewUnexpectedError("Unxepected database error")
	// 	} 
	// 	customers = append(customers, c)
	// }

	return customers, nil
}

func (d CustomerRepositoryDb) ById (id string) (*Customer, *errs.AppError){
	customerSql := "SELECT * FROM customers where customer_id = ?"
	// row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := d.client.Get(&c, customerSql, id)
	// err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil{
		if err == sql.ErrNoRows{
			return nil, errs.NewNotFoundError("Customer not found")
		}else{
			logger.Error("Error while scaning customer " + err.Error())
		
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &c, nil

}

func NewCustomerRepositoryDb() CustomerRepositoryDb {

	client, err := sqlx.Open("mysql", "root:changeMe@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}