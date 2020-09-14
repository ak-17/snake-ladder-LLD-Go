package model

import "fmt"

type Player struct {
	name string
	Position int64
	Won bool
}

func InitPlayer(name string) *Player  {
	return &Player{
		name:     name,
		Position: 0,
		Won:      false,
	}
}

func (p *Player) String() string  {
	return fmt.Sprintf("name: %s, position: %d",p.name,p.Position)
}