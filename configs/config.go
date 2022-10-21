package configs

import (
	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog/log"
)

type Config struct {
	KubefirstVersion string `env:"KUBEFIRST_VERSION"`
	HostedZoneName    string `env:"HOSTED_ZONE_NAME"`
}

func ReadConfig() *Config {
	config := Config{}

	if err := env.Parse(&config); err != nil {
		log.Err(err).Msg("something went wrong loading the environment variables")
		log.Panic().Send()
	}

	return &config
}
