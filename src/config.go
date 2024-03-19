package main

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// Config is the struct that holds the configuration
type Config struct {
	Templates []string
	Matrix    map[string][]string
	Exclude   []map[string]string
	Include   []map[string]string
}

// ParseConfig parses the variants.yaml file and returns a Config struct
func ParseConfig() *Config {
	yamlFile, err := os.ReadFile("variants.yaml")
	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}

	config := Config{}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &config
}
