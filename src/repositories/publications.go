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

func (repository Publications) Get_for_id(publication_id uint64) (models.Publications, error) {
	line, err := repository.db.Query(`
		select p.*, u.nick from 
		publications p inner join users u
		on u.id = p.author_id where p.id = ?
	`, publication_id)

	if err != nil {
		return models.Publications{}, err
	}

	defer line.Close()

	var publication models.Publications

	if line.Next() {
		if err = line.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreateAt,
			&publication.AuthorNick,
		); err != nil {
			return models.Publications{}, err
		}
	}

	return publication, nil
}

func (repository Publications) Get(user_id uint64) ([]models.Publications, error) {
	lines, err := repository.db.Query(`
		select distinct p.*, u.nick from publications p
		inner join users u on u.id = p.author_id 
		inner join followers s on p.author_id = s.user_id
		where u.id = ? or s.follower_id = ?
		order by 1 desc
	`, user_id, user_id)

	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var publications []models.Publications

	for lines.Next() {
		var publication models.Publications

		if err = lines.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreateAt,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}

func (repository Publications) Update(publication_id uint64, publication models.Publications) error {
	statement, err := repository.db.Prepare(
		"update publications set title = ?, content = ? where id = ?",
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(publication.Title, publication.Content, publication_id); err != nil {
		return err
	}

	return nil
}
