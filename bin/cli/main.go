package main

import (
	"fmt"
	"github.com/campaign-manager/internal/task_creation"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("variants.yaml")
	if err != nil {
		log.Fatal(err)
	}
	config, err := task_creation.ParseInput(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", config)

	variants := task_creation.GenerateVariants(config)

	fmt.Printf("%v\n", variants)

	// TODO generate files
}
