package game

import (
	// "github.com/hero-wars/utils"
	"github.com/hero-wars/config"
)


type BuildManager struct {
	playerBuilder PlayerBuilder
	cfg *config.Config
}

func (bm *BuildManager) SetBuilder (b PlayerBuilder) {
	bm.playerBuilder = b
}

func (bm *BuildManager) Construct(player struct{}) Player {
	bm.playerBuilder.SetHealth().
					SetStrength().
					SetDefence().
					SetSpeed().
					SetLuck()

	// TODO: fix
	if _, ok := bm.playerBuilder.(Hero); ok {
		bm.playerBuilder.SetHealth().
						 SetStrength()
	}

	return bm.playerBuilder.Build()
}