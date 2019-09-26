package gamelib

import (
	"log"
	"time"

	"github.com/andrei3131/hero-wars/config"
	"github.com/andrei3131/hero-wars/game"
	"github.com/andrei3131/hero-wars/player"

	"github.com/briandowns/spinner"
)


func BuildHeroWarsCharacters(hero player.PlayerBuilder,
							 villain player.PlayerBuilder) (*player.Player, *player.Special, // hero
	*player.Player, *player.Special) { // villain
	buildManager := player.BuildManager{}

	buildManager.SetBuilder(hero)
	heroPlayer, heroSpecial := buildManager.Construct()

	buildManager.SetBuilder(villain)
	villainPlayer, _villainSpecial := buildManager.Construct()

	return &heroPlayer, &heroSpecial, &villainPlayer, &_villainSpecial
}

func story() {
	log.Printf("%s\n", config.WAR_STORY)
	s := spinner.New(spinner.CharSets[39], 100 * time.Millisecond)
	s.Start()
	time.Sleep(3 * time.Second)
	s.Stop()
}

func GameLoop(battle *game.Battle) {
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
