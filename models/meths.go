package models

import (
	"net/http"

	"github.com/go-chi/render"
)

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest() render.Renderer {
	return &ErrResponse{
		// Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      "Invalid request.",
	}
}
