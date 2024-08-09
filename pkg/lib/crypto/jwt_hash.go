package crypto

import (
	"strconv"
	"swclabs/swix/internal/config"

	"golang.org/x/crypto/bcrypt"
)

// GenPassword generate password
func GenPassword(pass string) (string, error) {
	cost, _ := strconv.Atoi(config.JwtCost)
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(hash), err
}

// ComparePassword compare password
func ComparePassword(hashPass string, pass string) error {
	errCompare := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass))
	return errCompare
}
