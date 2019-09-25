package game

type Special struct {
	CriticalStrike struct {
		StrikeTwiceProbability float32
		StrikeThirdGivenTwiceProbability float32
	}
	Resilience struct {
		HalfDamageResilienceProbability float32
	}
}

type PlayerBuilder interface {
	SetHealth() PlayerBuilder
	SetStrength() PlayerBuilder
	SetDefence() PlayerBuilder
	SetSpeed() PlayerBuilder
	SetLuck() PlayerBuilder
	SetSpecialCriticalStrike() PlayerBuilder
	SetSpecialResilience() PlayerBuilder
	Build() (Player, Special)
}

