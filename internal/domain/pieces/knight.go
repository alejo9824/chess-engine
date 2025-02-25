package pieces

import "chess/internal/domain"

type Knight struct {
	X     int
	Y     int
	Color domain.Color `json:"color"`
	Name  string       `json:"name"`
}

func NewKnight(color domain.Color) *Knight {
	return &Knight{
		Color: color,
		Name:  "Knight",
	}
}

func (p *Knight) Move(x int, y int, x2 int, y2 int, targetEmpty bool) bool {

	if (x2 == x+2 || x2 == x-2) && (y2 == y+1 || y2 == y-1) {
		return true
	}
	if (x2 == x+1 || x2 == x-1) && (y2 == y+2 || y2 == y-2) {
		return true
	}

	return false

}

func (p *Knight) GetColor() domain.Color {
	return p.Color
}

func (p *Knight) GetName() string {
	return p.Name
}
