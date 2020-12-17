package address

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

func (rs *Controller) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/gen", rs.service.generateSegWitAddress)
	r.Get("/{user_id}", rs.service.getUserSegWitAddress)
	return r
}

func log(r *http.Request) logrus.FieldLogger {
	return logging.GetLogEntry(r)
}
