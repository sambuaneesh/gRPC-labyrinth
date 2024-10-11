package game

type Position struct {
	X, Y int
}

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

type Direction int

type Player struct {
	Score           int
	Health          int
	Position        Position
	RemainingSpells int
}

func NewPlayer() *Player {
	return &Player{
		Score:           0,
		Health:          3,
		Position:        Position{X: 0, Y: 0},
		RemainingSpells: 3,
	}
}

func (p *Player) CollectCoin() {
	p.Score += 1
}

func (p *Player) LoseHealth() {
	p.Health -= 1
}

func (p *Player) UseSpell() {
	p.RemainingSpells -= 1
}

func (p *Player) Move(direction Direction) Position {
	switch direction {
	case UP:
		p.Position.Y -= 1
	case DOWN:
		p.Position.Y += 1
	case LEFT:
		p.Position.X -= 1
	case RIGHT:
		p.Position.X += 1
	}
	return p.Position
}

func (p *Player) SetPosition(position Position) {
	p.Position = position
}
