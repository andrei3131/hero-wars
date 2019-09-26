package player_test

import (
	//"log"
	"fmt"

	gomock "github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/andrei3131/hero-wars/config"
	"github.com/andrei3131/hero-wars/gamelib"
	"github.com/andrei3131/hero-wars/mocks"
	"github.com/andrei3131/hero-wars/player"
)

var _ = Describe("PlayerBuilder", func() {
	var (
		mockCtrl *gomock.Controller //gomock struct
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Context("Fluent Interface Test", func() {
		var (
			heroMockCalls            [8]*gomock.Call
			villainMockCalls         [8]*gomock.Call
			mockHeroPlayerBuilder    *mocks.MockPlayerBuilder
			mockVillainPlayerBuilder *mocks.MockPlayerBuilder
		)

		BeforeEach(func() {
			mockHeroPlayerBuilder = mocks.NewMockPlayerBuilder(mockCtrl)
			mockVillainPlayerBuilder = mocks.NewMockPlayerBuilder(mockCtrl)

			heroMockCalls[0] = mockHeroPlayerBuilder.EXPECT().SetHealth().AnyTimes()
			heroMockCalls[1] = mockHeroPlayerBuilder.EXPECT().SetStrength().AnyTimes()
			heroMockCalls[2] = mockHeroPlayerBuilder.EXPECT().SetDefence().AnyTimes()
			heroMockCalls[3] = mockHeroPlayerBuilder.EXPECT().SetSpeed().AnyTimes()
			heroMockCalls[4] = mockHeroPlayerBuilder.EXPECT().SetLuck().AnyTimes()
			heroMockCalls[5] = mockHeroPlayerBuilder.EXPECT().SetSpecialCriticalStrike().AnyTimes()
			heroMockCalls[6] = mockHeroPlayerBuilder.EXPECT().SetSpecialResilience().AnyTimes()
			heroMockCalls[7] = mockHeroPlayerBuilder.EXPECT().Build().AnyTimes()

			villainMockCalls[0] = mockVillainPlayerBuilder.EXPECT().SetHealth().AnyTimes()
			villainMockCalls[1] = mockVillainPlayerBuilder.EXPECT().SetStrength().AnyTimes()
			villainMockCalls[2] = mockVillainPlayerBuilder.EXPECT().SetDefence().AnyTimes()
			villainMockCalls[3] = mockVillainPlayerBuilder.EXPECT().SetSpeed().AnyTimes()
			villainMockCalls[4] = mockVillainPlayerBuilder.EXPECT().SetLuck().AnyTimes()
			villainMockCalls[5] = mockVillainPlayerBuilder.EXPECT().SetSpecialCriticalStrike().AnyTimes()
			villainMockCalls[6] = mockVillainPlayerBuilder.EXPECT().SetSpecialResilience().AnyTimes()
			villainMockCalls[7] = mockVillainPlayerBuilder.EXPECT().Build().AnyTimes()
		})
		It("expects to receive all fluent interface calls once", func() {
			gamelib.BuildHeroWarsCharacters(mockHeroPlayerBuilder, mockVillainPlayerBuilder)
			for i := 0; i < len(heroMockCalls); i++ {
				heroMockCalls[i].Times(1)
			}
			for i := 0; i < len(villainMockCalls); i++ {
				villainMockCalls[i].Times(1)
			}
		})

	})

})

var _ = Describe("Player Builder Range Correctness", func() {
	cfg, _ := config.ReadConfig("../" + config.CONFIG_FILE)

	heroBuilder := player.NewHero(cfg)
	villainBuilder := player.NewVillain(cfg)

	for i := 1; i < 10000; i++ {
		heroPlayer, _, villainPlayer, _ := gamelib.BuildHeroWarsCharacters(heroBuilder, villainBuilder)

		Context(fmt.Sprintf("Hero Random Range Test: %d\n", i), func() {
			It("expects to generate a hero attribute within the configured range", func() {
				gamelib.BuildHeroWarsCharacters(heroBuilder, villainBuilder)
				Ω(heroPlayer.Health).Should(BeNumerically(">=", cfg.Hero.Common.Health.HealthStart))
				Ω(heroPlayer.Health).Should(BeNumerically("<=", cfg.Hero.Common.Health.HealthEnd))

				Ω(heroPlayer.Strength).Should(BeNumerically(">=", cfg.Hero.Common.Strength.StrengthStart))
				Ω(heroPlayer.Strength).Should(BeNumerically("<=", cfg.Hero.Common.Strength.StrengthEnd))

				Ω(heroPlayer.Defence).Should(BeNumerically(">=", cfg.Hero.Common.Defence.DefenceStart))
				Ω(heroPlayer.Defence).Should(BeNumerically("<=", cfg.Hero.Common.Defence.DefenceEnd))

				Ω(heroPlayer.Speed).Should(BeNumerically(">=", cfg.Hero.Common.Speed.SpeedStart))
				Ω(heroPlayer.Speed).Should(BeNumerically("<=", cfg.Hero.Common.Speed.SpeedEnd))

				Ω(heroPlayer.Luck).Should(BeNumerically(">=", cfg.Hero.Common.Luck.LuckStart))
				Ω(heroPlayer.Luck).Should(BeNumerically("<=", cfg.Hero.Common.Luck.LuckEnd))
			})
		})

		Context(fmt.Sprintf("Villain Random Range Test: %d\n", i), func() {
			It("expects to generate a hero attribute within the configured range", func() {
				gamelib.BuildHeroWarsCharacters(heroBuilder, villainBuilder)
				Ω(villainPlayer.Health).Should(BeNumerically(">=", cfg.Villain.Common.Health.HealthStart))
				Ω(villainPlayer.Health).Should(BeNumerically("<=", cfg.Villain.Common.Health.HealthEnd))

				Ω(villainPlayer.Strength).Should(BeNumerically(">=", cfg.Villain.Common.Strength.StrengthStart))
				Ω(villainPlayer.Strength).Should(BeNumerically("<=", cfg.Villain.Common.Strength.StrengthEnd))

				Ω(villainPlayer.Defence).Should(BeNumerically(">=", cfg.Villain.Common.Defence.DefenceStart))
				Ω(villainPlayer.Defence).Should(BeNumerically("<=", cfg.Villain.Common.Defence.DefenceEnd))

				Ω(villainPlayer.Speed).Should(BeNumerically(">=", cfg.Villain.Common.Speed.SpeedStart))
				Ω(villainPlayer.Speed).Should(BeNumerically("<=", cfg.Villain.Common.Speed.SpeedEnd))

				Ω(villainPlayer.Luck).Should(BeNumerically(">=", cfg.Villain.Common.Luck.LuckStart))
				Ω(villainPlayer.Luck).Should(BeNumerically("<=", cfg.Villain.Common.Luck.LuckEnd))
			})
		})
	}
})
