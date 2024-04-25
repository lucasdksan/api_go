package repositories

import (
	"api/src/models"
	"database/sql"
)

type Publications struct {
	db *sql.DB
}

func New_repository_publication(db *sql.DB) *Publications {
	return &Publications{}
}

func (repository Publications) Create(publication models.Publications) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into publications (title, content, author_id) values (?, ?, ?)",
	)

	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(publication.Title, publication.Content, publication.AuthorID)

	if err != nil {
		return 0, err
	}

	last_id_insert, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(last_id_insert), nil
}
