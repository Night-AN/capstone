package security

import (
	"golang.org/x/crypto/bcrypt"
)

var intensity int

func HashPassword(plain string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), intensity)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(plain, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}
