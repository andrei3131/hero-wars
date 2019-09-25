package game

import (
	"github.com/hero-wars/config"
)

type Battle struct {
	HeroPlayer *Player
	HeroSpecial *Special
	VillainPlayer *Player
}


func NewBattle(heroPlayer *Player, heroSpecial *Special, villainPlayer *Player) *Battle {
	return &Battle{
		HeroPlayer : heroPlayer,
		HeroSpecial : heroSpecial,
		VillainPlayer : villainPlayer,
	}
}

func (b *Battle) Duel() () {

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