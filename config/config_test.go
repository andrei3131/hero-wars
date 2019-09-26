package config_test

import (
	//"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/andrei3131/hero-wars/config"
)

var _ = Describe("Game Configuration Tests", func() {
	var (
		expectedConfig Config
	)


	cfg, err := ReadConfig("test_config.yml")


	BeforeEach(func() {
		expectedConfig = Config{}

		// Hero
		expectedConfig.Hero.Common.Health.HealthStart = 70
		expectedConfig.Hero.Common.Health.HealthEnd = 100
		
		expectedConfig.Hero.Common.Strength.StrengthStart = 70
		expectedConfig.Hero.Common.Strength.StrengthEnd= 80

		expectedConfig.Hero.Common.Defence.DefenceStart = 45
		expectedConfig.Hero.Common.Defence.DefenceEnd= 55

		expectedConfig.Hero.Common.Speed.SpeedStart = 40
		expectedConfig.Hero.Common.Speed.SpeedEnd= 50

		expectedConfig.Hero.Common.Luck.LuckStart = 0.1
		expectedConfig.Hero.Common.Luck.LuckEnd= 0.3

		expectedConfig.Hero.Special.CriticalStrike.StrikeTwiceProbability = 0.1
		expectedConfig.Hero.Special.CriticalStrike.StrikeThirdGivenTwiceProbability = 0.01

		expectedConfig.Hero.Special.Resilience.HalfDamageResilienceProbability = 0.2

		// Villain
		expectedConfig.Villain.Common.Health.HealthStart = 60
		expectedConfig.Villain.Common.Health.HealthEnd = 90
		
		expectedConfig.Villain.Common.Strength.StrengthStart = 60
		expectedConfig.Villain.Common.Strength.StrengthEnd= 90

		expectedConfig.Villain.Common.Defence.DefenceStart = 40
		expectedConfig.Villain.Common.Defence.DefenceEnd= 60

		expectedConfig.Villain.Common.Speed.SpeedStart = 40
		expectedConfig.Villain.Common.Speed.SpeedEnd= 60

		expectedConfig.Villain.Common.Luck.LuckStart = 0.25
		expectedConfig.Villain.Common.Luck.LuckEnd = 0.4
	})
	// End BeforeEach

	Describe("Reading config", func() {
		Context("From test file", func() {
            It("should give no error", func() {
		 		Expect(err).Should(BeNil())
			})
		})
    })

	// Hero tests
	Describe("Correctly reading reading configuration into template struct", func() {
        Context("With Hero Health Start", func() {
            It("should yield the expected value", func() {
                Expect(cfg.Hero.Common.Health.HealthStart).To(Equal(expectedConfig.Hero.Common.Health.HealthStart))
            })
        })
	})
	
	Describe("Correctly reading reading configuration into template struct", func() {
        Context("With Hero Health End", func() {
            It("should yield the expected value", func() {
				Expect(cfg.Hero.Common.Health.HealthEnd).To(Equal(expectedConfig.Hero.Common.Health.HealthEnd))
            })
        })
	})
	
		
	Describe("Correctly reading reading configuration into template struct", func() {
        Context("With Strength Strength Start", func() {
            It("should yield the expected value", func() {
				Expect(cfg.Hero.Common.Strength.StrengthStart).To(Equal(expectedConfig.Hero.Common.Strength.StrengthStart))
            })
        })
	})
			
	Describe("Correctly reading reading configuration into template struct", func() {
        Context("With Strength Strength End", func() {
            It("should yield the expected value", func() {
				Expect(cfg.Hero.Common.Strength.StrengthEnd).To(Equal(expectedConfig.Hero.Common.Strength.StrengthEnd))
            })
        })
	})
	
	Describe("Correctly reading reading configuration into template struct", func() {
        Context("With Strength Defence Start", func() {
            It("should yield the expected value", func() {
				Expect(cfg.Hero.Common.Defence.DefenceStart).To(Equal(expectedConfig.Hero.Common.Defence.DefenceStart))
            })
        })
	})
	

	Describe("Correctly reading reading configuration into template struct", func() {
        Context("With Strength Defence End", func() {
            It("should yield the expected value", func() {
				Expect(cfg.Hero.Common.Defence.DefenceEnd).To(Equal(expectedConfig.Hero.Common.Defence.DefenceEnd))
            })
        })
	})
	
	Describe("Correctly reading reading configuration into template struct", func() {
        Context("With Strength Speed Start", func() {
            It("should yield the expected value", func() {
				Expect(cfg.Hero.Common.Speed.SpeedStart).To(Equal(expectedConfig.Hero.Common.Speed.SpeedStart))
            })
        })
	})
	
	Describe("Correctly reading reading configuration into template struct", func() {
        Context("With Strength Speed End", func() {
            It("should yield the expected value", func() {
				Expect(cfg.Hero.Common.Speed.SpeedEnd).To(Equal(expectedConfig.Hero.Common.Speed.SpeedEnd))
            })
        })
	})

	Describe("Correctly reading reading configuration into template struct", func() {
        Context("With Strength Luck Start", func() {
            It("should yield the expected value", func() {
				Expect(cfg.Hero.Common.Luck.LuckStart).To(Equal(expectedConfig.Hero.Common.Luck.LuckStart))
            })
        })
	})


	Describe("Correctly reading reading configuration into template struct", func() {
        Context("With Strength Special Critical Strike", func() {
            It("should yield the expected value", func() {
				Expect(cfg.Hero.Special.CriticalStrike.StrikeTwiceProbability).To(Equal(expectedConfig.Hero.Special.CriticalStrike.StrikeTwiceProbability))
				Expect(cfg.Hero.Special.CriticalStrike.StrikeThirdGivenTwiceProbability).To(Equal(expectedConfig.Hero.Special.CriticalStrike.StrikeThirdGivenTwiceProbability))
            })
        })
	})

	Describe("Correctly reading reading configuration into template struct", func() {
        Context("With Strength Special Resilience", func() {
            It("should yield the expected value", func() {
				Expect(cfg.Hero.Special.Resilience.HalfDamageResilienceProbability).To(Equal(expectedConfig.Hero.Special.Resilience.HalfDamageResilienceProbability))
            })
        })
	})



	// Villain
	Describe("Correctly reading reading configuration into template struct", func() {
        Context("With Villain Health Start", func() {
            It("should yield the expected value", func() {
                Expect(cfg.Villain.Common.Health.HealthStart).To(Equal(expectedConfig.Villain.Common.Health.HealthStart))
            })
        })
	})
	
	Describe("Correctly reading reading configuration into template struct", func() {
        Context("With Villain Health End", func() {
            It("should yield the expected value", func() {
				Expect(cfg.Villain.Common.Health.HealthEnd).To(Equal(expectedConfig.Villain.Common.Health.HealthEnd))
            })
        })
	})
	
		
	Describe("Correctly reading reading configuration into template struct", func() {
        Context("With Strength Strength Start", func() {
            It("should yield the expected value", func() {
				Expect(cfg.Villain.Common.Strength.StrengthStart).To(Equal(expectedConfig.Villain.Common.Strength.StrengthStart))
            })
        })
	})
			
	Describe("Correctly reading reading configuration into template struct", func() {
        Context("With Strength Strength End", func() {
            It("should yield the expected value", func() {
				Expect(cfg.Villain.Common.Strength.StrengthEnd).To(Equal(expectedConfig.Villain.Common.Strength.StrengthEnd))
            })
        })
	})
	
	Describe("Correctly reading reading configuration into template struct", func() {
        Context("With Strength Defence Start", func() {
            It("should yield the expected value", func() {
				Expect(cfg.Villain.Common.Defence.DefenceStart).To(Equal(expectedConfig.Villain.Common.Defence.DefenceStart))
            })
        })
	})
	

	Describe("Correctly reading reading configuration into template struct", func() {
        Context("With Strength Defence End", func() {
            It("should yield the expected value", func() {
				Expect(cfg.Villain.Common.Defence.DefenceEnd).To(Equal(expectedConfig.Villain.Common.Defence.DefenceEnd))
            })
        })
	})
	
	Describe("Correctly reading reading configuration into template struct", func() {
        Context("With Strength Speed Start", func() {
            It("should yield the expected value", func() {
				Expect(cfg.Villain.Common.Speed.SpeedStart).To(Equal(expectedConfig.Villain.Common.Speed.SpeedStart))
            })
        })
	})
	
	Describe("Correctly reading reading configuration into template struct", func() {
        Context("With Strength Speed End", func() {
            It("should yield the expected value", func() {
				Expect(cfg.Villain.Common.Speed.SpeedEnd).To(Equal(expectedConfig.Villain.Common.Speed.SpeedEnd))
            })
        })
	})

	Describe("Correctly reading reading configuration into template struct", func() {
        Context("With Strength Luck Start", func() {
            It("should yield the expected value", func() {
				Expect(cfg.Villain.Common.Luck.LuckStart).To(Equal(expectedConfig.Villain.Common.Luck.LuckStart))
            })
        })
	})
})
