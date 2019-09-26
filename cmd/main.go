package main

import (
	"fmt"
	"log"

	"github.com/andrei3131/hero-wars/config"
	"github.com/andrei3131/hero-wars/game"
	"github.com/andrei3131/hero-wars/gamelib"
	"github.com/andrei3131/hero-wars/player"
)

func main() {
	cfg, err := config.ReadConfig(config.CONFIG_FILE)
	if err != nil {
		log.Fatal("Error reading configuration file")
	}

	heroBuilder := player.NewHero(cfg)
	villainBuilder := player.NewVillain(cfg)

	heroPlayer, heroSpecial, villainPlayer, _ := gamelib.BuildHeroWarsCharacters(heroBuilder, villainBuilder)

	fmt.Printf("Hero Player: %s\nHero Special: %s\n", heroPlayer, heroSpecial)
	fmt.Printf("Villain Player: %s\n", villainPlayer)

	strikeEnginge := &game.StrikeEngine{
		FirstTurnPlayed:          false,
		ResiliencePreviouslyUsed: false,
	}

	battle := game.NewBattle(heroPlayer, heroSpecial, villainPlayer, strikeEnginge)

	gamelib.GameLoop(battle)
}
