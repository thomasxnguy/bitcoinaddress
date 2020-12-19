package p2sh

// P2shResponse is the response object for /p2sh
type P2shResponse struct {
	P2shAddress string `json:"p2sh_address"`
}

func newP2shResponse(address string) *P2shResponse {
	resp := &P2shResponse{
		P2shAddress: address,
	}
	return resp
}
