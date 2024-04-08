package authentication

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func Create_token(user_id uint64) (string, error) {
	permissions := jwt.MapClaims{}

	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = user_id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte("Secrete"))
}
