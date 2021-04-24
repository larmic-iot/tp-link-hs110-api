package api

import (
	"net/http"
	http2 "tp-link-hs110-api/api/client/http"
)

func Handle404() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http2.NewErrorEncoder(w).Encode(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	})
}
