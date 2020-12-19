package p2sh

import (
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"github.com/thomasxnguy/bitcoinaddress/logging"
	"net/http"
)

type ctxKey int

// Controller provides application resources and handlers.
type Controller struct {
	service *Service
}

// NewController configures and returns application endpoints.
func NewController() (*Controller, error) {
	service := NewService()

	controller := &Controller{
		service: service,
	}
	return controller, nil
}

// Router exposes p2sh endpoints
func (rs *Controller) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/", rs.service.generateP2SHAddress)
	return r
}

func log(r *http.Request) logrus.FieldLogger {
	return logging.GetLogEntry(r)
}
