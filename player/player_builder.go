package player

// The go generate command is not used, as the mock needs to be adapted to test the fluent interface of PlayerBuilder
// [Unused] go:generate mockgen -destination=../mocks/mock_player_builder.go -package=mocks github.com/andrei3131/hero-wars/player PlayerBuilder

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
