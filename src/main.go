package main

import (
	"fmt"
)

func main() {
	config := ParseConfig()

	fmt.Printf("%v\n", config)

	variants := GenerateVariants(config)

	fmt.Printf("%v\n", variants)

	// TODO generate files
}
