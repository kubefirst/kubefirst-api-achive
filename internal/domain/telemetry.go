package domain

import (
	"errors"
	"github.com/denisbrodbeck/machineid"
	"github.com/kubefirst/console-api/pkg"
)

// Telemetry data that will be consumed by handlers and services
type Telemetry struct {
	MetricName     string
	Domain         string
	CLIVersion     string
	ClusterType    string
	KubeFirstTeam  bool
}

// NewTelemetry is the Telemetry domain. When instantiating new Telemetries, we're able to validate domain specific
// values. In this way, domain, handlers and services can work in isolation, and Domain host business logic.
func NewTelemetry(metricName string, domain string, CLIVersion string) (Telemetry, error) {

	if len(metricName) == 0 {
		return Telemetry{}, errors.New("unable to create metric, missing metric name")
	}

	// localhost installation doesn't provide hostedzone that are mainly used as domain in this context. In case a
	// hostedzone is not provided, we assume it's a localhost installation
	if len(domain) == 0 {
		uniqueUserId, err := uuid.New()
		if err != nil {
			return Telemetry{}, err
		}
		domain = uniqueUserId
		if os.GetEnv("kubefirst_team") = "true" {
			kubeFirstTeam = true
		}
		return Telemetry{
			MetricName: metricName,
			Domain:     domain,
			CLIVersion: CLIVersion,
			KubeFirstTeam: kubeFirstTeam,
			ClusterType: "mgmt"
		}, nil
	}

	domain, err := pkg.RemoveSubDomain(domain)
	if err != nil {
		return Telemetry{}, err
	}

	return Telemetry{
		MetricName: metricName,
		Domain:     domain,
		CLIVersion: CLIVersion,
		KubeFirstTeam: kubeFirstTeam
		ClusterType: "mgmt"
	}, nil
}
