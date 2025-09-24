package domain

import (
	"database/sql"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {
	findAllSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	rows, err := d.client.Query(findAllSQL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		var dbID int
		// Fixed: Use exact field names from Customer struct (Zipcode, DateofBirth)
		if err := rows.Scan(&dbID, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status); err != nil {
			return nil, err
		}
		c.ID = strconv.Itoa(dbID)
		customers = append(customers, c)
	}

	// Check for errors after iterating through all rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

func (d CustomerRepositoryDB) FindByID(id string) (*Customer, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, err // Invalid numeric ID
	}
	customerSQL := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	//row := d.client.QueryRow(customerSQL, id)

	var c Customer
	var dbID int
	err = d.client.QueryRow(customerSQL, idInt).Scan(
		&dbID,
		&c.Name,
		&c.City,
		&c.Zipcode,
		&c.DateofBirth,
		&c.Status,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	// Convert database int back to string for application
	c.ID = strconv.Itoa(dbID)

	return &c, nil
}

// func (d CustomerRepositoryDB) FindByID(id int) (*Customer, error) {
// 	customerSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?"
// 	row := d.client.QueryRow(customerSQL, id)

// 	var c Customer
// 	var dbID int
// 	err := row.Scan(&dbID, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
// 	if err != nil {
// 		log.Println("Error while scanning row:" + err.Error())
// 		return nil, err
// 	}
// 	c.ID = strconv.Itoa(dbID)
// 	return &c, nil
// }

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	db, err := sql.Open("mysql", "root:Mybeta#2026@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	// Fixed: Use 'db' variable (was declared but not used properly)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(time.Minute * 3)

	// Test the connection
	if err := db.Ping(); err != nil {
		panic(err)
	}

	return CustomerRepositoryDB{client: db} // Fixed: Use 'db' instead of undefined 'client'
}
