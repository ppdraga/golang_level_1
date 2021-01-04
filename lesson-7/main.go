package main

import (
	"fmt"
	"github.com/ppdraga/golang_level_1/lesson-7/confpack"
	"log"
)

func main() {
	config, err := confpack.GetConfigFromFile("conf.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(config)
}
