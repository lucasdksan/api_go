package security

import "golang.org/x/crypto/bcrypt"

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(pass_string, pass_hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(pass_hash), []byte(pass_string))
}
