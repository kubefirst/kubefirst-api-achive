package telemetry 

import (
	"github.com/kubefirst/console-api/internal/domain"
	"github.com/kubefirst/console-api/internal/handlers"
	"github.com/kubefirst/console-api/internal/services"
	"github.com/kubefirst/console-api/pkg"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/analytics-go"
)

func SendMetric(metricName string) error {
	// Instantiates a SegmentIO client to use send messages to the segment API.
	segmentIOClient := analytics.New(pkg.SegmentIOWriteKey)

	// SegmentIO library works with queue that is based on timing, we explicit close the http client connection
	// to force flush in case there is still some pending message in the SegmentIO library queue.
	defer func(segmentIOClient analytics.Client) {
		err := segmentIOClient.Close()
		if err != nil {
			log.Error().Err(err).Msg("An error occurred while closing the segment connection")
		}
	}(segmentIOClient)

	// validate telemetryDomain data
	telemetryDomain, err := domain.NewTelemetry(
		metricName,
		"gh.mgmt.kubefirst.com", // ToDo: Detokenize
		"1.9.9", // ToDo
	)
	if err != nil {
		log.Error().Err(err).Msg("An error occurred while creating new telemetry data")
		return err
	}

	telemetryService := services.NewSegmentIoService(segmentIOClient)
	telemetryHandler := handlers.NewTelemetryHandler(telemetryService)

	err = telemetryHandler.SendCountMetric(telemetryDomain)
	if err != nil {
		log.Error().Err(err).Msg("An error occurred while sending count metric")
		return err
	}

	return nil
}
