package repository

import (
	"database/sql"

	"github.com/dannielss/goflix/model"
)

func NewUserRepository(mysqlClient *sql.DB) UserRepositoryInterface {
	return &userRepository{mysqlClient}
}

type UserRepositoryInterface interface {
	ShowAll() (*sql.Rows, error)
	Insert(user *model.User) error
	Update(u *model.User) (int64, error)
	Delete(id int) (int64, error)
}

type userRepository struct {
	mysqlClient *sql.DB
}

func (ur *userRepository) ShowAll() (*sql.Rows, error) {
	query := "SELECT * FROM users"
	rows, err := ur.mysqlClient.Query(query)

	if err != nil {
		return rows, err
	}

	return rows, nil
}

func (ur *userRepository) Insert(u *model.User) error {
	query := "INSERT INTO users(name, email, password) VALUES (?, ?, ?)"
	_, err := ur.mysqlClient.Exec(query, u.Name, u.Email, u.Password)

	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) Update(u *model.User) (int64, error) {
	query := "UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?"

	res, err := ur.mysqlClient.Exec(query, u.Name, u.Email, u.Password, u.Id)

	if err != nil {
		return 0, err
	}

	val, error := res.RowsAffected()

	if error != nil {
		return 0, error
	}

	return val, nil
}

func (ur *userRepository) Delete(id int) (int64, error) {
	query := "DELETE FROM users WHERE id = ?"
	res, err := ur.mysqlClient.Exec(query, id)

	if err != nil {
		return 0, err
	}

	val, error := res.RowsAffected()

	if error != nil {
		return 0, error
	}

	return val, nil
}
