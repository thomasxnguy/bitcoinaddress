package common

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

var (
	ValidateNetwork = validation.NewStringRule(isCorrectPath, "must be a correct path")
)

func isCorrectPath(str string) bool {
	return true
}
