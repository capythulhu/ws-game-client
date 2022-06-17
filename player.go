package main

type Player struct {
	Actor
	velocity int
	nick     rune
}

func (p *Player) Move(direction Coordinate) {
	switch true {
	case direction.x < 0:
		p.Actor.position.x = max(p.Actor.position.x-p.velocity, 0)
	case direction.x > 0:
		p.Actor.position.x = min(p.Actor.position.x+p.velocity, mapHeight-1)
	}
	switch true {
	case direction.y < 0:
		p.Actor.position.y = max(p.Actor.position.y-p.velocity, 0)
	case direction.y > 0:
		p.Actor.position.y = min(p.Actor.position.y+p.velocity, mapWidth-1)
	}
}
