package game

type Villain struct {
	player Player
}

func (h *Villain) SetHealth() PlayerBuilder {

}

func (h *Villain) SetStrength() PlayerBuilder {

}

func (h *Villain) SetDefence() PlayerBuilder {

}

func (h *Villain) SetSpeed() PlayerBuilder {

}

func (h *Villain) SetLuck() PlayerBuilder {

}

func (h *Villain) Build() Player {
	return h.Player
}