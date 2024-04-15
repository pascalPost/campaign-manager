package main

import (
	"fmt"
	cm "github.com/campaign-manager/internal"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("variants.yaml")
	if err != nil {
		log.Fatal(err)
	}
	config, err := cm.ParseInput(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", config)

	variants := cm.GenerateVariants(config)

	fmt.Printf("%v\n", variants)

	// TODO generate files
}
