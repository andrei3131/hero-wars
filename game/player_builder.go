package game

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

