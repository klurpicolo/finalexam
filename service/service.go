package service

import (
	"github.com/klurpicolo/finalexam/database"
	"github.com/klurpicolo/finalexam/models"
)

//Insert comment
func Insert(customer *models.Customer) (string, error) {
	stmt, err := database.Conn().Prepare("INSERT INTO customers (name, email, status) values ($1, $2, $3)  RETURNING id")
	if err != nil {
		return "", err
	}

	row, err2 := stmt.Query(customer.Name, customer.Email, customer.Status)
	if err2 != nil {
		return "", err
	}

	// createdCustomer := models.Customer{
	// 	ID:     "",
	// 	Name:   customer.Name,
	// 	Email:  customer.Email,
	// 	Status: customer.Status,
	// }
	var ID *string
	err3 := row.Scan(&ID)
	if err3 != nil {
		return "", err
	}

	return *ID, nil
}

//FindAll Comment
func FindAll() ([]*models.Customer, error) {
	stmt, err := database.Conn().Prepare("SELECT id, name, email, status FROM customers")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	customers := []*models.Customer{}
	for rows.Next() {
		each := &models.Customer{}
		err := rows.Scan(&each.ID, &each.Name, &each.Email, &each.Status)
		if err != nil {
			return nil, err
		}
		customers = append(customers, each)
	}

	return customers, nil
}

//FindbyID Comment
func FindbyID(id string) (*models.Customer, error) {
	stmt, err := database.Conn().Prepare("SELECT id, name, email, status FROM customers where id=$1")
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRow(id)

	customer := &models.Customer{}

	err = row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
