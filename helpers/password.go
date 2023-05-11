package helpers

import "golang.org/x/crypto/bcrypt"

func HashThePassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(passwordHash), err
}

func VerifyThePassword(hashPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}

type ResponseData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseNotData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
