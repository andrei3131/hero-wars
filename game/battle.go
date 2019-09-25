package game

import (
	"log"

	"github.com/hero-wars/config"
)

type Battle struct {
	HeroPlayer     *Player
	HeroSpecial    *Special
	VillainPlayer  *Player
	VillainSpecial *Special
	TurnManager *BattleTurnManager
}

func NewBattle(heroPlayer *Player, heroSpecial *Special, villainPlayer *Player) *Battle {
	return &Battle{
		HeroPlayer:    heroPlayer,
		HeroSpecial:   heroSpecial,
		VillainPlayer: villainPlayer,
		VillainSpecial: nil,
		TurnManager : &BattleTurnManager{
			FirstTurnPlayed : false,
			ResiliencePreviouslyUsed : false,
			AttackerPlayer : nil,
			AttackerSpecial: nil,
			DefenderPlayer: nil,
			DefenderSpecial: nil,
		},
	}
}

func (b *Battle) InitSetHeroAttacker() {
	b.TurnManager.AttackerPlayer  = b.HeroPlayer
	b.TurnManager.AttackerSpecial = b.HeroSpecial
	b.TurnManager.DefenderPlayer  = b.VillainPlayer
	b.TurnManager.DefenderSpecial = b.VillainSpecial
}

func (b *Battle) InitSetVillainAttacker() {
	b.TurnManager.AttackerPlayer  = b.VillainPlayer 
	b.TurnManager.AttackerSpecial = b.VillainSpecial
	b.TurnManager.DefenderPlayer  = b.HeroPlayer
	b.TurnManager.DefenderSpecial = b.HeroSpecial
}


func (b *Battle) Duel() {
	if !b.TurnManager.FirstTurnPlayed {
	    if b.HeroPlayer.Speed > b.VillainPlayer.Speed {
			log.Printf("%s\n", "Hero has higher speed, so he attacks first")
			b.InitSetHeroAttacker()
	    }
	   
	    if b.HeroPlayer.Speed < b.VillainPlayer.Speed {
			log.Printf("%s\n", "Villain has higher speed, so he attacks first")
			b.InitSetVillainAttacker()
		}
		 
		if b.VillainPlayer.Speed == b.HeroPlayer.Speed {
			log.Printf("%s\n", "Hero and Villain have the same speed, so the luckiest of them attacks first")
			if b.HeroPlayer.Luck > b.VillainPlayer.Luck {
				log.Printf("%s\n", "Hero has higher luck, so he attacks first")
				b.InitSetHeroAttacker()
			} else {
				log.Printf("%s\n", "Villain has higher luck, so he attacks first")
				b.InitSetVillainAttacker()
			}	
		}

	   b.TurnManager.FirstTurnPlayed = true;
	}

	b.TurnManager.Attack()
	b.TurnManager.SwitchRoles()
	b.TurnManager.Attack()
	b.TurnManager.SwitchRoles()

}

func (b *Battle) GetWinner() (config.GameOutcome, *Player) {
	if b.HeroPlayer.Health > b.VillainPlayer.Health {
		return config.HERO_WINS, b.HeroPlayer
	}

	if b.HeroPlayer.Health == b.VillainPlayer.Health {
		return config.DRAW, nil
	}

	return config.VILLAIN_WINS, b.VillainPlayer
}
