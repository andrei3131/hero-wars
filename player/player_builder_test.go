package player_test

import (
	//"log"
	//"fmt"

	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"
	gomock "github.com/golang/mock/gomock"

	"github.com/andrei3131/hero-wars/mocks"
	//"github.com/andrei3131/hero-wars/player"
	"github.com/andrei3131/hero-wars/gamelib"
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
