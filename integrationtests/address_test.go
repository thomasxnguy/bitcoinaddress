package integrationtests

import (
	"testing"
)

func TestGenerateAddress_WithNoParamater_ShouldReturnSegwitAddresses(t *testing.T) {
	const schema = `{"native_segwit_address": "string","segwit_address": "string","user_id": "string"}`

	mockServer.
		Get("/address/gen").
		Expect(t).
		Status(200).
		Type("application/json").
		JSONSchema(schema).
		Done()
}

func TestGetAddress_NotValidId_ShouldReturn404(t *testing.T) {
	mockServer.
		Get("/address/ss").
		Expect(t).
		Status(422).
		JSON(map[string]interface{}{
			"code":   "invalid_request",
			"errors": "Must be a valid user_id",
			"status": "Unprocessable Entity",
		}).
		Done()
}

func TestGetAddress_UserNotExist_ShouldReturn404(t *testing.T) {
	mockServer.
		Get("/address/89dc3ae0-a93a-483d-87d5-6227f450ada1").
		Expect(t).
		Type("application/json").
		Status(404).
		JSON(map[string]interface{}{
			"code":   "not_found",
			"errors": "User is not found",
			"status": "Not Found",
		}).
		Done()
}

func TestGetAddress_EmptyId_ShouldReturn404(t *testing.T) {
	mockServer.
		Get("/address").
		Expect(t).
		Status(404).
		Done()
}
