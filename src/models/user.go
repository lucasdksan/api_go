package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID       uint64    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	CreateAt time.Time `json:"CreateAt,omitempty"`
}

func (u *User) User_init() error {
	if err := u.validate(); err != nil {
		return err
	}

	u.format()

	return nil
}

func (u *User) validate() error {
	if u.Name == "" {
		return errors.New("o nome eh obrigatorio e nao pode estar em branco")
	}

	if u.Email == "" {
		return errors.New("o e-mail eh obrigatorio e nao pode estar em branco")
	}

	if u.Nick == "" {
		return errors.New("o nick eh obrigatorio e nao pode estar em branco")
	}

	if u.Password == "" {
		return errors.New("a senha eh obrigatorio e nao pode estar em branco")
	}

	return nil
}

func (u *User) format() {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
}
