package game

type PlayerBuilder interface {
	SetHealth() PlayerBuilder
	SetStrength() PlayerBuilder
	SetDefence() PlayerBuilder
	SetSpeed() PlayerBuilder
	Build() Player
}

