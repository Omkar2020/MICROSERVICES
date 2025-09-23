package domain

type Customer struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateofBirth string `json:"dateofbirth"`
	Status      string `json:"status"`
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
