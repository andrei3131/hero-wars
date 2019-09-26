package player

type BuildManager struct {
	playerBuilder PlayerBuilder
}

func (bm *BuildManager) SetBuilder(b PlayerBuilder) {
	bm.playerBuilder = b
}

func (bm *BuildManager) Construct() (Player, Special) {
	bm.playerBuilder.SetHealth().
		SetStrength().
		SetDefence().
		SetSpeed().
		SetLuck().
		SetSpecialCriticalStrike().
		SetSpecialResilience()

	return bm.playerBuilder.Build()
}
