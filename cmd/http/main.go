package main

import (
	sw "github.com/kubefirst/console-api/internal/api"
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

	log.Printf("Console API started")
	router := sw.NewRouter()

	port := ":3000"
	log.Info().Msgf("API listening at %q port", port[1:])
	if err := http.ListenAndServe(port, router); err != nil {
		log.Panic().Err(err).Msg("API is down")
	}
}
