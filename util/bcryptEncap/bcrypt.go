package bcryptEncap

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateFromPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CompareHashAndPassword(realPassword, checkPassword []byte) error {
	return bcrypt.CompareHashAndPassword(realPassword, checkPassword)
}
