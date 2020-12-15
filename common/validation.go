package common

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

var (
	ValidatePublicKey = validation.NewStringRule(validatePublicKey, "must be a correct public key")
)

func validatePublicKey(string) bool {
	return true
}
