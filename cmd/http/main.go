package main

import (
	"net/http"
	"os"
	
	sw "github.com/kubefirst/console-api/internal/api"
	"github.com/kubefirst/console-api/internal/telemetry"
	"github.com/kubefirst/console-api/pkg"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// setup logging with color and code line on logs
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout}).With().Caller().Logger()

	err := telemetry.SendMetric(pkg.MetricHealth)

	if err != nil {
		log.Error().Err(err).Msg("An error occurred while sending telemetry metric")
		return
	}

	log.Printf("Kubefirst API started")
	router := sw.NewRouter()

	port := ":3000"
	log.Info().Msgf("API listening at %q port", port[1:])
	if err := http.ListenAndServe(port, router); err != nil {
		log.Panic().Err(err).Msg("API is down")
	}
}
