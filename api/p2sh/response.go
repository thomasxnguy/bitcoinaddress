package p2sh

type P2shResponse struct {
	P2shAddress string `json:"p2sh_address"`
}

func newP2shResponse() *P2shResponse {
	resp := &P2shResponse{
		P2shAddress: "",
	}
	return resp
}
