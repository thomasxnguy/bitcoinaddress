package address

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/thomasxnguy/bitcoinaddress/common"
	"net/http"
)

type p2shRequest struct {
	N   int `json:"n"`
	M int `json:"m"`
	PublicKeys []string `json:"public_keys"`
}

func (body *p2shRequest) Bind(r *http.Request) error {
	return validation.ValidateStruct(body,
		validation.Field(&body.PublicKeys, validation.Required,  validation.Each(common.ValidatePublicKey)),
		validation.Field(&body.M, validation.Required, validation.Min(0)),
		validation.Field(&body.N, validation.Required, validation.Min(0)),
	)
}
