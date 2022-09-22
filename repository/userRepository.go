package repository

import (
	"database/sql"

	"github.com/dannielss/goflix/database"
	"github.com/dannielss/goflix/model"
)

func ShowAll() (*sql.Rows, error) {
	query := "SELECT * FROM users"
	rows, err := database.DBCon.Query(query)

	if err != nil {
		return rows, err
	}

	return rows, nil
}

func Insert(u model.User) error {
	query := "INSERT INTO users(name, email, password) VALUES (?, ?, ?)"
	_, err := database.DBCon.Exec(query, u.Name, u.Email, u.Password)

	if err != nil {
		return err
	}

	return nil
}

func Update(u model.User) (int64, error) {
	query := "UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?"

	res, err := database.DBCon.Exec(query, u.Name, u.Email, u.Password, u.Id)

	if err != nil {
		return 0, err
	}

	val, error := res.RowsAffected()

	if error != nil {
		return 0, error
	}

	return val, nil
}

func Delete(id int) (int64, error) {
	query := "DELETE FROM users WHERE id = ?"
	res, err := database.DBCon.Exec(query, id)

	if err != nil {
		return 0, err
	}

	val, error := res.RowsAffected()

	if error != nil {
		return 0, error
	}

	return val, nil
}
