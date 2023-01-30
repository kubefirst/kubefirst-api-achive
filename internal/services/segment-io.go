package services

import (
	"github.com/segmentio/analytics-go"
)

// SegmentIoService hosts SegmentIO requirements
type SegmentIoService struct {
	SegmentIOClient analytics.Client
}

// NewSegmentIoService instantiate a new SegmentIO service.
func NewSegmentIoService(segmentIoClient analytics.Client) SegmentIoService {
	return SegmentIoService{
		SegmentIOClient: segmentIoClient,
	}
}

// EnqueueCountMetric use the service SegmentIO client that also has a http client to communicate with SegmentIO API.
func (service SegmentIoService) EnqueueCountMetric(metricName string, domain string, cliVersion string, kubeFirstTeam) error {

	// Enqueues a track event that will be sent asynchronously.
	err := service.SegmentIOClient.Enqueue(analytics.Track{
		UserId: domain,
		Event:  metricName,
		Properties: analytics.NewProperties().
			Set("domain", domain).
			Set("cli_version", cliVersion),
			Set("kubefirst_team", KubeFirstTeam)
	})
	if err != nil {
		return err
	}

	return nil
}
