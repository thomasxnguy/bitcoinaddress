package api

import (
	"github.com/thomasxnguy/bitcoinaddress/api/address"
	"github.com/thomasxnguy/bitcoinaddress/api/p2sh"
	"github.com/thomasxnguy/bitcoinaddress/logging"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

// New configures application resources and routes.
func New(enableCORS bool) (*chi.Mux, error) {
	logger := logging.NewLogger()

	addressAPI, err := address.NewController()
	if err != nil {
		logger.WithField("module", "address").Error(err)
		return nil, err
	}

	p2shAPI, err := p2sh.NewController()
	if err != nil {
		logger.WithField("module", "address").Error(err)
		return nil, err
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(SetRequestID())
	r.Use(middleware.Compress(0))
	r.Use(middleware.Timeout(15 * time.Second))

	r.Use(logging.NewStructuredLogger(logger))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// use CORS middleware if client is not served by this api, e.g. from other domain or CDN
	if enableCORS {
		r.Use(corsConfig().Handler)
	}

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"Status":"Ok"}`))
	})
	r.Mount("/address", addressAPI.Router())
	r.Mount("/p2sh", p2shAPI.Router())

	return r, nil
}

func corsConfig() *cors.Cors {
	// Basic CORS
	return cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           86400, // Maximum value not ignored by any of major browsers
	})
}

// SetRequestID set the X-Request-id in the header.
func SetRequestID() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			w.Header().Set(middleware.RequestIDHeader, middleware.GetReqID(ctx))
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
