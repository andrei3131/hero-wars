package game

import (
	"math/rand"
	"time"

	"github.com/andrei3131/hero-wars/config"
	"github.com/andrei3131/hero-wars/player"
)

type Engine struct {
	AttackerPlayer  *player.Player
	AttackerSpecial *player.Special
	DefenderPlayer  *player.Player
	DefenderSpecial *player.Special
	StrikeEngine    StrikeEngineInterface
}

func (engine *Engine) Attack() {
	// The Villain has nil special skill,
	// so check if the attacker's special is not neil to enable it
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))

	engine.StrikeEngine.UpdateStrikeEngine(engine.AttackerPlayer, engine.AttackerSpecial,
		engine.DefenderPlayer, engine.DefenderSpecial)

	if engine.AttackerSpecial == nil {
		// attacker is Villain
		engine.StrikeEngine.Attack(config.VILLAIN_NAME, config.HERO_NAME, generator)
	} else {
		// attacker is Hero
		engine.StrikeEngine.Attack(config.HERO_NAME, config.VILLAIN_NAME, generator)
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

func (engine *Engine) IsFirstTurnPlayed() bool {
	return engine.StrikeEngine.IsFirstTurnPlayed()
}

func (engine *Engine) SetFirstTurnPlayed() {
	engine.StrikeEngine.SetFirstTurnPlayed()
}
