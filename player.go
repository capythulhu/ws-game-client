package main

import "github.com/thzoid/ws-game-server/shared"

type Player struct {
	shared.Actor
	velocity int
	nick     rune
}

func (p *Player) Move(direction shared.Coordinate) {
	switch true {
	case direction.X < 0:
		p.Actor.Position.X = max(p.Actor.Position.X-p.velocity, 0)
	case direction.X > 0:
		p.Actor.Position.X = min(p.Actor.Position.X+p.velocity, mapHeight-1)
	}
	switch true {
	case direction.Y < 0:
		p.Actor.Position.Y = max(p.Actor.Position.Y-p.velocity, 0)
	case direction.Y > 0:
		p.Actor.Position.Y = min(p.Actor.Position.Y+p.velocity, mapWidth-1)
	}
}
