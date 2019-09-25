package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hero-wars/config"
	"github.com/hero-wars/game"

	"github.com/briandowns/spinner"
)

func buildHeroWarsCharacters(cfg *config.Config) (*game.Player, *game.Special, // hero
	*game.Player, *game.Special) { // villain
	buildManager := game.BuildManager{}

	hero := game.NewHero(cfg)
	buildManager.SetBuilder(hero)
	heroPlayer, heroSpecial := buildManager.Construct()

	villain := game.NewVillain(cfg)
	buildManager.SetBuilder(villain)
	villainPlayer, _villainSpecial := buildManager.Construct()

	return &heroPlayer, &heroSpecial, &villainPlayer, &_villainSpecial
}

func story() {
	log.Printf("%s\n", config.WAR_STORY)
	s := spinner.New(spinner.CharSets[39], 100*time.Millisecond)
	s.Start()
	time.Sleep(3 * time.Second)
	s.Stop()
}

func gameLoop(battle *game.Battle) {
	numRounds := 0
	for {
		story()

		battle.Duel()

		numRounds += 1
		if numRounds == config.MAX_ROUNDS {
			log.Printf("[GAME OVER] The maximum number of rounds (%d) has been reached.\n", config.MAX_ROUNDS)
			outcome, winner := battle.GetWinner()
			if outcome == config.DRAW {
				log.Printf("[%s] %s\n%s", outcome, battle.HeroPlayer, battle.VillainPlayer)
			} else {
				log.Printf("[%s] %s\n", outcome, winner)
			}
			break
		}

		if !battle.HeroPlayer.IsAlive() {
			log.Printf("[GAME OVER] Hero is dead: %s\n", battle.HeroPlayer)
			break
		}

		if !battle.VillainPlayer.IsAlive() {
			log.Printf("[GAME OVER] Villain is dead: %s\n", battle.VillainPlayer)
			break
		}
	}
}

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal("Error reading configuration file")
	}

	heroPlayer, heroSpecial, villainPlayer, _ := buildHeroWarsCharacters(cfg)

	fmt.Printf("Hero Player: %s\nHero Special: %+v\n", heroPlayer, heroSpecial)
	fmt.Printf("Villain Player: %s\n", villainPlayer)

	battle := game.NewBattle(heroPlayer, heroSpecial, villainPlayer)

	gameLoop(battle)
}
