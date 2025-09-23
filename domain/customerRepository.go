package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepository {
	customers := []Customer{
		{ID: "1", Name: "Omkar", City: "Pune", Zipcode: "411057", DateofBirth: "01-01-1990", Status: "active"},
		{ID: "2", Name: "John", City: "New York", Zipcode: "10001", DateofBirth: "02-02-1985", Status: "inactive"},
		{ID: "3", Name: "Alice", City: "Los Angeles", Zipcode: "90001", DateofBirth: "03-03-1992", Status: "active"},
	}
	return CustomerRepositoryStub{customers: customers}
}
