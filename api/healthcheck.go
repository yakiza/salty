package api

import (
	"net/http"
)

func healthCheckController() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Running"))
		if err != nil {
			return
		}
	})

}
