package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/Omkar2020/MICROSERVICES/service"
)

// Define CustomerResponse at the TOP of the file
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
	// Get customers from service
	customers, err := ch.service.GetAllCustomers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert domain.Customer to CustomerResponse
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

	// Set response content type based on request header
	if r.Header.Get("Accept") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customerResponses)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customerResponses)
	}
}
