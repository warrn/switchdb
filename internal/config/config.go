package config

import (
	"context"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"github.com/heetch/confita/backend/flags"
)

type DevConfig struct {
	IP         string `config:"listen-ip"`
	AdminPort  uint   `config:"admin-port"`
	PublicPort uint   `config:"public-port"`
}

type Config struct {
	DevConfig
	Production bool `config:"production"`
}

var defaultVals = Config{
	Production: false,
	DevConfig: DevConfig{
		IP:         "",
		AdminPort:  9091,
		PublicPort: 9090,
	},
}

func LoadConfig() (*Config, error) {
	loader := confita.NewLoader(
		file.NewOptionalBackend("/etc/switchdb/conf.yaml"),
		env.NewBackend(),
		flags.NewBackend(),
	)

	config := defaultVals

	err := loader.Load(context.Background(), &config)
	if err != nil {
		return nil, err
	}

	return &config, err
}
