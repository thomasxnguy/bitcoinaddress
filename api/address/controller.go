package address

import (
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"github.com/thomasxnguy/bitcoinaddress/database"
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
	accountStorer := database.NewMockAccountStore()
	service := NewService(accountStorer)

	controller := &Controller{
		service: service,
	}
	return controller, nil
}

// Router exposes address endpoints
func (rs *Controller) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/gen", rs.service.generateAddresses)
	r.Get("/{user_id}", rs.service.getUserAddresses)
	return r
}

func log(r *http.Request) logrus.FieldLogger {
	return logging.GetLogEntry(r)
}
