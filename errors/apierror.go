package apierrors

import (
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation"
	"net/http"
)

var (
	// BadRequest return status 400 Bad Request for malformed request body.
	BadRequest = &ErrResponse{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest)}
	// NotFound returns status 404 Not Found for invalid resource request.
	NotFound = &ErrResponse{HTTPStatusCode: http.StatusNotFound, StatusText: http.StatusText(http.StatusNotFound)}
	// Unauthorized returns 401 Unauthorized.
	Unauthorized = &ErrResponse{HTTPStatusCode: http.StatusUnauthorized, StatusText: http.StatusText(http.StatusUnauthorized)}
	// Forbidden returns status 403 Forbidden for unauthorized request.
	Forbidden = &ErrResponse{HTTPStatusCode: http.StatusForbidden, StatusText: http.StatusText(http.StatusForbidden)}
	// InternalServerError returns status 500 Internal Server Error.
	InternalServerError = &ErrResponse{HTTPStatusCode: http.StatusInternalServerError, StatusText: http.StatusText(http.StatusInternalServerError)}
)

// ErrResponse renderer type for handling all sorts of errors.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime errors
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText       string            `json:"status"`                      // user-level status message
	AppCode          string            `json:"code,omitempty"`              // application-specific errors code
	ErrorText        string            `json:"errors,omitempty"`            // application-level errors message, for debugging
	ValidationErrors validation.Errors `json:"validation_errors,omitempty"` // user level model validation errors
}

// Render sets the application-specific errors code in AppCode.
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

// ErrInvalidRequest returns status 422 Unprocessable Entity including errors message.
func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnprocessableEntity,
		StatusText:     http.StatusText(http.StatusUnprocessableEntity),
		ErrorText:      err.Error(),
		AppCode:        "invalid_request",
	}
}

// ErrValidation returns status 422 Unprocessable Entity stating validation errors.
func ErrValidation(err error, valErr validation.Errors) render.Renderer {
	return &ErrResponse{
		Err:              err,
		HTTPStatusCode:   http.StatusUnprocessableEntity,
		StatusText:       http.StatusText(http.StatusUnprocessableEntity),
		ErrorText:        err.Error(),
		ValidationErrors: valErr,
		AppCode:          "validation_error",
	}
}

// ErrUnauthorized renders status 401 Unauthorized with custom errors message.
func ErrUnauthorized(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnauthorized,
		StatusText:     http.StatusText(http.StatusUnauthorized),
		ErrorText:      err.Error(),
		AppCode:        "unauthorized",
	}
}

// ErrNotFound renders status 404 NotFound with custom errors message.
func ErrNotFound(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusNotFound,
		StatusText:     http.StatusText(http.StatusNotFound),
		ErrorText:      err.Error(),
		AppCode:        "not_found",
	}
}
