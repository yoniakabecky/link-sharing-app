package password

import (
	"golang.org/x/crypto/bcrypt"
)

// cost is used by HashPassword; set in tests to trigger error path.
var cost = bcrypt.DefaultCost

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func ComparePassword(password, hashed string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}
