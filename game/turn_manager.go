package game

import (
	"time"
	"log"
	"math/rand"

	"github.com/hero-wars/utils"
)

type BattleTurnManager struct {
	FirstTurnPlayed bool
	ResiliencePreviouslyUsed bool
	AttackerPlayer     *Player
	AttackerSpecial    *Special
	DefenderPlayer  *Player
	DefenderSpecial *Special
}

func (btm *BattleTurnManager) printStats(attacker, defender string) {
	// print hero information first
	if attacker == "Hero" {
	   log.Printf("%s: %s\n", attacker, btm.AttackerPlayer)	
	   log.Printf("%s: %s\n", defender, btm.DefenderPlayer)	
	} else {
	   log.Printf("%s: %s\n", defender, btm.DefenderPlayer)	
	   log.Printf("%s: %s\n", attacker, btm.AttackerPlayer)	
	}
}

func (btm *BattleTurnManager) strike(attacker, defender string, generator *rand.Rand) int {
	log.Printf("%s attacks %s\n", attacker, defender)

	damage := utils.MaxTwoIntegers(0, btm.AttackerPlayer.Strength - btm.DefenderPlayer.Defence)

	log.Printf("Damage computed as Attacker Strength (%d) - Defender Defence (%d)\n", btm.AttackerPlayer.Strength, btm.DefenderPlayer.Defence)

	if generator.Float32() < btm.DefenderPlayer.Luck {
		log.Printf("%s is lucky, so it does not suffer any damage\n", defender)
		damage = 0
	}

	return damage
}

func (btm *BattleTurnManager) subtractDefenderDamage(attacker, defender string, damage int) {
	log.Printf("Following %s's attack, the %s loses %d of his health\n", attacker, defender, damage)
	btm.DefenderPlayer.Health -= damage
	if btm.DefenderPlayer.Health < 0 {
	   btm.DefenderPlayer.Health = 0
	}
}

func (btm *BattleTurnManager) attack(attacker, defender string, generator *rand.Rand) {
	damage := btm.strike(attacker, defender, generator)


	if attacker == "Villain" {
		// Resilience (Special Skill)
		if !btm.ResiliencePreviouslyUsed {
			if generator.Float32() < btm.DefenderSpecial.Resilience.HalfDamageResilienceProbability {
				oldDamage := damage
				damage = damage / 2
				log.Printf("%s is resilient and only takes half damage from %s: %d instead of %d\n", defender, attacker, damage, oldDamage)
				btm.ResiliencePreviouslyUsed = true
			}
		} else {
			log.Printf("%s cannot use his resiliency skill as, he has used it in a previous round\n", defender)
			btm.ResiliencePreviouslyUsed = false
		}
	}


	btm.subtractDefenderDamage(attacker, defender, damage)
	btm.printStats(attacker, defender)

	if attacker == "Hero" {
		if generator.Float32() < btm.AttackerSpecial.CriticalStrike.StrikeTwiceProbability {
			// Hero strikes twice (Special Skill)
			log.Printf("%s uses his special skill and strikes second time", attacker)
			btm.strike(attacker, defender, generator)
			btm.subtractDefenderDamage(attacker, defender, damage)
			btm.printStats(attacker, defender)

		    if generator.Float32() < btm.AttackerSpecial.CriticalStrike.StrikeThirdGivenTwiceProbability {
			    // Hero strikes third time given he has struck twice on this turn (Special Skill)
				log.Printf("%s uses his special skill and strikes third time", attacker)
				btm.strike(attacker, defender, generator)
				btm.subtractDefenderDamage(attacker, defender, damage)
				btm.printStats(attacker, defender)
		    }
		}		
	}

}


func (btm *BattleTurnManager) Attack() {
	// The Villain has nil special skill, 
	// so check if the attacker's special is not neil to enable it
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	if btm.AttackerSpecial == nil {
		// attacker is Villain
		btm.attack("Villain", "Hero", generator)
	} else {
		// attacker is Hero
		btm.attack("Hero", "Villain", generator)
	}
}

func (btm *BattleTurnManager) SwitchRoles() {
	tempPlayer  := btm.AttackerPlayer
	tempSpecial := btm.AttackerSpecial

	btm.AttackerPlayer  = btm.DefenderPlayer
	btm.AttackerSpecial = btm.DefenderSpecial

	btm.DefenderPlayer  = tempPlayer
	btm.DefenderSpecial = tempSpecial
}