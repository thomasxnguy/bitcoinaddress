package address

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
	apierrors "github.com/thomasxnguy/bitcoinaddress/errors"
	"net/http"
)

// Service to manage HD SegWit addresses
type Service struct {
}

func NewService() *Service {
	return &Service{}
}

// Endpoint to generate a HD SegWit address
func (rs *Service) generateSegWitAddress(w http.ResponseWriter, r *http.Request) {
	userId := uuid.New()
	render.Respond(w, r, newGenerateAddressResponse(&userId))
}

// Endpoint to get the address of a perticular user
func (rs *Service) getUserSegWitAddress(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "user_id")
	err := validation.Validate(userId,
		validation.Required, // not empty
		is.UUID, //is UUID
	)
	if err != nil {
		render.Render(w, r, apierrors.ErrInvalidRequest(err))
		return
	}
	render.Respond(w, r, newGetAddressResponse())
}
