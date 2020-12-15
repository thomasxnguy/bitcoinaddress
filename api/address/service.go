package address

import (
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation"
	apierrors "github.com/thomasxnguy/bitcoinaddress/errors"
	"net/http"
	"errors"
)

// Service to manage bitcoin addresses
type Service struct {
}

func NewService() *Service {
	return &Service{}
}

// Endpoint to generate a p2sh address from bitcoin addresses
func (rs *Service) generateP2SHAddress(w http.ResponseWriter, r *http.Request) {
	body := &p2shRequest{}
	if err := render.Bind(r, body); err != nil {
		switch err.(type) {
		case validation.Errors:
			render.Render(w, r, apierrors.ErrValidation(errors.New("p2sh validation errors"), err.(validation.Errors)))
			return
		}
		render.Render(w, r, apierrors.ErrInvalidRequest(err))
		return
	}

	// Check that m >= n in a n-of-m scheme
	if body.M > body.N {
		render.Render(
			w,
			r,
			apierrors.ErrInvalidRequest(
				errors.New("m needs to be equal or superior to n in a n-of-m scheme")))
		return
	}

	// Check that number of public key match
	if len(body.PublicKeys) != body.M {
		render.Render(
			w,
			r,
			apierrors.ErrInvalidRequest(
				errors.New("number of public key needs to match m")))
		return
	}
	render.Respond(w, r, newHealthCheckResponse())
}
