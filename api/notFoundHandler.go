package api

import (
	"net/http"
	http3 "tp-link-hs110-api/api/http"
)

func Handle404() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http3.NewErrorEncoder(w).Encode(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	})
}
