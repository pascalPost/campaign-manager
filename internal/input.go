package internal

import (
	"gopkg.in/yaml.v3"
	"log/slog"
)

// Input is the struct that holds the configuration
type Input struct {
	Templates []string
	Matrix    map[string][]string
	Exclude   []map[string]string
	Include   []map[string]string
}

// ParseInput parses the yaml input and returns a Input struct
func ParseInput(inputYAML []byte) (*Input, error) {
	config := Input{}

	if err := yaml.Unmarshal(inputYAML, &config); err != nil {
		slog.Warn("Unmarshal error: %v", err)
		return nil, err
	}

	return &config, nil
}
