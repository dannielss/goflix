package repository

import (
	"database/sql"

	"github.com/dannielss/goflix/model"
)

func NewMovieRepository(mysqlClient *sql.DB) MovieRepositoryInterface {
	return &movieRepository{mysqlClient}
}

type MovieRepositoryInterface interface {
	ShowAll() (*sql.Rows, error)
	Insert(body *model.PayloadMovie) error
}

type payloadMovie struct {
	movie      model.Movie
	categoryId int64
}

type movieRepository struct {
	mysqlClient *sql.DB
}

func (mr *movieRepository) ShowAll() (*sql.Rows, error) {
	query := "SELECT B.id, B.title, B.description, C.name as category FROM movies_categories  A INNER JOIN movies B INNER JOIN categories C ON A.movie_id  = B.id AND A.category_id = C.id;"

	rows, err := mr.mysqlClient.Query(query)

	if err != nil {
		return rows, err
	}

	return rows, nil
}

func (mr *movieRepository) Insert(body *model.PayloadMovie) error {
	out := make(chan output)

	transaction, error := mr.mysqlClient.Begin()

	if error != nil {
		return error
	}

	defer transaction.Rollback()

	go func() {
		query := "INSERT INTO movies(title, description) VALUES (?, ?)"

		stmt, err := transaction.Prepare(query)

		if err != nil {
			out <- output{Id: 0, err: err}
		}

		val, err := stmt.Exec(body.Movie.Title, body.Movie.Description)

		if err != nil {
			out <- output{Id: 0, err: err}
		}

		movieId, err := val.LastInsertId()

		if err != nil {
			out <- output{Id: 0, err: err}
		}

		out <- output{Id: movieId, err: nil}
	}()

	queryRelationship := "INSERT INTO movies_categories(category_id, movie_id) VALUES (?, ?)"

	stmt, err := transaction.Prepare(queryRelationship)

	if err != nil {
		return err
	}

	res := <-out

	if res.err != nil {
		return res.err
	}

	_, hasError := stmt.Exec(body.CategoryId, res.Id)

	if hasError != nil {
		return hasError
	}

	lastError := transaction.Commit()

	if lastError != nil {
		return lastError
	}

	return nil
}

type output struct {
	Id  int64
	err error
}
