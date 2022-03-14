package config

type Config struct {
	YoutubeAPIKey string `yaml:"youtube_api_key"`
}

func Init(pathToConfig string) (*Config, error) {
	return nil, nil
}
