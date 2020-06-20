package service

import (
	"github.com/klurpicolo/finalexam/database"
	"github.com/klurpicolo/finalexam/models"
)

//Comment
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
