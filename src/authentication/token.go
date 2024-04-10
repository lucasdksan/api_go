package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func extract_token(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func Extract_user_id(r *http.Request) (uint64, error) {
	token_string := extract_token(r)
	token, err := jwt.Parse(token_string, return_key)

	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user_id, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)

		if err != nil {
			return 0, err
		}

		return user_id, nil
	}

	return 0, errors.New("token Invalido")
}

func return_key(t *jwt.Token) (interface{}, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("metodo de assinatura inesperada %v", t.Header["alg"])
	}

	return config.Secret_key, nil
}

func Create_token(user_id uint64) (string, error) {
	permissions := jwt.MapClaims{}

	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = user_id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(config.Secret_key))
}

func Validate_token(r *http.Request) error {
	token_string := extract_token(r)

	token, err := jwt.Parse(token_string, return_key)

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token invalido")
}
