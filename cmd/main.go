package main

import (
	"fmt"
	"github.com/yakiza/salty/api"
	"net/http"
)

func main() {
	r := api.NewHandler()
	fmt.Println("Running on :3000")
	http.ListenAndServe(":3000", r)
	// Graceful shutdown
}
