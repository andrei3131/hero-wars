package game

//go:generate mockgen -destination=../mocks/mock_strike_engine.go -package=mocks github.com/andrei3131/hero-wars/game StrikeEngineInterface

import (
	"log"
	"math/rand"

	"github.com/andrei3131/hero-wars/config"
	"github.com/andrei3131/hero-wars/player"
	"github.com/andrei3131/hero-wars/utils"
)

type StrikeEngineInterface interface {
	Attack(attacker, defender string, generator *rand.Rand)
	UpdateStrikeEngine(*player.Player, *player.Special,
		*player.Player, *player.Special)
	IsFirstTurnPlayed() bool
	SetFirstTurnPlayed()
}

type StrikeEngine struct {
	FirstTurnPlayed          bool
	ResiliencePreviouslyUsed bool
	AttackerPlayer           *player.Player
	AttackerSpecial          *player.Special
	DefenderPlayer           *player.Player
	DefenderSpecial          *player.Special
}

func (strikeEngine *StrikeEngine) strike(attacker, defender string, generator *rand.Rand) int {
	log.Printf("%s attacks %s\n", attacker, defender)

	damage := utils.MaxTwoIntegers(0, strikeEngine.AttackerPlayer.Strength-strikeEngine.DefenderPlayer.Defence)

	log.Printf("Damage computed as Attacker Strength (%d) - Defender Defence (%d)\n", strikeEngine.AttackerPlayer.Strength, strikeEngine.DefenderPlayer.Defence)

	if generator.Float64() < strikeEngine.DefenderPlayer.Luck {
		log.Printf("%s is lucky, so it does not suffer any damage\n", defender)
		damage = 0
	}

	return damage
}

func (strikeEngine *StrikeEngine) subtractDefenderDamage(attacker, defender string, damage int) {
	oldDefederHealth := strikeEngine.DefenderPlayer.Health
	strikeEngine.DefenderPlayer.Health -= damage
	if strikeEngine.DefenderPlayer.Health < 0 {
		log.Printf("Heavy damage: Defender's health is now negative %d, capping at 0.\n", strikeEngine.DefenderPlayer.Health)
		strikeEngine.DefenderPlayer.Health = 0
	}
	log.Printf("Following %s's attack, the %s loses %d of his health: Defender Health was %d, Defender Health becomes %d\n",
		attacker, defender, damage, oldDefederHealth, strikeEngine.DefenderPlayer.Health)
}

func (strikeEngine *StrikeEngine) UpdateStrikeEngine(attackerPlayer *player.Player,
	attackerSpecial *player.Special,
	defenderPlayer *player.Player,
	defenderSpecial *player.Special) {
	strikeEngine.AttackerPlayer = attackerPlayer
	strikeEngine.AttackerSpecial = attackerSpecial
	strikeEngine.DefenderPlayer = defenderPlayer
	strikeEngine.DefenderSpecial = defenderSpecial
}

func (strikeEngine *StrikeEngine) Attack(attacker, defender string,
	generator *rand.Rand) {
	damage := strikeEngine.strike(attacker, defender, generator)

	if attacker == config.VILLAIN_NAME {
		// Resilience (Special Skill)
		if !strikeEngine.ResiliencePreviouslyUsed {
			if generator.Float64() < strikeEngine.DefenderSpecial.Resilience.HalfDamageResilienceProbability {
				oldDamage := damage
				damage = damage / 2
				log.Printf("%s is resilient and only takes half damage from %s: %d instead of %d\n", defender, attacker, damage, oldDamage)
				strikeEngine.ResiliencePreviouslyUsed = true
			}
		} else {
			log.Printf("%s cannot use his resiliency skill as, he has used it in a previous round\n", defender)
			strikeEngine.ResiliencePreviouslyUsed = false
		}
	}

	strikeEngine.subtractDefenderDamage(attacker, defender, damage)
	strikeEngine.printStats(attacker, defender)

	if attacker == config.HERO_NAME {
		if generator.Float64() < strikeEngine.AttackerSpecial.CriticalStrike.StrikeTwiceProbability {
			// Hero strikes twice (Special Skill)
			log.Printf("%s uses his special skill and strikes second time", attacker)
			damage = strikeEngine.strike(attacker, defender, generator)
			strikeEngine.subtractDefenderDamage(attacker, defender, damage)
			strikeEngine.printStats(attacker, defender)

			if generator.Float64() < strikeEngine.AttackerSpecial.CriticalStrike.StrikeThirdGivenTwiceProbability {
				// Hero strikes third time given he has struck twice on this turn (Special Skill)
				log.Printf("%s uses his special skill and strikes third time", attacker)
				damage = strikeEngine.strike(attacker, defender, generator)
				strikeEngine.subtractDefenderDamage(attacker, defender, damage)
				strikeEngine.printStats(attacker, defender)
			}
		}
	}

}

func (strikeEngine *StrikeEngine) IsFirstTurnPlayed() bool {
	return strikeEngine.FirstTurnPlayed
}

func (strikeEngine *StrikeEngine) SetFirstTurnPlayed() {
	strikeEngine.FirstTurnPlayed = true
}

func (strikeEngine *StrikeEngine) printStats(attacker, defender string) {
	// print hero information first
	if attacker == config.HERO_NAME {
		log.Printf("%s: %s\n", attacker, strikeEngine.AttackerPlayer)
		log.Printf("%s: %s\n", defender, strikeEngine.DefenderPlayer)
	} else {
		log.Printf("%s: %s\n", defender, strikeEngine.DefenderPlayer)
		log.Printf("%s: %s\n", attacker, strikeEngine.AttackerPlayer)
	}
}
