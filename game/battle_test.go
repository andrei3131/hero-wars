package game_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/andrei3131/hero-wars/player"
	"github.com/andrei3131/hero-wars/game"

)

var _ = Describe("Battle Hero Starts", func() {
	heroPlayer := &player.Player {
		Health   : 1,
		Strength : 1,
		Defence  : 1,
		Speed    : 2,
		Luck     : 0.1,
	}

	villainPlayer := &player.Player {
		Health   : 1,
		Strength : 1,
		Defence  : 1,
		Speed    : 1,
		Luck     : 0.1,
	}


	battle := game.NewBattle(heroPlayer, &player.Special{}, villainPlayer, &game.StrikeEngine{})

	// One round
	battle.Duel()

	Describe("Battle start", func() {
		Context("With hero with higher speed", func() {
			It("should yield hero to be the attacker", func() {
				Expect(battle.Engine.AttackerPlayer).To(Equal(heroPlayer))
				Expect(battle.Engine.DefenderPlayer).To(Equal(villainPlayer))

			})
		})
	})
})


var _ = Describe("Battle Villain Starts", func() {
	heroPlayer := &player.Player {
		Health   : 1,
		Strength : 1,
		Defence  : 1,
		Speed    : 1,
		Luck     : 0.1,
	}

	villainPlayer := &player.Player {
		Health   : 1,
		Strength : 1,
		Defence  : 1,
		Speed    : 2,
		Luck     : 0.1,
	}


	battle := game.NewBattle(heroPlayer, &player.Special{}, villainPlayer, &game.StrikeEngine{})

	// One round
	battle.Duel()

	Describe("Battle start", func() {
		Context("With hero with higher speed", func() {
			It("should yield hero to be the attacker", func() {
				Expect(battle.Engine.AttackerPlayer).To(Equal(villainPlayer))
				Expect(battle.Engine.DefenderPlayer).To(Equal(heroPlayer))
			})
		})
	})
})



var _ = Describe("Battle When Speeds Equal, Highest Luck starts (Villain Luckier)", func() {
	heroPlayer := &player.Player {
		Health   : 1,
		Strength : 1,
		Defence  : 1,
		Speed    : 1,
		Luck     : 0.1,
	}

	villainPlayer := &player.Player {
		Health   : 1,
		Strength : 1,
		Defence  : 1,
		Speed    : 1,
		Luck     : 0.5,
	}


	battle := game.NewBattle(heroPlayer, &player.Special{}, villainPlayer, &game.StrikeEngine{})

	// One round
	battle.Duel()

	Describe("Battle start", func() {
		Context("With hero with higher speed", func() {
			It("should yield hero to be the attacker", func() {
				Expect(battle.Engine.AttackerPlayer).To(Equal(villainPlayer))
				Expect(battle.Engine.DefenderPlayer).To(Equal(heroPlayer))
			})
		})
	})
})


var _ = Describe("Battle When Speeds Equal, Highest Luck starts (Hero Luckier)", func() {
	heroPlayer := &player.Player {
		Health   : 1,
		Strength : 1,
		Defence  : 1,
		Speed    : 1,
		Luck     : 0.2,
	}

	villainPlayer := &player.Player {
		Health   : 1,
		Strength : 1,
		Defence  : 1,
		Speed    : 1,
		Luck     : 0.1,
	}


	battle := game.NewBattle(heroPlayer, &player.Special{}, villainPlayer, &game.StrikeEngine{})

	// One round
	battle.Duel()

	Describe("Battle start", func() {
		Context("With hero with higher speed", func() {
			It("should yield hero to be the attacker", func() {
				Expect(battle.Engine.AttackerPlayer).To(Equal(heroPlayer))
				Expect(battle.Engine.DefenderPlayer).To(Equal(villainPlayer))
			})
		})
	})
})



var _ = Describe("A duel attack/defend/attack/defend leaves roles unchanged", func() {
	heroPlayer := &player.Player {
		Health   : 1,
		Strength : 1,
		Defence  : 1,
		Speed    : 1,
		Luck     : 0.2,
	}

	villainPlayer := &player.Player {
		Health   : 1,
		Strength : 1,
		Defence  : 1,
		Speed    : 1,
		Luck     : 0.1,
	}


	heroSpecial := &player.Special{}
	battle := game.NewBattle(heroPlayer, heroSpecial, villainPlayer, &game.StrikeEngine{})

	
	// Hero starts, he is luckiest
	Describe("A duel", func() {
		Context("When hero starts first (any could start)", func() {
			It("should restore initial attacker/defender roles", func() {
				// first round
				battle.Duel()
				Expect(battle.Engine.AttackerPlayer).To(Equal(heroPlayer))
				Expect(battle.Engine.AttackerSpecial).To(Equal(heroSpecial))
				Expect(battle.Engine.DefenderPlayer).To(Equal(villainPlayer))
				Expect(battle.Engine.DefenderSpecial).To(BeNil())
				battle.Duel()
				Expect(battle.Engine.AttackerPlayer).To(Equal(heroPlayer))
				Expect(battle.Engine.AttackerSpecial).To(Equal(heroSpecial))
				Expect(battle.Engine.DefenderPlayer).To(Equal(villainPlayer))
				Expect(battle.Engine.DefenderSpecial).To(BeNil())
			})
		})
	})
})