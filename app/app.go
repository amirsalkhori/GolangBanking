package app

import (
	"_/Users/amir/Desktop/GoWithRealProject/domain"
	"_/Users/amir/Desktop/GoWithRealProject/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// mux := http.NewServeMux()
	router := mux.NewRouter()
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)


	http.ListenAndServe("localhost:8001", router)
}

// func getCustomerItem(w http.ResponseWriter, r *http.Request){
// 	vars := mux.Vars(r)
// 	fmt.Println(vars)
// 	fmt.Fprint(w, vars["customer_id"])
// }

// func createCustomer(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprint(w, "Create customer")
// }