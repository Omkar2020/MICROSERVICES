package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"regexp"

	"github.com/Omkar2020/MICROSERVICES/service"
	"github.com/gorilla/mux"
)

type CustomerResponse struct {
	ID          string `json:"id" xml:"id"`
	Name        string `json:"name" xml:"name"`
	City        string `json:"city" xml:"city"`
	ZipCode     string `json:"zip_code" xml:"zip_code"`
	DateOfBirth string `json:"date_of_birth" xml:"date_of_birth"`
	Status      string `json:"status" xml:"status"`
}

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	customerResponses := make([]CustomerResponse, 0)
	for _, customer := range customers {
		customerResponses = append(customerResponses, CustomerResponse{
			ID:          customer.ID,
			Name:        customer.Name,
			City:        customer.City,
			ZipCode:     customer.Zipcode,
			DateOfBirth: customer.DateofBirth,
			Status:      customer.Status,
		})
	}

	if r.Header.Get("Accept") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customerResponses)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customerResponses)
	}
}

func (ch *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Validate numeric ID using regex
	matched, _ := regexp.MatchString(`^\d+$`, idStr)
	if !matched {
		http.Error(w, "Customer ID must be numeric", http.StatusBadRequest)
		return
	}

	// Pass the string ID directly to service
	customer, err := ch.service.GetCustomerByID(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if customer == nil {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
