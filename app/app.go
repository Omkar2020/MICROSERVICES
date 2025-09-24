package app

import (
	"log"
	"net/http"

	"github.com/Omkar2020/MICROSERVICES/domain"
	"github.com/Omkar2020/MICROSERVICES/service"
	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	// Wiring the handlers
	// ch := CustomerHandler{
	// 	service: service.NewCustomerService(
	// 		domain.NewCustomerRepositoryStub(),
	// 	),
	// ch := CustomerHandler{
	// 	service: service.NewCustomerService(
	// 		domain.NewCustomerRepositoryDB()),
	// }

	ch := CustomerHandler{
		service: service.NewCustomerService(
			domain.NewCustomerRepositoryDB()),
	}

	// Define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	//router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	//router.HandleFunc("/customers/{id}", ch.getCustomer).Methods("GET")
	// Enforce numeric IDs in the route pattern
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomer).Methods("GET")

	// Starting Server
	log.Println("Server starting on localhost:9000")
	log.Fatal(http.ListenAndServe("localhost:9000", router))
}
