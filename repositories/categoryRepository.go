package repositories

import (
	"database/sql"

	"github.com/dannielss/goflix/model"
)

func NewCategoryRepository(mysqlClient *sql.DB) CategoryRepositoryInterface {
	return &categoryRepository{mysqlClient}
}

type CategoryRepositoryInterface interface {
	ShowAll() (*sql.Rows, error)
	AddCategory(*model.Category) error
}

type categoryRepository struct {
	mysqlClient *sql.DB
}

func (cr *categoryRepository) ShowAll() (*sql.Rows, error) {
	query := "SELECT * FROM categories"
	rows, err := cr.mysqlClient.Query(query)

	if err != nil {
		return rows, err
	}

	return rows, nil
}

func (cr *categoryRepository) AddCategory(category *model.Category) error {
	query := "INSERT INTO categories (name) VALUES (?)"

	stmt, err := cr.mysqlClient.Prepare(query)

	if err != nil {
		return err
	}

	_, execError := stmt.Exec(&category.Name)

	if execError != nil {
		return execError
	}

	return nil
}
