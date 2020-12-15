package public

import (
	"github.com/go-chi/render"
	"net/http"
)

// Service
type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (rs *Service) getHealthCheck(w http.ResponseWriter, r *http.Request) {

	render.Respond(w, r, newHealthCheckResponse())
}
