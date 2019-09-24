package game

import (
	"fmt"
)


type BuildManager struct {
	playerBuilder PlayerBuilder
}

func (bm *BuildManager) SetBuilder (b PlayerBuilder) {
	bm.playerBuilder = b
}

func (bm *BuildManager) Construct() Player {
	
}