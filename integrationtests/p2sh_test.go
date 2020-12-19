package integrationtests

import (
	"testing"
)

func TestP2SH_WithGoodParameter_ShouldReturnOk(t *testing.T) {
	mockServer.
		Post("/p2sh").
		JSON(map[string]interface{}{
			"req":         1,
			"public_keys": [1]string{"04a882d414e478039cd5b52a92ffb13dd5e6bd4515497439dffd691a0f12af9575fa349b5694ed3155b136f09e63975a1700c9f4d4df849323dac06cf3bd6458cd"}}).
		Expect(t).
		Status(200).
		Type("application/json").
		JSON(map[string]string{"p2sh_address": "3JFyvF4YnrjBKLr3Qs6s7JkrnBEYmPygDF"}).
		Done()
}

func TestP2SH_WithNoReq_ShouldReturnValidationError(t *testing.T) {
	mockServer.
		Post("/p2sh").
		JSON(map[string]interface{}{
			"public_keys": [1]string{"04a882d414e478039cd5b52a92ffb13dd5e6bd4515497439dffd691a0f12af9575fa349b5694ed3155b136f09e63975a1700c9f4d4df849323dac06cf3bd6458cd"}}).
		Expect(t).
		Status(422).
		Type("application/json").
		JSON(map[string]interface{}{
			"code":   "validation_error",
			"errors": "p2sh validation errors",
			"status": "Unprocessable Entity",
			"validation_errors": map[string]string{
				"req": "cannot be blank",
			},
		}).
		Done()
}

func TestP2SH_WithNoPubKey_ShouldReturnValidationError(t *testing.T) {
	mockServer.
		Post("/p2sh").
		JSON(map[string]interface{}{
			"req": 1}).
		Expect(t).
		Status(422).
		Type("application/json").
		JSON(map[string]interface{}{
			"code":   "validation_error",
			"errors": "p2sh validation errors",
			"status": "Unprocessable Entity",
			"validation_errors": map[string]string{
				"public_keys": "cannot be blank",
			},
		}).
		Done()
}

func TestP2SH_ReqExceedPubKey_ShouldReturnOk(t *testing.T) {
	mockServer.
		Post("/p2sh").
		JSON(map[string]interface{}{
			"req":         3,
			"public_keys": [1]string{"04a882d414e478039cd5b52a92ffb13dd5e6bd4515497439dffd691a0f12af9575fa349b5694ed3155b136f09e63975a1700c9f4d4df849323dac06cf3bd6458cd"}}).
		Expect(t).
		Status(422).
		Type("application/json").
		JSON(map[string]interface{}{
			"code":   "invalid_request",
			"errors": "Req needs to be less or equal than number of public key provided",
			"status": "Unprocessable Entity",
		}).
		Done()
}

func TestP2SH_NotValidPubKey_ShouldReturnValidationError(t *testing.T) {
	mockServer.
		Post("/p2sh").
		JSON(map[string]interface{}{
			"req":         1,
			"public_keys": [1]string{"aa"}}).
		Expect(t).
		Status(422).
		Type("application/json").
		JSON(map[string]interface{}{
			"code":   "validation_error",
			"errors": "p2sh validation errors",
			"status": "Unprocessable Entity",
			"validation_errors": map[string]interface{}{
				"public_keys": map[string]string{
					"0": "must be a correct public key",
				},
			},
		}).
		Done()
}
