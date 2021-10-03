package util

import passwordvalidator "github.com/wagslane/go-password-validator"

const (
	MIN_PASSWORD_LENGTH  = 8
	MIN_PASSWORD_ENTROPY = 50
)

func CheckPasswordRequierements(password string) bool {

	if password == "" || len(password) < MIN_PASSWORD_LENGTH || passwordvalidator.GetEntropy(password) < MIN_PASSWORD_ENTROPY {
		return false
	}

	return true
}
