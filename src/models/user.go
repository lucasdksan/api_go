package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	CreateAt time.Time `json:"createAt,omitempty"`
}

func (u *User) User_init(state string) error {
	if err := u.validate(state); err != nil {
		return err
	}

	if err := u.format(state); err != nil {
		return err
	}

	return nil
}

func (u *User) validate(state string) error {
	if u.Name == "" {
		return errors.New("o nome eh obrigatorio e nao pode estar em branco")
	}

	if u.Email == "" {
		return errors.New("o e-mail eh obrigatorio e nao pode estar em branco")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("o e-mail possui um formato invalido")
	}

	if u.Nick == "" {
		return errors.New("o nick eh obrigatorio e nao pode estar em branco")
	}

	if state == "register" && u.Password == "" {
		return errors.New("a senha eh obrigatorio e nao pode estar em branco")
	}

	return nil
}

func (u *User) format(state string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)

	if state == "register" {
		pass_hash, err := security.Hash(u.Password)

		if err != nil {
			return err
		}

		u.Password = string(pass_hash)
	}

	return nil
}
