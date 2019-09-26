package player

import (
	"math/rand"
	"time"

	"github.com/andrei3131/hero-wars/config"
	"github.com/andrei3131/hero-wars/utils"
)

type Villain struct {
	player    Player
	cfg       *config.Config
	generator *rand.Rand
}

func NewVillain(cfg *config.Config) *Villain {
	return &Villain{
		cfg:       cfg,
		generator: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (v *Villain) SetHealth() PlayerBuilder {
	intRange := utils.NewIntRange(v.cfg.Villain.Common.Health.HealthStart,
		v.cfg.Villain.Common.Health.HealthEnd,
		v.generator)
	v.player.Health = intRange.NextRandomInt()

	return v
}

func (v *Villain) SetStrength() PlayerBuilder {
	intRange := utils.NewIntRange(v.cfg.Villain.Common.Strength.StrengthStart,
		v.cfg.Villain.Common.Strength.StrengthEnd,
		v.generator)
	v.player.Strength = intRange.NextRandomInt()

	return v
}

func (v *Villain) SetDefence() PlayerBuilder {
	intRange := utils.NewIntRange(v.cfg.Villain.Common.Defence.DefenceStart,
		v.cfg.Villain.Common.Defence.DefenceEnd,
		v.generator)
	v.player.Defence = intRange.NextRandomInt()

	return v
}

func (v *Villain) SetSpeed() PlayerBuilder {
	intRange := utils.NewIntRange(v.cfg.Villain.Common.Speed.SpeedStart,
		v.cfg.Villain.Common.Speed.SpeedEnd,
		v.generator)
	v.player.Speed = intRange.NextRandomInt()

	return v
}

func (v *Villain) SetLuck() PlayerBuilder {
	intRange := utils.NewFloat64Range(v.cfg.Villain.Common.Luck.LuckStart,
		v.cfg.Villain.Common.Luck.LuckEnd,
		v.generator)
	v.player.Luck = intRange.NextRandomFloat64()

	return v
}

func (v *Villain) SetSpecialCriticalStrike() PlayerBuilder {
	// Villain has no special powers
	return v
}

func (v *Villain) SetSpecialResilience() PlayerBuilder {
	// Villain has no special powers
	return v
}

func (v *Villain) Build() (Player, Special) {
	return v.player, Special{}
}
