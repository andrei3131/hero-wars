package game

import (
	"math/rand"
	"time"

	"github.com/hero-wars/config"
	"github.com/hero-wars/utils"
)

type Hero struct {
	player    Player
	special   Special
	cfg       *config.Config
	generator *rand.Rand
}

func NewHero(cfg *config.Config) *Hero {
	return &Hero{
		cfg:       cfg,
		generator: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (h *Hero) SetHealth() PlayerBuilder {
	intRange := utils.NewIntRange(h.cfg.Hero.Common.Health.HealthStart,
		h.cfg.Hero.Common.Health.HealthEnd,
		h.generator)
	h.player.Health = intRange.NextRandomInt()

	return h
}

func (h *Hero) SetStrength() PlayerBuilder {
	intRange := utils.NewIntRange(h.cfg.Hero.Common.Strength.StrengthStart,
		h.cfg.Hero.Common.Strength.StrengthEnd,
		h.generator)
	h.player.Strength = intRange.NextRandomInt()

	return h
}

func (h *Hero) SetDefence() PlayerBuilder {
	intRange := utils.NewIntRange(h.cfg.Hero.Common.Defence.DefenceStart,
		h.cfg.Hero.Common.Defence.DefenceEnd,
		h.generator)
	h.player.Defence = intRange.NextRandomInt()

	return h
}

func (h *Hero) SetSpeed() PlayerBuilder {
	intRange := utils.NewIntRange(h.cfg.Hero.Common.Speed.SpeedStart,
		h.cfg.Hero.Common.Speed.SpeedEnd,
		h.generator)
	h.player.Speed = intRange.NextRandomInt()

	return h
}

func (h *Hero) SetLuck() PlayerBuilder {
	intRange := utils.NewFloat32Range(h.cfg.Hero.Common.Luck.LuckStart,
		h.cfg.Hero.Common.Luck.LuckEnd,
		h.generator)
	h.player.Luck = intRange.NextRandomFloat32()

	return h
}

func (h *Hero) SetSpecialCriticalStrike() PlayerBuilder {
	h.special.CriticalStrike.StrikeTwiceProbability = h.cfg.Hero.Special.CriticalStrike.StrikeTwiceProbability
	h.special.CriticalStrike.StrikeThirdGivenTwiceProbability = h.cfg.Hero.Special.CriticalStrike.StrikeThirdGivenTwiceProbability

	return h
}

func (h *Hero) SetSpecialResilience() PlayerBuilder {
	h.special.Resilience.HalfDamageResilienceProbability = h.cfg.Hero.Special.Resilience.HalfDamageResilienceProbability

	return h
}

func (h *Hero) Build() (Player, Special) {
	return h.player, h.special
}
