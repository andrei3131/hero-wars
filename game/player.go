package game

import (
	//"os"
	// "fmt"

	"github.com/jedib0t/go-pretty/table"
)

type Special struct {
	CriticalStrike struct {
		StrikeTwiceProbability float32
		StrikeThirdGivenTwiceProbability float32
	}
	Resilience struct {
		HalfDamageResilienceProbability float32
	}
}

type Player struct {
	Health int
	Strength int
	Defence int 
	Speed int
	Luck float32
}


func (p *Player) IsAlive() bool {
	return p.Health > 0
}

func (p Player) String() string {
	t := table.NewWriter()
    //t.SetOutputMirror(os.Stdout)
    t.AppendHeader(table.Row{"Health", "Strength", "Defence", "Speed", "Luck"})
    t.AppendRow([]interface{}{p.Health, p.Health, p.Defence, p.Speed, p.Luck})
    return "\n\n" + t.Render() + "\n\n"
}