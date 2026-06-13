package learngo

import (
	"fmt"
	"math/rand"
)

type Player interface {
	KickBall()
}

type FootballPlayer struct {
	name    string
	stamina int
	power   int
}

func (p FootballPlayer) KickBall() {
	shot := p.power + p.stamina

	fmt.Printf("Player: %v can kick ball with power ^%v", p.name, shot)
}

func showPlayer() {
	players := make([]Player, 11)

	for i := range players {
		players[i] = FootballPlayer{
			name:    "Messi",
			stamina: rand.Intn(10),
			power:   rand.Intn(10),
		}
	}
}
