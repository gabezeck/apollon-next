package main

import (
	"log"
	"net/http"

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

	log.Fatal(http.ListenAndServe(":8080", r))
}
