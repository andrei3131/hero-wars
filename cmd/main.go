package main

import (
	"log"
	"fmt"
	"github.com/hero-wars/config"
)


func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal("Error reading configuration file")
	}

	fmt.Printf("%+v\n", cfg)
}