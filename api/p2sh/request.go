package p2sh

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/thomasxnguy/bitcoinaddress/common"
	"net/http"
)

// P2shRequest is the request object for /p2sh
type P2shRequest struct {
	Req        int      `json:"req"`
	PublicKeys []string `json:"public_keys"`
}

func (body *P2shRequest) Bind(r *http.Request) error {
	return validation.ValidateStruct(body,
		validation.Field(&body.PublicKeys, validation.Required, validation.Each(common.ValidatePublicKey)),
		validation.Field(&body.Req, validation.Required, validation.Min(0)),
	)
}
