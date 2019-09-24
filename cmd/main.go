package main

import (
	"log"
	"fmt"
	
	"github.com/hero-wars/game"
	"github.com/hero-wars/config"
)


func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal("Error reading configuration file")
	}

	fmt.Printf("%+v\n", cfg)

	playerBuilder := game.PlayerBuilder{}

	hero := &game.Hero{}
	playerBuilder.SetBuilder(hero)
	playerBuilder.Construct(cfg.Hero)
	playerBuilder.PrintPlayer()

	villain := &game.Villain{}
	playerBuilder.SetBuilder(villain)
	playerBuilder.Construct(cfg.Villain)
	playerBuilder.PrintPlayer()
}