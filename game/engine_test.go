package game_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/andrei3131/hero-wars/game"
	"github.com/andrei3131/hero-wars/player"
)

var _ = Describe("Game Engine can switch roles", func() {
	heroPlayer := &player.Player{
		Health:   1,
		Strength: 1,
		Defence:  1,
		Speed:    2,
		Luck:     0.1,
	}

	villainPlayer := &player.Player{
		Health:   1,
		Strength: 1,
		Defence:  1,
		Speed:    1,
		Luck:     0.1,
	}

	heroSpecial := &player.Special{}
	engine := &game.Engine{
		AttackerPlayer:  heroPlayer,
		AttackerSpecial: heroSpecial,
		DefenderPlayer:  villainPlayer,
		DefenderSpecial: nil,
	}

	Describe("Role switching", func() {
		Context("Starting with hero as the attacker", func() {
			It("should revert the roles", func() {
				Expect(engine.AttackerPlayer).To(Equal(heroPlayer))
				Expect(engine.AttackerSpecial).To(Equal(heroSpecial))
				Expect(engine.DefenderPlayer).To(Equal(villainPlayer))
				Expect(engine.DefenderSpecial).To(BeNil())
				engine.SwitchRoles()
				Expect(engine.AttackerPlayer).To(Equal(villainPlayer))
				Expect(engine.AttackerSpecial).To(BeNil())
				Expect(engine.DefenderPlayer).To(Equal(heroPlayer))
				Expect(engine.DefenderSpecial).To(Equal(heroSpecial))

			})
		})
	})
})
