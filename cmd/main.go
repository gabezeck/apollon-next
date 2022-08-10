package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gabezeck/test-api/internal/api"
	"github.com/gabezeck/test-api/internal/config"
	"github.com/gabezeck/test-api/internal/deps"

	"github.com/gorilla/mux"
	_ "github.com/sakirsensoy/genv/dotenv/autoload"
)

func main() {
	r := mux.NewRouter()

	cfg := config.New()
	deps := deps.New(cfg)
	api.RegisterRoutes(r, deps)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	deps.Logger.Info("Server Started")

	<-done
	deps.Logger.Info("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// Handle shutting down deps, as needed
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		deps.Logger.Fatalf("Server Shutdown Failed:%+v", err)
	}
	deps.Logger.Info("Server Exited Properly")
}
