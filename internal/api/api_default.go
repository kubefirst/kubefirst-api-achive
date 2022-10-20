/*
 * Kubefirst API
 *
 * Kubefirst API to serve console
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"encoding/json"
	"github.com/kubefirst/console-api/internal/telemetry"
	"github.com/kubefirst/console-api/pkg"
	"github.com/rs/zerolog/log"

	"net/http"
)

func HealthzGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)
	log.Debug().Msg("I'm health!")

	err := telemetry.SendMetric(pkg.MetricHealth)

	if err != nil {
		w.WriteHeader(http.StatusOK)
		log.Error().Err(err).Msg("An error occurred while sending telemetry metric")
		return
	}

	jsonData, err := json.Marshal(true)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		log.Error().Err(err).Msg("An error occurred while parsing response")
		return
	}

	_, err = w.Write(jsonData)

	if err != nil {
		w.WriteHeader(http.StatusOK)
		log.Error().Err(err).Msg("An error occurred while writing response")
		return
	}
}