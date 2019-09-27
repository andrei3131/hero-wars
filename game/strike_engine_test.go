package game_test

import (
	//"log"

	gomock "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"

	"github.com/andrei3131/hero-wars/config"
	"github.com/andrei3131/hero-wars/mocks"
	"github.com/andrei3131/hero-wars/player"
	"github.com/andrei3131/hero-wars/game"

)

var _ = Describe("PlayerBuilder", func() {
	var (
		heroPlayer *player.Player
		heroSpecial *player.Special
		villainPlayer *player.Player

		// strikeEngineMockCallAttack    *gomock.Call
		// strikeEngineMockCallDefence    *gomock.Call
		
		mockStrikeEngine *mocks.MockStrikeEngineInterface
		mockCtrl *gomock.Controller //gomock struct
	)

	heroSpecial = &player.Special{}
	heroPlayer = &player.Player {
		Health   : 1,
		Strength : 1,
		Defence  : 1,
		Speed    : 2,
		Luck     : 0.2,
	}

	villainPlayer = &player.Player {
		Health   : 1,
		Strength : 1,
		Defence  : 1,
		Speed    : 1,
		Luck     : 0.1,
	}

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("Strike Engine", func() {
		BeforeEach(func() {
			mockStrikeEngine = mocks.NewMockStrikeEngineInterface(mockCtrl)
		})
		It("expects alternate strikes by both players, first hero attacks then villain attacks", func() {
			
			mockStrikeEngine.EXPECT().IsFirstTurnPlayed().Return(false).Times(1)
			mockStrikeEngine.EXPECT().SetFirstTurnPlayed().Times(1)

			mockStrikeEngine.EXPECT().Attack(config.HERO_NAME, config.VILLAIN_NAME, gomock.Any()).Times(1)
			mockStrikeEngine.EXPECT().Attack(config.VILLAIN_NAME, config.HERO_NAME, gomock.Any()).Times(1)			

			mockStrikeEngine.EXPECT().UpdateStrikeEngine(heroPlayer, heroSpecial, villainPlayer, nil).Times(1)			
			mockStrikeEngine.EXPECT().UpdateStrikeEngine(villainPlayer, nil, heroPlayer, heroSpecial).Times(1)			

			battle := game.NewBattle(heroPlayer, heroSpecial, villainPlayer, mockStrikeEngine)
			battle.Duel()

		})

	})

})



