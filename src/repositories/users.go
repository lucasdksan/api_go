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

func (repository Users) Get_for_id(user_id uint64) (models.User, error) {
	var user models.User
	lines, err := repository.db.Query(
		"select id, name, nick, email, createAt from users where id = ?",
		user_id,
	)

	if err != nil {
		return models.User{}, err
	}

	defer lines.Close()

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreateAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
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

func (repository Users) Update(user_id uint64, user models.User) error {
	statement, err := repository.db.Prepare("update users set name = ?, nick = ?, email = ? where id = ?")

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(
		user.Name,
		user.Nick,
		user.Email,
		user_id,
	); err != nil {
		return err
	}

	return nil
}

func (repository Users) Delete(user_id uint64) error {
	statement, err := repository.db.Prepare("delete from users where id = ?")

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(user_id); err != nil {
		return err
	}

	return nil
}

func (repository Users) Search_user(email string) (models.User, error) {
	var user models.User
	line, err := repository.db.Query("select id, password from users where email = ?", email)

	if err != nil {
		return models.User{}, err
	}

	defer line.Close()

	if line.Next() {
		if err = line.Scan(
			&user.ID,
			&user.Password,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository Users) Follow(follower_id, user_id uint64) error {
	statement, err := repository.db.Prepare(
		"insert ignore into followers (user_id, follower_id) values (?, ?)",
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(user_id, follower_id); err != nil {
		return err
	}

	return nil
}

func (repository Users) Un_follow(user_id, follower_id uint64) error {
	statement, err := repository.db.Prepare(
		"delete from followers where user_id = ? and follower_id = ?",
	)

	if err != nil {
		return err
	}

	statement.Close()

	if _, err = statement.Exec(user_id, follower_id); err != nil {
		return err
	}

	return nil
}

func (repository Users) Get_followers(user_id uint64) ([]models.User, error) {
	lines, err := repository.db.Query(
		`select u.id, u.name, u.nick, u.email, u.createAt
		from users u inner join followers s on u.id = s.follower_id where s.user_id = ?
		`, user_id,
	)

	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var followers []models.User

	for lines.Next() {
		var follower models.User

		if err := lines.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreateAt,
		); err != nil {
			return nil, err
		}

		followers = append(followers, follower)
	}

	return followers, nil
}
