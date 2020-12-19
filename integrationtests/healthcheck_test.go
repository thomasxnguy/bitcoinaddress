package integrationtests

import (
	"testing"
)

func TestHealthCheck_WithNoParamater_ShouldReturnOk(t *testing.T) {
	mockServer.
		Get("/health").
		Expect(t).
		Status(200).
		Type("application/json").
		JSON(map[string]string{"Status": "Ok"}).
		Done()
}
