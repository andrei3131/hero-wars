package game

import (
	"log"

	"github.com/hero-wars/config"
)

type TurnManager struct {
	firstTurnPlayed bool
	AttackerPlayer     *Player
	AttackerSpecial    *Special
	DefenderPlayer  *Player
	DefenderSpecial *Special
}


func (tm *TurnManager) Attack() {

}

func (tm *TurnManager) SwitchRoles() {
	
}