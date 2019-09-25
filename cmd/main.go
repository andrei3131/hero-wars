package main

import (
	"log"
	"fmt"
	"time"
	
	"github.com/hero-wars/game"
	"github.com/hero-wars/config"

	"github.com/briandowns/spinner"
)


func buildHeroWarsCharacters(cfg *config.Config) (*game.Player,*game.Special,    // hero
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


func gameLoop(heroPlayer *game.Player, heroSpecial *game.Special, villainPlayer *game.Player) {
	numRounds := 0
	for {
		story()


	
		numRounds += 1
		if numRounds == config.MAX_ROUNDS {
		   log.Printf("[GAME OVER] The maximum number of rounds (%d) has been reached.\n", config.MAX_ROUNDS)
		   outcome, winner := game.GetWinner(heroPlayer, villainPlayer)
		   if outcome == config.DRAW {
		   	   log.Printf("[%s] %s\n%s", outcome, heroPlayer, villainPlayer)
		   } else {
			   log.Printf("[%s] %s\n", outcome, winner)
		   }
		   break
		}

		if !heroPlayer.IsAlive() {
			log.Printf("[GAME OVER] Hero is dead: %s\n", heroPlayer)
			break
		}

		if !villainPlayer.IsAlive() {
			log.Printf("[GAME OVER] Villain is dead: %s\n", villainPlayer)
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

	gameLoop(heroPlayer, heroSpecial, villainPlayer)
}