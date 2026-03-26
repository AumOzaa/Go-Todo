package models

import (
	"bytes"
	"net/http"
)

type CustomResponseWriter struct {
	http.ResponseWriter
	Buf                *bytes.Buffer
	OriginalStatusCode int
	Always200          bool
}
