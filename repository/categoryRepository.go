package repository

import "database/sql"

func NewCategoryRepository(mysqlClient *sql.DB) CategoryRepositoryInterface {
	return &categoryRepository{mysqlClient}
}

type CategoryRepositoryInterface interface {
	ShowAll() (*sql.Rows, error)
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
