package player

import (
	"github.com/jedib0t/go-pretty/table"
)

type Special struct {
	CriticalStrike struct {
		StrikeTwiceProbability           float64
		StrikeThirdGivenTwiceProbability float64
	}
	Resilience struct {
		HalfDamageResilienceProbability float64
	}
}

func (s Special) String() string {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"Critical Strike Twice Probability", "Critical Strike Third Given Twice Probaility", "Resilience Half Damage Probability"})
	t.AppendRow([]interface{}{s.CriticalStrike.StrikeTwiceProbability, s.CriticalStrike.StrikeThirdGivenTwiceProbability, s.Resilience.HalfDamageResilienceProbability})
	return "\n" + t.Render() + "\n"
}

type Player struct {
	Health   int
	Strength int
	Defence  int
	Speed    int
	Luck     float64
}

func (p *Player) IsAlive() bool {
	return p.Health > 0
}

func (p Player) String() string {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"Health", "Strength", "Defence", "Speed", "Luck"})
	t.AppendRow([]interface{}{p.Health, p.Strength, p.Defence, p.Speed, p.Luck})
	return "\n" + t.Render() + "\n"
}
