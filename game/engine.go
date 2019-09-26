package game

import (
	"log"
	"math/rand"
	"time"

	"github.com/hero-wars/player"
	"github.com/hero-wars/utils"
	"github.com/hero-wars/config"
)

type Engine struct {
	FirstTurnPlayed          bool
	ResiliencePreviouslyUsed bool
	AttackerPlayer           *player.Player
	AttackerSpecial          *player.Special
	DefenderPlayer           *player.Player
	DefenderSpecial          *player.Special
}

func (engine *Engine) strike(attacker, defender string, generator *rand.Rand) int {
	log.Printf("%s attacks %s\n", attacker, defender)

	damage := utils.MaxTwoIntegers(0, engine.AttackerPlayer.Strength-engine.DefenderPlayer.Defence)

	log.Printf("Damage computed as Attacker Strength (%d) - Defender Defence (%d)\n", engine.AttackerPlayer.Strength, engine.DefenderPlayer.Defence)

	if generator.Float64() < engine.DefenderPlayer.Luck {
		log.Printf("%s is lucky, so it does not suffer any damage\n", defender)
		damage = 0
	}

	return damage
}

func (engine *Engine) subtractDefenderDamage(attacker, defender string, damage int) {
	log.Printf("Following %s's attack, the %s loses %d of his health\n", attacker, defender, damage)
	engine.DefenderPlayer.Health -= damage
	if engine.DefenderPlayer.Health < 0 {
		engine.DefenderPlayer.Health = 0
	}
}

func (engine *Engine) attack(attacker, defender string, generator *rand.Rand) {
	damage := engine.strike(attacker, defender, generator)

	if attacker == config.VILLAIN_NAME {
		// Resilience (Special Skill)
		if !engine.ResiliencePreviouslyUsed {
			if generator.Float64() < engine.DefenderSpecial.Resilience.HalfDamageResilienceProbability {
				oldDamage := damage
				damage = damage / 2
				log.Printf("%s is resilient and only takes half damage from %s: %d instead of %d\n", defender, attacker, damage, oldDamage)
				engine.ResiliencePreviouslyUsed = true
			}
		} else {
			log.Printf("%s cannot use his resiliency skill as, he has used it in a previous round\n", defender)
			engine.ResiliencePreviouslyUsed = false
		}
	}

	engine.subtractDefenderDamage(attacker, defender, damage)
	engine.printStats(attacker, defender)

	if attacker == config.HERO_NAME {
		if generator.Float64() < engine.AttackerSpecial.CriticalStrike.StrikeTwiceProbability {
			// Hero strikes twice (Special Skill)
			log.Printf("%s uses his special skill and strikes second time", attacker)
			engine.strike(attacker, defender, generator)
			engine.subtractDefenderDamage(attacker, defender, damage)
			engine.printStats(attacker, defender)

			if generator.Float64() < engine.AttackerSpecial.CriticalStrike.StrikeThirdGivenTwiceProbability {
				// Hero strikes third time given he has struck twice on this turn (Special Skill)
				log.Printf("%s uses his special skill and strikes third time", attacker)
				engine.strike(attacker, defender, generator)
				engine.subtractDefenderDamage(attacker, defender, damage)
				engine.printStats(attacker, defender)
			}
		}
	}

}

func (engine *Engine) Attack() {
	// The Villain has nil special skill,
	// so check if the attacker's special is not neil to enable it
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	if engine.AttackerSpecial == nil {
		// attacker is Villain
		engine.attack(config.VILLAIN_NAME, config.HERO_NAME, generator)
	} else {
		// attacker is Hero
		engine.attack(config.HERO_NAME, config.VILLAIN_NAME, generator)
	}
}

func (engine *Engine) SwitchRoles() {
	tempPlayer := engine.AttackerPlayer
	tempSpecial := engine.AttackerSpecial

	engine.AttackerPlayer = engine.DefenderPlayer
	engine.AttackerSpecial = engine.DefenderSpecial

	engine.DefenderPlayer = tempPlayer
	engine.DefenderSpecial = tempSpecial
}

func (engine *Engine) printStats(attacker, defender string) {
	// print hero information first
	if attacker == config.HERO_NAME {
		log.Printf("%s: %s\n", attacker, engine.AttackerPlayer)
		log.Printf("%s: %s\n", defender, engine.DefenderPlayer)
	} else {
		log.Printf("%s: %s\n", defender, engine.DefenderPlayer)
		log.Printf("%s: %s\n", attacker, engine.AttackerPlayer)
	}
}
