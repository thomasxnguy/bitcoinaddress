package address

import (
	"github.com/google/uuid"
)

// GenerateAddressResponse is the response object for /address/gen
type GenerateAddressResponse struct {
	UserId              *uuid.UUID `json:"user_id"`
	SegwitAddress       string     `json:"segwit_address"`
	NativeSegwitAddress string     `json:"native_segwit_address"`
}

func newGenerateAddressResponse(userId *uuid.UUID, segwit string, nativesegwit string) *GenerateAddressResponse {
	resp := &GenerateAddressResponse{
		UserId:              userId,
		SegwitAddress:       segwit,
		NativeSegwitAddress: nativesegwit,
	}
	return resp
}

// GetAddressResponse is the response object for /address/:user_id
type GetAddressResponse struct {
	SegwitAddress       string `json:"segwit_address"`
	NativeSegwitAddress string `json:"native_segwit_address"`
}

func newGetAddressResponse(segwit string, nativesegwit string) *GetAddressResponse {
	resp := &GetAddressResponse{
		SegwitAddress:       segwit,
		NativeSegwitAddress: nativesegwit,
	}
	return resp
}
