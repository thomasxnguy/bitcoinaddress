package address

import (
	"github.com/google/uuid"
)

type GenerateAddressResponse struct {
	UserId *uuid.UUID `json:"user_id"`
	Address string `json:"address"`
}

func newGenerateAddressResponse(userId *uuid.UUID) *GenerateAddressResponse {
	resp := &GenerateAddressResponse{
		UserId: userId,
		Address: "",
	}
	return resp
}

type GetAddressResponse struct {
	Address string `json:"address"`
}

func newGetAddressResponse() *GetAddressResponse {
	resp := &GetAddressResponse{
		Address: "",
	}
	return resp
}
