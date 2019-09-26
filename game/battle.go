package game

import (
	"log"

	"github.com/hero-wars/config"
	"github.com/hero-wars/player"
)

type Battle struct {
	HeroPlayer     *player.Player
	HeroSpecial    *player.Special
	VillainPlayer  *player.Player
	VillainSpecial *player.Special
	Engine         *Engine
}

func NewBattle(heroPlayer *player.Player, heroSpecial *player.Special,
	villainPlayer *player.Player, strikeEngine StrikeEngineInterface) *Battle {
	return &Battle{
		HeroPlayer:     heroPlayer,
		HeroSpecial:    heroSpecial,
		VillainPlayer:  villainPlayer,
		VillainSpecial: nil,
		Engine: &Engine{
			AttackerPlayer:  nil,
			AttackerSpecial: nil,
			DefenderPlayer:  nil,
			DefenderSpecial: nil,
			StrikeEngine:    strikeEngine,
		},
	}
}

func (b *Battle) InitSetHeroAttacker() {
	b.Engine.AttackerPlayer = b.HeroPlayer
	b.Engine.AttackerSpecial = b.HeroSpecial
	b.Engine.DefenderPlayer = b.VillainPlayer
	b.Engine.DefenderSpecial = b.VillainSpecial
}

func (b *Battle) InitSetVillainAttacker() {
	b.Engine.AttackerPlayer = b.VillainPlayer
	b.Engine.AttackerSpecial = b.VillainSpecial
	b.Engine.DefenderPlayer = b.HeroPlayer
	b.Engine.DefenderSpecial = b.HeroSpecial
}

func (b *Battle) Duel() {
	if !b.Engine.IsFirstTurnPlayed() {
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

		b.Engine.SetFirstTurnPlayed()
	}

	b.Engine.Attack()
	b.Engine.SwitchRoles()
	b.Engine.Attack()
	b.Engine.SwitchRoles()

}

func (b *Battle) GetWinner() (config.GameOutcome, *player.Player) {
	if b.HeroPlayer.Health > b.VillainPlayer.Health {
		return config.HERO_WINS, b.HeroPlayer
	}

	if b.HeroPlayer.Health == b.VillainPlayer.Health {
		return config.DRAW, nil
	}

	return config.VILLAIN_WINS, b.VillainPlayer
}
