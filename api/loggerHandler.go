package api

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("%s\t%s\t%s\t(in progress)", r.Method, r.RequestURI, name)

		inner.ServeHTTP(w, r)

		log.Printf("%s\t%s\t%s\ttakes %s", r.Method, r.RequestURI, name, time.Since(start))
	})
}
