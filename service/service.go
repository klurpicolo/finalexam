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

	row := stmt.QueryRow(customer.Name, customer.Email, customer.Status)

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

//UpdateByID comment
func UpdateByID(id string, customer *models.Customer) error {
	stmt, err := database.Conn().Prepare("UPDATE customers SET name=$2, email=$3, status=$4 WHERE id=$1")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(id, customer.Name, customer.Email, customer.Status); err != nil {
		return err
	}

	return nil
}

func DeleteByID(id string) error {
	stmt, err := database.Conn().Prepare("DELETE FROM customers Where id = $1")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(id); err != nil {
		return err
	}

	return nil
}
