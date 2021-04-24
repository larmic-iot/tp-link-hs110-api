package http

import (
	"encoding/json"
	"net/http"
	"tp-link-hs110-api/api/model"
)

type ErrorEncoder struct {
	w http.ResponseWriter
}

func NewErrorEncoder(w http.ResponseWriter) *ErrorEncoder {
	return &ErrorEncoder{w: w}
}

func (e *ErrorEncoder) Encode(statusCode int, message string) {
	e.w.Header().Add("Content-Type", "application/json; charset=UTF-8")
	e.w.WriteHeader(statusCode)

	_ = json.
		NewEncoder(e.w).
		Encode(
			model.ProtocolError{
				Code:    statusCode,
				Message: message,
			})
}
