package cm

import (
	"gopkg.in/yaml.v3"
	"log/slog"
)

// Config is the struct that holds the configuration
type Config struct {
	Templates []string
	Matrix    map[string][]string
	Exclude   []map[string]string
	Include   []map[string]string
}

// ParseConfig parses the yaml input and returns a Config struct
func ParseConfig(inputYAML []byte) (*Config, error) {
	config := Config{}

	if err := yaml.Unmarshal(inputYAML, &config); err != nil {
		slog.Warn("Unmarshal error: %v", err)
		return nil, err
	}

	return &config, nil
}
