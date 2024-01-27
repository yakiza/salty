package main

import (
	"github.com/yakiza/salty/api"
	"net/http"
)

func main() {
	r := api.NewHandler()
	http.ListenAndServe(":3000", r)
	// Graceful shutdown
}
