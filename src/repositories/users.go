package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB
}

func New_repository_user(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) Get(identifier string) ([]models.User, error) {
	identifier = fmt.Sprintf("%%%s%%", identifier)

	lines, err := repository.db.Query(
		"select id, name, email, nick, createAt from users where name LIKE ? or nick LIKE ?",
		identifier, identifier,
	)

	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err := lines.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Nick,
			&user.CreateAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository Users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into users (name, nick, email, password) values(?, ?, ?, ?)",
	)

	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	last_id_insert, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(last_id_insert), nil
}
