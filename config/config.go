package config

import (
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Config struct {
	YoutubeAPIKey string `yaml:"youtube_api_key"`
}

// unmarshalles yaml file to config.Config struct.
func Parse(pathToConfig string, dst *Config) error {
	data, err := os.ReadFile(pathToConfig)
	if err != nil {
		return errors.Wrap(err, "reading config")
	}

	if err = yaml.Unmarshal(data, dst); err != nil {
		return errors.Wrap(err, "marshalling config")
	}

	return nil
}
