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


	buildManager := game.BuildManager{}

	hero := game.NewHero(cfg)
	buildManager.SetBuilder(hero)
	heroPlayer, heroSpecial := buildManager.Construct()

	villain := game.NewVillain(cfg)
	buildManager.SetBuilder(villain)
	villainPlayer, _ := buildManager.Construct()

	fmt.Printf("Hero Player: %+v\nHero Special: %+v\n", heroPlayer, heroSpecial)
	fmt.Printf("Villain Player: %+v\n", villainPlayer)

}