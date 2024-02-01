package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/yakiza/salty/adapters/fakedb"
	"github.com/yakiza/salty/api"
	"github.com/yakiza/salty/internal/customer"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Config contains all the relevant information for the http server
type Config struct {
	Address       string        `env:"SALTY_HTTP_ADDRESS" envDefault:":8080"`
	ReadTimeout   time.Duration `env:"SALTY_HTTP_READ_TIMEOUT" envDefault:"0"`
	WriteTimeout  time.Duration `env:"SALTY_HTTP_WRITE_TIMEOUT" envDefault:"0"`
	IdleTimeout   time.Duration `env:"SALTY_HTTP_IDLE_TIMEOUT" envDefault:"0"`
	PublicAddress string        `env:"SALTY_HTTP_PUBLIC_ADDRESS" envDefault:"http://localhost:8080"`
}

func main() {
	// Parse configuration coming from environment variables.
	var config Config
	if err := env.Parse(&config); err != nil {
		log.Fatal(err.Error())
	}

	// Repositories initialization "fakeDB for now"
	db := fakedb.NewCustomerRepository()

	// domainLogger main logging pipeline for the domain level functionality
	domainLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("Cannot start logic logger: ", err)
	}

	// Domain services initialization
	customerUseCases := customer.NewCustomerUseCase(db, domainLogger)

	// Passing on domain services to be wired in the API layer
	r := api.NewHandler(customerUseCases)

	// Use of system calls SIGINT and SIGTERM signals that cause a graceful stop.
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGTERM, syscall.SIGINT)

	// TODO: Add a server with the configurations from Config struct.
	if err := http.ListenAndServe(":8080", r); err != nil {
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
