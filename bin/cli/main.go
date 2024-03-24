package main

import (
	"fmt"
	"github.com/campaign-manager/src"
	"log"
	"os"
)

func main() {
	file, err := os.ReadFile("variants.yaml")
	if err != nil {
		log.Fatal(err)
	}
	config, err := cm.ParseConfig(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", config)

	variants := cm.GenerateVariants(config)

	fmt.Printf("%v\n", variants)

	// TODO generate files
}
