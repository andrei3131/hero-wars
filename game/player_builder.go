package game

type PlayerBuilder interface {
	SetHealth() PlayerBuilder
	SetStrength() PlayerBuilder
	SetDefence() PlayerBuilder
	SetSpeed() PlayerBuilder
	SetLuck() PlayerBuilder
	Build() Player
}

