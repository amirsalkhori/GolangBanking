package app

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
	"_/Users/amir/Desktop/GoWithRealProject/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request){
	log.Println("Get all customers")

	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml"{
		w.Header().Add("Content-Type", "application/xml")
		w.WriteHeader(http.StatusOK)
		xml.NewEncoder(w).Encode(customers)
	}else{
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers)
	}

}