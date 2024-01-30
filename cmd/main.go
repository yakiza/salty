package main

import (
	"fmt"
	"github.com/yakiza/salty/api"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	r := api.NewHandler()
	fmt.Println("Running on :3000")
	// Use of system calls SIGINT and SIGTERM signals that cause a gracefully stop.
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGTERM, syscall.SIGINT)

	if err := http.ListenAndServe(":8081", r); err != nil {
		return
	}

	// Graceful shutdown
	switch s := <-signalCh; s {
	case syscall.SIGTERM:
		log.Println("Terminating gracefully.")
		//if err := srv.Shutdown(context.Background()); err != http.ErrServerClosed {
		//	log.Println("Failed to shutdown server:", err)
		//}
	case syscall.SIGINT:
		log.Println("Terminating eagerly.")
		os.Exit(-int(syscall.SIGINT))
	}
}
