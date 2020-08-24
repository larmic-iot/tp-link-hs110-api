package api

import (
	"encoding/json"
	"net/http"
)

func Handle404() http.Handler {
	type ProtocolError struct {
		Code    int
		Message string
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.
			NewEncoder(w).
			Encode(
				ProtocolError{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				})
	})
}
