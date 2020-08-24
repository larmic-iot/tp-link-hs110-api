package api

import (
	"encoding/json"
	"net/http"
)

func Handle404() http.Handler {
	type ProtocolError struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)

		_ = json.
			NewEncoder(w).
			Encode(
				ProtocolError{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				})
	})
}
