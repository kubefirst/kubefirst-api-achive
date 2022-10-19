package main

import (
	"github.com/gorilla/mux"
	"github.com/kubefirst/console-api/internal/handlers"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

func main() {
	// setup logging with color and code line on logs
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout}).With().Caller().Logger()

	// mux router
	r := mux.NewRouter().StrictSlash(true)
	r.Use(mux.CORSMethodMiddleware(r))

	appHandler := handlers.NewApp()

	r.HandleFunc("/healthz", appHandler.Healthz).Methods(http.MethodGet, http.MethodOptions)

	port := ":3000"
	log.Info().Msgf("API listening at %q port", port[1:])
	if err := http.ListenAndServe(port, r); err != nil {
		log.Panic().Err(err).Msg("API is down")
	}
}
