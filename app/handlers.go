package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customers struct {
	Name    string `json:"name" xml:"name"`
	Email   string `json:"email" xml:"email"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Again Go!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hello, Again Go!")

	customers := []Customers{
		{Name: "Ashish", Email: "ashish@example.com", City: "Pune", ZipCode: "411001"},
		{Name: "John", Email: "john@example.com", City: "New York", ZipCode: "10001"},
		{Name: "Jane", Email: "jane@example.com", City: "London", ZipCode: "E1 6AN"},
	}
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(customers)
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}
