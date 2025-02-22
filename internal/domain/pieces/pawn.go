package pieces

import "chess/internal/domain"

type Pawn struct {
	X     int
	Y     int
	Color domain.Color
	Name  string `json:"name"`
}

func NewPawn(color domain.Color) *Pawn {
	return &Pawn{
		Color: color,
		Name:  "pawn",
	}
}

func (p *Pawn) Move(x int, y int) (newX int, newY int) {
	return x, y
}

func (p *Pawn) GetColor() domain.Color {
	return p.Color
}
